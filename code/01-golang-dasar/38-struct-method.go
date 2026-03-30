package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func main() {
	bani := Customer{"Bani", "Cirebon", 23}

	bani.sayHello()
}
