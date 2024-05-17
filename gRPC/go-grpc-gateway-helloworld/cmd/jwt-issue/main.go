package main

import (
	"fmt"
	"os"

	"github.com/QuanDN22/BE/gRPC/go-grpc-gateway-helloworld/pkg/simplejwt"
)

func main() {
	issuer, err := simplejwt.NewIssuer(os.Args[1])
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
}
