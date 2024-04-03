package main

import (
	"fmt"
)

func main() {
	// single variable
	//
	var number int = 20
	fmt.Println(number)

	//
	number = 10
	fmt.Println(number)

	// type inference
	var email = "abc@gmail.com"
	fmt.Println(email)

	// some variable
	// same type
	var a, b int
	a = 1
	b = 2
	fmt.Println(a)
	fmt.Println(b)

	//
	var c, d int = 4, 5
	fmt.Println(c)
	fmt.Println(d)

	//
	var x, y = 6, 7
	fmt.Println(x)
	fmt.Println(y)

	// different type
	var (
		name    string
		address string
		age     int
	)
	name = "Quan"
	address = "BN"
	age = 21
	fmt.Println(name, address, age)

	//
	var (
		name1    string = "name1"
		address1 string = "address1"
		age1     int    = 21
	)
	fmt.Println(name1, address1, age1)

	//
	var name2, address2, age2 = "name2", "address2", 25
	fmt.Println(name2, address2, age2)

	language, num := "Golang", 5
	fmt.Println(language, num)
}
