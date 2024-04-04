package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	fmt.Println("Starting")
	t := triangle{3, 5}
	s := square{4}

	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (tr triangle) getArea() float64 {
	return tr.height * tr.base * 0.5
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
