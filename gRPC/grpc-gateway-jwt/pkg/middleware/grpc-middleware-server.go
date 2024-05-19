package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (m *Middleware) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// fmt.Println(info.FullMethod)
	// check login request
	if info.FullMethod == "/auth.AuthService/Login" {
		fmt.Println("4. start /auth.AuthService/Login")
		// fmt.Println(info.FullMethod)
		// log.Println("gRPC-middleware-server UnaryServerInterceptor - /auth.AuthService/Login")
		return handler(ctx, req)
	}

	fmt.Println("4. start")
	// fmt.Println("gRPC-middleware-server UnaryServerInterceptor - /auth.AuthService/Login")

	// rip the token from the metadata from the context
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}

	tokens := headers.Get("jwt")
	if len(tokens) < 1 {
		return nil, status.New(codes.Unauthenticated, "no auth provided").Err()
	}

	tokenString := tokens[0] // just use the first, ignore repeated headers

	token, err := m.GetToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Save the token in the context for later use
	ctx = ContextWithToken(ctx, token)

	// fmt.Println("* gRPC SERVER middleware validated and set token")

	// call the next handler, with the updated context
	return handler(ctx, req)
}
