package main

import "fmt"

func main() {
	// arrays
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// array with short statement
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	// array with range
	for i, v := range arr2 {
		fmt.Println(i, v)
	}

	// array with for loop
	for i := 0; i < len(arr2); i++ {
		fmt.Println(i, arr2[i])
	}
}
