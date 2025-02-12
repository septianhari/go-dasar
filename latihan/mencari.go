package main

import (
	"fmt"
)

func Operation(n int) int {
	if n == 1 {
		return n
	} else {
		return n + Operation(n-1)
	}
}

func main() {
	fmt.Println(Operation(10))
}
