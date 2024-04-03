package main

import "fmt"

func main() {
	// if else
	number := 10

	if number == 10 {
		fmt.Println("number = 10")
	} else if number < 10 && number != 0 {
		fmt.Println("number < 10")
	} else {
		fmt.Println("number > 10")
	}

	// With init statement
	if num := 9; num < 10 { // num is a local variable
		fmt.Println(num)
	}

	// switch case
	switch number {
	case 10:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	//
	switch {
	case number == 10:
		fmt.Println(1)
		break
	case number > 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}

	// for
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}

	for number < 20 {
		fmt.Print(number, " ")
		number++
	} 
}
