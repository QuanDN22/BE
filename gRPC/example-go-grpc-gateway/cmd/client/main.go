// ./cmd/client/main.go

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/QuanDN22/BE/gRPC/example-go-grpc-gateway/protogen/golang/orders"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the order server.
	orderServiceAddr := "localhost:50051"
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	if err = orders.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	// start listening to requests from the gateway server
	addr := "localhost:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}


	// gwmux := runtime.NewServeMux()
	// // Register Greeter
	// err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway:", err)
	// }

	// gwServer := &http.Server{
	// 	Addr:    ":8090",
	// 	Handler: gwmux,
	// }

	// log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	// log.Fatalln(gwServer.ListenAndServe())
}
