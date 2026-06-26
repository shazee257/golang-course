package main

import "fmt"

type paymenter interface {
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Razorpay", amount)
}

func (p payment) makePayment(amount float32) {
	razorpayGateway := razorpay{}
	razorpayGateway.pay(amount)
}

func main() {
	razorpayGateway := razorpay{}
	newPayment := payment{
		gateway: razorpayGateway,
	}
	newPayment.makePayment(100)
}
