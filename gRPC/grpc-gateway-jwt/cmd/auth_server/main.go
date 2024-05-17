package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/pkg/middleware"
	authpb "github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/proto/auth"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	authpb.UnimplementedAuthServiceServer
	issuer *middleware.Issuer
}

func NewServer(iss *middleware.Issuer) (*AuthServer, error) {
	return &AuthServer{
		issuer: iss,
	}, nil
}

func (s *AuthServer) SayHello(ctx context.Context, in *authpb.HelloRequest) (*authpb.HelloReply, error) {
	// not having a token is now an exceptional state and we can just
	// let the context helper panic if that happens
	// token := middleware.MustContextGetToken(ctx)
	token, err := middleware.ContextGetToken(ctx)
	if err != nil {
		panic(err)
	}

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &authpb.HelloReply{
		Message: fmt.Sprintf("Auth service hi %s, roles %v", in.GetName(), roles),
	}, nil
}

func (s *AuthServer) Ping(ctx context.Context, _ *emptypb.Empty) (*authpb.HelloReply, error) {
	return &authpb.HelloReply{
		Message: "Pong",
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, in *authpb.LoginRequest) (*authpb.LoginReply, error) {
	fmt.Printf("Login request %s, %s", in.GetUsername(), in.GetPassword())
	if in.GetUsername() != "admin" || in.GetPassword() != "pass" {
		return nil, fmt.Errorf("Invalid username or password")
	}

	tokenString, err := s.issuer.IssueToken("admin", []string{"admin"})
	if err != nil {
		return nil, err
	}

	return &authpb.LoginReply{
		AccessToken: tokenString,
	}, nil
}

func main() {
	iss, err := middleware.NewIssuer("./auth.ed")
	if err != nil {
		log.Fatalln("Failed to create issuer: ", err)
	}

	authserver, err := NewServer(iss)
	if err != nil {
		log.Fatalln("Failed to create server: ", err)
	}

	mw, err := middleware.NewMiddleware(os.Args[1])
	// mw, err := middleware.NewMiddleware("./auth.ed.pub")
	// mw, err := middleware.NewMiddleware("./auth.pub.key")
	if err != nil {
		log.Fatalln("Failed to create middleware: ", err)
	}

	const addr = ":8080"
	// Create listening on TCP port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to listen: ", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer(
		grpc.UnaryInterceptor(mw.UnaryServerInterceptor),
	)
	// Attach the Greeter service to the server
	authpb.RegisterAuthServiceServer(s, authserver)

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatalln(s.Serve(lis))
}