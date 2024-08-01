package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var customer Customer
	customer.Name = "Har"
	customer.Address = "Jakarta"
	customer.Age = 22
	fmt.Println(customer)

	fmt.Println(customer.Name)

	Atun := Customer{
		Name:    "Atun",
		Address: "Jakarta",
		Age:     22,
	}
	fmt.Println(Atun)

	Atun2 := Customer{"Atun", "Jakarta", 22}
	fmt.Println(Atun2)

	Har := Customer{"Har", "Jakarta", 22}
	fmt.Println(Har)

	Atun2.sayHi("Hari")
	Har.sayHi("Syifa")
}
