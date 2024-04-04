package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("HelloWorld")
	jim := Person{firstName: "John", lastName: "John"}
	jim.UpdateName()
	jim.print()
}

func (p Person) print() {
	fmt.Println(p)
}

func (p *Person) UpdateName() {
	(*p).firstName = "Waker"
}
