package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func add(x, y, z int) int {
	return x + y + z
}

func main() {
	sayHello()

	fmt.Println(add(1, 2, 6))
	fmt.Println(add(10, 24, 73))
}
