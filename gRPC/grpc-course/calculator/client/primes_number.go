package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Printf("doPrimes was invoked")

	req := &pb.PrimesRequest{
		Number: 120,
	}

	stream, err := c.PrimesNumber(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while receiving message: %v\n", err)
		}

		log.Printf("Primes received: %d\n", msg.Primes)
	}
}
