package main

import (
	"context"
	"log"
	"net/http"

	helloworldpb "github.com/QuanDN22/BE/gRPC/go-grpc-gateway-helloworld/proto/helloworld"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// // Create a client connection to the gRPC server we just created
	// // This is where the gRPC-gateway proxies the requests

	// approach 1
	// conn, err := grpc.NewClient(
	// 	"0.0.0.0:8080",
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// )
	// if err != nil {
	// 	log.Fatalln("Failed to dial server:", err)
	// }

	// gwmux := runtime.NewServeMux()
	// // Register Greeter
	// err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// approach 2
	gwmux := runtime.NewServeMux()
	
	dialOptions := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Register Greeter
	err := helloworldpb.RegisterGreeterHandlerFromEndpoint(ctx, gwmux, "0.0.0.0:8080", dialOptions)
	if err != nil {
		log.Fatalln("Failed to register gateway: ", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway is running on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
