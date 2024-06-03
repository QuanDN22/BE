package consumer

import (
	"context"
	"fmt"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
// const (
// 	topic         = "message-log"
// 	BrokerAddress = "localhost:9092"
// )

type Consumer struct {
	Consumer *kafka.Reader
}

func NewConsumer(ctx context.Context, BrokerAddress string, topic string, groupID string) *Consumer {
	// func NewConsumer(ctx context.Context) *Consumer {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		// Brokers:     []string{BrokerAddress},
		// Topic:       topic,
		// GroupID:     "my-group",
		Brokers:     []string{BrokerAddress},
		Topic:       topic,
		GroupID:     groupID,
		StartOffset: kafka.FirstOffset,
	})

	return &Consumer{
		Consumer: r,
	}
}

func Woker(msg kafka.Message) {
	fmt.Println("received: ", string(msg.Value))
	fmt.Printf("key: %s, offset: %d\n", string(msg.Key), msg.Offset)

	split := strings.Split(string(msg.Value), ",")
	user_id := split[0]
	status := split[1]

	// after receiving the message, log its value
	fmt.Printf("user_id: %s, status: %s\n", user_id, status)
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := c.Consumer.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		go Woker(msg)

		// split := strings.Split(string(msg.Value), ",")
		// user_id := split[0]
		// status := split[1]

		// // after receiving the message, log its value
		// fmt.Println("received: ", string(msg.Value))
		// fmt.Println("user_id: ", user_id)
		// fmt.Println("status: ", status)
	}
}
