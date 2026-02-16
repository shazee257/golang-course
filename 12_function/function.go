package main

import "fmt"

// func getLanguages() (string, string, string) {
// 	return "Go", "Python", "Java"
// }

// func add(a int, b int) int {
// 	return a + b
// }

func processIt(fn func(a int) int) int {
	return fn(10)
}

func main() {
	fn := func(a int) int {
		return a + 10
	}

	fmt.Println(processIt(fn))

	// fmt.Println(getLanguages())
}
