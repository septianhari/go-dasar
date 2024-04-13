package main

import "fmt"

func swap(a, b *int) {
	c := *a
	*a = *b
	*b = c

	fmt.Println("Di dalam swap", *a, *b)
}

func main() {
	x := 1
	y := 2
	fmt.Println("Sebelum swap", x, y)
	swap(&x, &y)
	fmt.Println("Setelah swap", x, y)

	// 	var number int
	// 	number = 100
	// 	var numberPointer *int
	// 	numberPointer = &number

	// fmt.Println(&number)
	// fmt.Println(*numberPointer)
}
