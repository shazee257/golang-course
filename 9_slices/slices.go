package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	nums2 := make([]int, 3)

	// nums2 := nums
	copy(nums2, nums)
	// print address of nums
	fmt.Printf("%p\n", nums)
	fmt.Printf("%p\n", nums2)

	i := 2
	// print type of i
	fmt.Printf("%T\n", i)

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
