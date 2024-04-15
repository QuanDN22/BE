package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMaxNumber(c pb.CalculatorServiceClient) {
	log.Println("doMaxNumber was invoked")

	stream, err := c.MaxNumber(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*pb.MaxRequest{
		{Number: 1}, 
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	waitc := make(chan struct{})
 
	go func() {
		for _, req := range reqs {
			log.Printf("Sending request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %v\n", err)
				break
			}

			log.Printf("Received response: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}
