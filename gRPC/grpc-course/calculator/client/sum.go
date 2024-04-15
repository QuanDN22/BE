package main

import (
	"context"
	pb "grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked") 
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 2,
	}) 

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum : %v\n", res.Result)
}