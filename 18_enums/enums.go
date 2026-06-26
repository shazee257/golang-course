package main

import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota
	Shipped
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Status:", status)
}

func main() {
	changeOrderStatus(Delivered)
}
