package main

import "fmt"

func main() {
	rows := 5

	for i := rows; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
