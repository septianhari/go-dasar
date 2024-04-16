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

	time.Sleep(10 * time.Millisecond)

	fmt.Println("Main thread ended")
}
