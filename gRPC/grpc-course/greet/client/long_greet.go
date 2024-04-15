package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Qu"},
		{FirstName: "Quan"},
		{FirstName: "QuanDN"},
		{FirstName: "QuanDN2"},
		{FirstName: "QuanDN22"},
		{FirstName: "QuanDN222"},
		{FirstName: "QuanDN2222"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		// time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
