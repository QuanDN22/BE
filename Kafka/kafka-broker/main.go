package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "192.168.1.61:9092"
	topic := "test"
	groupID := "my-group"

	config := kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		GroupID:     groupID,
		Topic:       topic,
		MaxWait:     100,
		StartOffset: kafka.LastOffset,
	}

	reader := kafka.NewReader(config)
	defer reader.Close()

	fmt.Println("Consumer started...")


	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:   "test",
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	go func() {
		for {
			message, err := reader.ReadMessage(context.Background())
			if err != nil {
				fmt.Printf("Error reading message: %v\n", err)
				continue
			}
	
			fmt.Printf("Received message: %s\n", string(message.Value))
		}
	}()
	time.Sleep(10 * time.Minute)
}
