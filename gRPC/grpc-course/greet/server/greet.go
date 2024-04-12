package main

import (
	"context"
	pb "grpc-go-course/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetRespone, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetRespone{
		Result: "Hello" + in.FirstName,
	}, nil
}
