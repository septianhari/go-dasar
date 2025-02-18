package main

import (
	"fmt"
)

func main() {
	var x uint64 = 10
	y := &x
	z := x

	i := 10

	for {
		if i < 1 {
			break
		}
		*y += uint64(12)
		i--
	}

	fmt.Println("x", x)
	fmt.Println("y", *y)
	fmt.Println("z", z)
}
