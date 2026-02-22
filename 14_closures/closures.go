package main

import "fmt"

func count() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

func main() {
	increment := count()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
