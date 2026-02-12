package main

import "fmt"

const name = "Shahzad"

func main() {
	const (
		port int    = 8080
		host string = "localhost"
	)
	fmt.Println(port)
	fmt.Println(host)

	fmt.Println(name)
}
