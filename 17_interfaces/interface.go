package main

import "fmt"

type paymnent struct {
}

func (p *paymnent) makePayment(amount float32) {
	razorpayGateway := razorpay{}
	razorpayGateway.pay(amount)

}

type razorpay struct {
}

func (r *razorpay) pay(amount float32) {
	fmt.Println("Razorpay", amount)
}

func main() {
	newPayment := paymnent{}
	newPayment.makePayment(100)

}
