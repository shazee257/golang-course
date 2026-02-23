package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func (o *Order) updateStatus(status string) {
	o.status = status
}

func (o *Order) getStatus() string {
	return o.status
}

func main() {
	order := Order{
		id:     "1",
		amount: 10.99,
		status: "pending",
	}

	order.createdAt = time.Now()

	order.updateStatus("completed")
	fmt.Println("order: ", order)

	fmt.Println("Status: ", order.getStatus())
}
