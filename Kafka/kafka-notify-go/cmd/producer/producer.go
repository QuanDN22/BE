package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/QuanDN22/Kafka/kafka-notify-go/pkg/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	ProducerPort       = ":8080"
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

// ============== HELPER FUNCTIONS ==============
var ErrUserNotFoundInProducer = errors.New("user not found")

func findUserByID(id int, users []models.User) (models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, ErrUserNotFoundInProducer
}

func getIDFromRequest(formValue string, ctx *gin.Context) (int, error) {
	fmt.Printf("formValueConvert: %s, type %T\n", formValue, formValue)
	formValueConvert := ctx.PostForm(formValue)
	fmt.Printf("formValueConvert: %s\n", formValueConvert)
	id, err := strconv.Atoi(formValueConvert)
	fmt.Printf("id: %d\n", id)
	if err != nil {
		return 0, fmt.Errorf("failed to parse ID from form value %s: %w", formValue, err)
	}

	return id, nil
}

// ============== KAFKA RELATED FUNCTIONS ==============
func sendKafkaMessage(producer sarama.SyncProducer, users []models.User, ctx *gin.Context, fromID int, toID int) error {
	// This function starts by retrieving the message from the context
	message := ctx.PostForm("message")
	fmt.Printf("message: %s\n", message)

	// Attempts to find both the sender and the recipient using their IDs.
	fromUser, err := findUserByID(fromID, users)
	if err != nil {
		return err
	}

	toUser, err := findUserByID(toID, users)
	if err != nil {
		return err
	}

	// Initializes a Notification struct that encapsulates information
	// about the sender, the recipient, and the actual message
	notification := models.Notification{
		From:    fromUser,
		To:      toUser,
		Message: message,
	}

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %w", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(toUser.ID)),
		Value: sarama.StringEncoder(notificationJSON),
	}

	// Sends the constructed message to the "notifications" topic
	_, _, err = producer.SendMessage(msg)
	return err
}

func SendMessageHandler(producer sarama.SyncProducer, users []models.User) gin.HandlerFunc {
	// This function serves as an endpoint handler for the /send POST request.
	// It processes the incoming request to ensure valid sender and recipient IDs are provided.
	return func(ctx *gin.Context) {
		fromID, err := getIDFromRequest("fromID", ctx)
		if err != nil {
			// a 400 Bad Request for invalid IDs
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		toID, err := getIDFromRequest("toID", ctx)
		if err != nil {
			// a 400 Bad Request for invalid IDs
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = sendKafkaMessage(producer, users, ctx, fromID, toID)
		if errors.Is(err, ErrUserNotFoundInProducer) {
			// a 404 Not Found for nonexistent users
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		}
		if err != nil {
			// a 500 Internal Server Error for other failures
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Notification sent successfully!",
		})
	}
}

func setupProducer() (sarama.SyncProducer, error) {
	// Initializes a new default configuration for Kafka.
	config := sarama.NewConfig()

	// Ensures that the producer receives an acknowledgment
	// once the message is successfully stored in the "notifications" topic.
	config.Producer.Return.Successes = true

	// Initializes a synchronous Kafka producer
	// that connects to the Kafka broker running at localhost:9092
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}

	return producer, nil
}

func main() {
	users := []models.User{
		{ID: 1, Name: "User#1"},
		{ID: 2, Name: "User#2"},
		{ID: 3, Name: "User#3"},
		{ID: 4, Name: "User#4"},
	}

	// Initialize a Kafka producer
	producer, err := setupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)

	// Create a Gin router, setting up a web server.
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET"}
	router.Use(cors.New(config))

	// Define a POST endpoint /send to handle notifications.
	// This endpoint expects the sender and recipient’s IDs and the message content.
	router.POST("/send", SendMessageHandler(producer, users))

	fmt.Printf("Kafka Producer started at http://localhost%s\n", ProducerPort)
	if err := router.Run(ProducerPort); err != nil {
		log.Fatalf("failed to run the server : %v", err)
	}
}
