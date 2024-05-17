package main

import (
	"fmt"
	"os"

	"github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/pkg/middleware"
)

func main() {
	issuer, err := middleware.NewIssuer(os.Args[1])
	if err != nil {
		fmt.Printf("unable to create issuer: %v\n", err)
		os.Exit(1)
	}

	token, err := issuer.IssueToken("admin", []string{"admin", "basic"})
	if err != nil {
		fmt.Printf("unable to issue token: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(token)

	//

	v, err := middleware.NewValidator(os.Args[2])
	if err != nil {
		fmt.Printf("unable to create validator: %v\n", err)
		os.Exit(1)
	}

	token1, err := v.GetToken(token)
	if err != nil {
		fmt.Printf("unable to get validated token: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(token1.Claims)
}
