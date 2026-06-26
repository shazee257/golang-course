package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseconds precision
	customer  customer
}

func (o *Order) updateStatus(status string) {
	o.status = status
}

func (o *Order) getStatus() string {
	return o.status
}

func newOrder(id string, amount float32, status string) *Order {
	return &Order{
		id:     id,
		amount: amount,
		status: status,
	}
}

func main() {
	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{name: "Go", isGood: true}
	// fmt.Println("language: ", language)
	// fmt.Printf("%+v\n", language) // %+v prints the struct with field names

	// newCustomer := customer{
	// 	name:  "John Doe",
	// 	phone: "123-456-7890",
	// }

	manualOrder := Order{
		id:     "123",
		amount: 100.0,
		status: "pending",
		// customer: newCustomer,
		customer: customer{
			name:  "John Doe",
			phone: "123-456-7890",
		},
	}

	fmt.Println("manualOrder: ", manualOrder)

	order := newOrder("123", 100.0, "pending")
	fmt.Println("order: ", order)

	order.createdAt = time.Now()

	order.updateStatus("completed")
	fmt.Println("order: ", order)

	fmt.Println("Status: ", order.getStatus())
}
