package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Babi" {
		return "***"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Budi", spamFilter)
}
