package main

import (
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) MaxNumber(stream pb.CalculatorService_MaxNumberServer) error {
	log.Printf("MaxNumber function was invoked with %v\n", stream)

	var maxNumber int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		// maxNumber = max(maxNumber, req.Number)
		// err = stream.Send(&pb.MaxResponse{
		// 	Result: maxNumber,
		// })

		if number := req.Number; number > maxNumber {
			maxNumber = number
			err := stream.Send(&pb.MaxResponse{
				Result: maxNumber,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
