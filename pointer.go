package main

import "fmt"

func main() {
	var number int
	number = 100
	var numberPointer *int
	numberPointer = &number

	fmt.Println(&number)
	fmt.Println(*numberPointer)
}
