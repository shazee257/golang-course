package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x int = 5
	fmt.Println(x)

	name := "Shahzad"
	fmt.Println(name)

	// float
	price := 10.2
	fmt.Println(unsafe.Sizeof(price))

	// var isStudent bool = true
	// fmt.Println(isStudent)

	// var height float64 = 5.9
	// fmt.Println(height)
}
