package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	//create a connection
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)
	// go doGreetManyTimes(c)
	// go doGreet(c)
	// doLongGreet(c)
	doGreetEveryone(c)
	time.Sleep(5 * time.Second)
}
