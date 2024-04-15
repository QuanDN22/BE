package main

import (
	"fmt"
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetRespone{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		// log.Printf("LongGreet function was invoked with %v", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName) 
	}
}
