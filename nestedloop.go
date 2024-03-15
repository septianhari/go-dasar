package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}
