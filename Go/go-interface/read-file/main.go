package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		// log.Fatal(err)
		fmt.Println("err:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
