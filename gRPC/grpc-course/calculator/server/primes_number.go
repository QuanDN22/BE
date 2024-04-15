package main

import (
	pb "grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) PrimesNumber(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesNumberServer) error {
	log.Printf("GreetManyTimes function was invoked with %v", in)

	N := in.Number
	var k int32 = 2

	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.PrimesResponse{
				Primes: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}
