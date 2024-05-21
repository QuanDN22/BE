package main

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/Kafka/ping-pong-service/config"
	"github.com/QuanDN22/Kafka/ping-pong-service/producer"
)

func main() {
	// config
	cfg, err := config.NewConfig("./", ".env")
	if err != nil {
		log.Fatalf("failed get config %v", err)
	}
	fmt.Println("producer")
	ctx := context.Background()
	pd := producer.NewProducer(ctx, cfg.KafkaBrokerAddress, cfg.KafkaTopic)
	// pd := producer.NewProducer(ctx)
	pd.Start(ctx)
}
