package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("changeNum: ", *num)
}

func main() {
	num := 10
	changeNum(&num)
	fmt.Println("main: ", num)

}
