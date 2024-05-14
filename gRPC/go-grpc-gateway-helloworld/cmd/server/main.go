package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	helloworldpb "github.com/QuanDN22/BE/gRPC/go-grpc-gateway-helloworld/proto/helloworld"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/QuanDN22/BE/gRPC/go-grpc-gateway-helloworld/pkg/simplejwt"
)

type server struct {
	helloworldpb.UnimplementedGreeterServer
	validator *simplejwt.Validator
}

func NewServer(validator *simplejwt.Validator) (*server, error) {
	return &server{
		validator: validator,
	}, nil
}

// SayHello implements helloworld.GreeterServer it requires a valid
// token in the context and prints the included preferred name from the request
// and roles from the token
func (s *server) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	token, err := s.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &helloworldpb.HelloReply{
		Message: fmt.Sprintf(
			"Hello %s! I am the backend. You have roles %v",
			in.GetName(), roles),
	}, nil
}

func (s *server) SayHi(ctx context.Context, _ *emptypb.Empty) (*helloworldpb.HelloReply, error) {
	token, err := s.tokenFromContextMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get token: %w", err)
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &helloworldpb.HelloReply{
		Message: fmt.Sprintf(
			"Hi, I am the backend. You have roles %v", roles),
	}, nil
}

func (b *server) tokenFromContextMetadata(ctx context.Context) (*jwt.Token, error) {
	// rip the token from the metadata via the context
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("no metadata found in context")
	}
	tokens := headers.Get("jwt")
	if len(tokens) < 1 {
		return nil, errors.New("no token found in metadata")
	}
	tokenString := tokens[0]

	token, err := b.validator.GetToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	return token, nil
}

func main() {
	// init
	validator, err := simplejwt.NewValidator(os.Args[1])
	if err != nil {
		panic(err)
	}

	server, err := NewServer(validator)
	if err != nil {
		panic(err)
	}

	const addr = ":8083"
	// Create listening on TCP port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, server)

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8083")
	log.Fatalln(s.Serve(lis))
	// go func() {
	// 	log.Fatalln(s.Serve(lis))
	// }()
}
