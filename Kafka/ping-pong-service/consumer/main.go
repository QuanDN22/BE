package main

import (
	"context"
	"fmt"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
const (
	topic         = "message-log"
	BrokerAddress = "localhost:9092"
)

type Consumer struct{}

func (c *Consumer) consumer(ctx context.Context) {
	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{BrokerAddress},
		Topic:       topic,
		GroupID:     "my-group",
		StartOffset: kafka.FirstOffset,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		split := strings.Split(string(msg.Value), ",")
		user_id := split[0]
		status := split[1]

		// after receiving the message, log its value
		fmt.Println("received: ", string(msg.Value))
		fmt.Println("user_id: ", user_id)
		fmt.Println("status: ", status)
	}
}

func main() {
	var c Consumer
	c.consumer(context.Background())
}
