package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"

	pb "github.com/QuanDN22/BE/gRPC/grpc-gateway-httpbody-message/proto/httpbody_example"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HttpBodyExampleService struct {
	pb.UnimplementedHttpBodyExampleServiceServer
}

func (s *HttpBodyExampleService) Helloworld(_ context.Context, in *emptypb.Empty) (*httpbody.HttpBody, error) {
	return &httpbody.HttpBody{
		ContentType: "text/html",
		Data:        []byte("Hello World"),
	}, nil
}

func (s *HttpBodyExampleService) Download(_ *emptypb.Empty, stream pb.HttpBodyExampleService_DownloadServer) error {
	msgs := []*httpbody.HttpBody{
		{
			ContentType: "text/html",
			Data:        []byte("Hello 1"),
		},
		{
			ContentType: "text/html",
			Data:        []byte("Hello 2"),
		},
	}

	for _, msg := range msgs {
		if err := stream.Send(msg); err != nil {
			return err
		}
	}

	return nil
}

func (s *HttpBodyExampleService) Upload(stream pb.HttpBodyExampleService_UploadServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&httpbody.HttpBody{
				ContentType: "text/html",
				Data:        []byte("Upload received"),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}

		type server struct {
			Server_Names  string
			Server_IPv4   string
			Server_Status string
		}

		var res server
		err = json.Unmarshal(req.Content, &res)
		if err != nil {
			log.Fatalf("Error while unmarshalling: %v", err)
		}
		log.Println(res)

		// res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}

func main() {
	const addr = ":8083"
	// Create listening on TCP port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterHttpBodyExampleServiceServer(s, &HttpBodyExampleService{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8083")
	log.Fatalln(s.Serve(lis))
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
}
