package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked") 
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "QuanDN",
	}) 

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting : %s\n", res.Result)
}