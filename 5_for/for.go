package main

import "fmt"

func main() {
	// // for loop
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println("for loop", i)
	// }

	// // while loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println("while loop", i)
	// 	i++
	// }

	// // for loop with multiple conditions
	// for i, j := 0, 10; i < j; i, j = i+1, j-1 {
	// 	fmt.Println("for loop with multiple conditions", i, j)
	// }

	// range
	for i := range 3 {
		fmt.Println("range", i)
	}

}
