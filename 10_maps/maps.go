package main

import (
	"fmt"
)

func main() {
	// make a map
	m := map[string]int{"four": 4}

	// add values to map
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	v, ok := m["three"]
	fmt.Println(v, ok)
}
