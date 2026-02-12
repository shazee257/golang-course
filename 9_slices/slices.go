package main

import "fmt"

func main() {
	var nums = make([]int, 5)
	fmt.Println(nums)

	// uninitialized slice
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// nums = append(nums, 1)
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(cap(nums))

	// // slices
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println(slice)

	// // slice with range
	// for i, v := range slice {
	// 	fmt.Println(i, v)
	// }

	// // slice with for loop
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(i, slice[i])
	// }
}
