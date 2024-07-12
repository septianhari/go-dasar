package main

import (
	"fmt"
	"time"
)

func greet(str string) {
	fmt.Println("Hi", str)
}

func main() {
	fmt.Println("Main thread started")

	go greet("Adit")
	go greet("Levi")

	fmt.Println("Before sleep")

	time.Sleep(5 * time.Millisecond)

	time.Sleep(10 * time.Millisecond)

	fmt.Println("Main thread ended")
}
