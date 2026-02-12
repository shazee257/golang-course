package main

import "fmt"

func main() {
	// num := 10
	// if num > 5 {
	// 	fmt.Println("num is greater than 5")
	// } else {
	// 	fmt.Println("num is less than 5")
	// }

	// if with short statement
	if num := 10; num > 5 {
		fmt.Println("num is greater than 5")
	} else {
		fmt.Println("num is less than 5")
	}
}
