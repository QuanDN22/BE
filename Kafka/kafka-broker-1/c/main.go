package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/BE/Kafka/kafka-broker-1/config"
	"github.com/QuanDN22/BE/Kafka/kafka-broker-1/consumer"
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
