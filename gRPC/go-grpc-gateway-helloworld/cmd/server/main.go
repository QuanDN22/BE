package main

import (
	"context"
	"log"
	"net"

	helloworldpb "github.com/QuanDN22/BE/gRPC/go-grpc-gateway-helloworld/proto/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{
		Message: in.Name + "world",
	}, nil
}

func (s *server) SayHi(ctx context.Context, _ *emptypb.Empty) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{
		Message: "hi world",
	}, nil
}

func main() {
	const addr = ":8080"
	// Create listening on TCP port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, &server{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatalln(s.Serve(lis))
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
}
