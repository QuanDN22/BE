package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Kafka/ping-pong-service/config"
	"github.com/QuanDN22/Kafka/ping-pong-service/consumer"
)

func main() {
	// config
	cfg, err := config.NewConfig("./", ".env")
	if err != nil {
		log.Fatalf("failed get config %v", err)
	}
	fmt.Println("consumer")
	ctx := context.Background()
	cs := consumer.NewConsumer(ctx, cfg.KafkaBrokerAddress, cfg.KafkaTopic, cfg.KafkaConsumerGroupId)
	// cs := consumer.NewConsumer(ctx)
	cs.Start(ctx)
}
