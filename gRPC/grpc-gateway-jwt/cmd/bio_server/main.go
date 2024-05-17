package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/pkg/middleware"
	biopb "github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/proto/bio"
	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BioServer struct {
	biopb.UnimplementedBioServiceServer
}

// func NewServer() *server {
// 	return &server{}
// }

func (s *BioServer) SayHi(ctx context.Context, _ *emptypb.Empty) (*biopb.HelloReply, error) {
	// not having a token is now an exceptional state and we can just
	// let the context helper panic if that happens
	token := middleware.MustContextGetToken(ctx)

	// dig the roles from the claims
	roles := token.Claims.(jwt.MapClaims)["roles"]

	return &biopb.HelloReply{
		Message: fmt.Sprintf("BioServer %v", roles),
	}, nil
}

func main() {
	// mw, err := middleware.NewMiddleware(os.Args[1])
	mw, err := middleware.NewMiddleware("auth.ed.pub")
	if err != nil {
		log.Fatalln("Failed to create middleware: ", err)
	}

	const addr = ":8085"
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
	biopb.RegisterBioServiceServer(s, &BioServer{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8085")
	log.Fatalln(s.Serve(lis))
}
