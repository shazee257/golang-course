package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", nums)

	for i, v := range nums {
		fmt.Println(i, v)
	}
}
