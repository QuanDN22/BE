package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Boolean
	var myBool bool = true
	fmt.Println(myBool)

	// String
	var myString string = "Quan"
	fmt.Println(myString)

	// int
	var myInt int = 5
	fmt.Println(myInt)

	// int 8 16 32 64
	// Range
	fmt.Println("Int8: ", math.MinInt8, math.MaxInt8, bits.OnesCount8(math.MaxUint8))
	fmt.Println("Int16: ", math.MinInt16, math.MaxInt16, math.MaxUint16)
	fmt.Println("Int32: ", math.MinInt32, math.MaxInt32, math.MaxUint32)
	fmt.Println("Int64: ", math.MinInt64, math.MaxInt64, bits.OnesCount64(math.MaxUint64))

	// uint 
	var myUint uint = 10
	fmt.Println(myUint)

}
