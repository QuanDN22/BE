package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

// the topic and broker address are initialized as constants
const (
	topic         = "message-log"
	BrokerAddress = "localhost:9092"
)

type Producer struct{}

func (p *Producer) producer(ctx context.Context) {
	// i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{BrokerAddress},
		Topic:   topic,
	})

	for {
		var src = rand.NewSource(time.Now().UnixNano())
		var r = rand.New(src)

		user_id := r.Intn(100) + 1
		status := r.Intn(2) == 1

		// each kafka message has a key and value. The key is used
		// to decide which partition (and consequently, which broker)
		// the message gets published on
		err := w.WriteMessages(ctx, kafka.Message{
			// Key: []byte(strconv.Itoa(i)),
			Key: []byte(strconv.Itoa(user_id)),
			// create an arbitrary message payload for the value
			// Value: []byte("this is message" + strconv.Itoa(i)),
			// Value: []byte(strconv.Itoa(user_id)+status),
			Value: []byte(fmt.Sprintf("%d,%t", user_id, status)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		// log a confirmation once the message is written
		// fmt.Println("writes:", i)
		// i++
		msg := fmt.Sprintf(`{"user_id": %d, "status": %t}`, user_id, status)
		fmt.Println(msg)
		// sleep for a second
		time.Sleep(time.Second*3)
	}
}

func main() {
	var p Producer
	p.producer(context.Background())
}
