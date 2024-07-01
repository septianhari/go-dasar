package main

import (
	"fmt"
)

func People() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
}

func main() {
	for i := 1; i <= 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
}
