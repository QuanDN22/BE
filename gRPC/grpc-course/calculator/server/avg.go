package main

import (
	pb "grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Printf("Avg function was invoked")

	var total int32 = 0
	var count int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			if count == 0 {
				return stream.SendAndClose(&pb.AvgResponse{
					Result: 0,
				})
			} else {
				return stream.SendAndClose(&pb.AvgResponse{
					Result: float32(total) / float32(count),
				})
			}
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		log.Printf("Avg function was invoked with %v", req)
		total += req.Number
		count++
	}
}
