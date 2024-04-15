package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Printf("doAvg was invoked")

	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
		{Number: 6},
		{Number: 7},
		{Number: 8},
		{Number: 9},
		{Number: 10},
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}
