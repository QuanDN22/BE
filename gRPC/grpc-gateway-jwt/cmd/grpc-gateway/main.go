package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/pkg/middleware"
	authpb "github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/proto/auth"
	biopb "github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/proto/bio"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func withLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/hello" {
			fmt.Println("it is login")
			h.ServeHTTP(w, r)
		} else {
			fmt.Printf("Run request, http_method %s and http_url ", r.Method)
			fmt.Println(r.URL)
			h.ServeHTTP(w, r)
		}
	})
}

func main() {
	// create middleware using the given public key path
	mw, err := middleware.NewMiddleware("./auth.ed.pub")
	if err != nil {
		panic(err)
	}

	// // Create a client connection to the gRPC server we just created
	// // This is where the gRPC-gateway proxies the requests

	// approach 1
	conn_auth, err := grpc.NewClient(
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	conn_bio, err := grpc.NewClient(
		"0.0.0.0:8085",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(mw.UnaryClientInterceptor),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = authpb.RegisterAuthServiceHandler(context.Background(), gwmux, conn_auth)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	err = biopb.RegisterBioServiceHandler(context.Background(), gwmux, conn_bio)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// approach 2
	// gwmux := runtime.NewServeMux()

	// dialOptions := []grpc.DialOption{
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()),
	// }

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // Register Greeter
	// // Register gRPC server endpoint
	// // Note: Make sure the gRPC server is running properly and accessible
	// err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, "0.0.0.0:8080", dialOptions)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway from auth service: ", err)
	// }
	// err = biopb.RegisterBioServiceHandlerFromEndpoint(ctx, gwmux, "0.0.0.0:8085", dialOptions)
	// if err != nil {
	// 	log.Fatalln("Failed to register gateway from bio service: ", err)
	// }

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: withLogger(mw.HandleHTTP(gwmux)),
	}

	log.Println("Serving gRPC-Gateway is running on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
