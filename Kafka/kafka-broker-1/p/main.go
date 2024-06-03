package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/QuanDN22/BE/Kafka/kafka-broker-1/config"
	"github.com/QuanDN22/BE/Kafka/kafka-broker-1/producer"
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

	time.Sleep(5 * time.Second)

	// pd := producer.NewProducer(ctx)
	pd.Start(ctx)
}
