package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/IBM/sarama"
	"github.com/QuanDN22/Kafka/kafka-notify-go/pkg/models"
	"github.com/gin-gonic/gin"
)

const (
	ConsumerGroup      = "notifications-group"
	ConsumerTopic      = "notifications"
	ConsumerPort       = ":8081"
	KafkaServerAddress = "localhost:9092"
)

// ============== HELPER FUNCTIONS ==============
var ErrNoMessageFound = errors.New("no message found")

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessageFound
	}
	return userID, nil
}

// ====== NOTIFICATION STORAGE ======
type UserNotifications map[string][]models.Notification

type NotificationStore struct {
	data UserNotifications
	mu   sync.RWMutex
}

func (ns *NotificationStore) Add(userID string, notification models.Notification) {
	ns.mu.Lock()
	defer ns.mu.Unlock()
	ns.data[userID] = append(ns.data[userID], notification)
}

func (ns *NotificationStore) Get(userID string) []models.Notification {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	return ns.data[userID]
}

// ============== KAFKA RELATED FUNCTIONS ==============
type Consumer struct {
	store *NotificationStore // which is a reference to the NotificationStore to keep track of the received notifications.
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// The consumer listens for new messages on the topic
	for msg := range claim.Messages() {
		// For each message, it fetches the userID (the Key of the message)
		userID := string(msg.Key)
		
		var notification models.Notification
		// Un-marshals the message into a Notification struct.
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}

		// Adds the notification to the NotificationStore.
		consumer.store.Add(userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	// Initializes a new default configuration for Kafka.
	config := sarama.NewConfig()

	// Creates a new Kafka consumer group that connects to the broker running on localhost:9092
	// The group name is "notifications-group".
	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

// This function sets up the Kafka consumer group, 
// listens for incoming messages, 
// and processes them using the Consumer struct methods.
func setupConsumerGroup(ctx context.Context, store *NotificationStore) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	// consuming messages from the “notifications” topic 
	// and processing any errors that arise.
	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func handleNotifications(ctx *gin.Context, store *NotificationStore) {
	// It attempts to retrieve the userID from the request.
	userID, err := getUserIDFromRequest(ctx)

	// If it doesn’t exist, it returns a 404 Not Found status.
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	// Fetches the notifications for the provided user ID from the NotificationStore
	notes := store.Get(userID)
	
	// an empty notifications slice
	if len(notes) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message":       "No notifications found for user",
				"notifications": []models.Notification{},
			})
		return
	}

	// Sends back the current notifications
	ctx.JSON(http.StatusOK, gin.H{"notifications": notes})
}

func main() {
	// Creates an instance of NotificationStore to hold the notifications
	store := &NotificationStore{
		data: make(UserNotifications),
	}

	// Sets up a cancellable context that can be used to stop the consumer group.
	ctx, cancel := context.WithCancel(context.Background())
	
	// Starts the consumer group in a separate Goroutine, 
	// allowing it to operate concurrently without blocking the main thread.
	go setupConsumerGroup(ctx, store)
	defer cancel()

	// Create a Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Define a GET endpoint
	// Fetch the notifications for a specific user 
	// via the handleNotifications() function when accessed.
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		handleNotifications(ctx, store)
	})

	fmt.Printf("Kafka CONSUMER (Group: %s)"+
		"started at http://localhost%s\n", ConsumerGroup, ConsumerPort)

	if err := router.Run(ConsumerPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
