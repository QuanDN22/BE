package main

import (
	"fmt"
	"reflect"
)

func plus(a int, b int) int {
	return a + b
}

func mod(a, b int) {
	fmt.Println(a % b)
}

func cal(a, b int) (int, int) {
	return a + b, a % b
}

// Passing Address to a Function
func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	// return
}

// Variadic Functions
// A variadic function is a function that accepts a variable number of arguments
func sum(nums ...int) int {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

// Select single argument from all arguments of variadic function
func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

// Passing multiple string arguments to a variadic function
func variadicExample1(s ...string) {
	fmt.Println(s)
}

// Pass different types of arguments in variadic function
func variadicExample2(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main() {
	//
	fmt.Println(plus(1, 2))
	mod(2, 3)
	fmt.Println(cal(3, 4))

	// Passing Address to a Function
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)

	fmt.Println("After :", text, age)

	// anonymous function
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	// Closures Function
	// Closures are a special case of anonymous functions.
	// Closures are anonymous functions
	// which access the variables defined outside the body of the function.
	l := 20
	b := 30

	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

	// Variadic Function
	fmt.Println(sum(1, 2, 3))

	// Select single argument from all arguments of variadic function
	variadicExample("red", "blue", "green", "yellow")

	// Passing multiple string arguments to a variadic function
	variadicExample1()
	variadicExample1("red", "blue")
	variadicExample1("red", "blue", "green")
	variadicExample1("red", "blue", "green", "yellow")

	//
	variadicExample2(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
