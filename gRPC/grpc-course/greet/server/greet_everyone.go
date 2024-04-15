package main

import (
	pb "grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked with %v", stream)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err!= nil {
            log.Fatalf("Error while reading client stream: %v\n", err)
        }

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetRespone{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}


}
