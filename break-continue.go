package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("Perulangan ke - ", i)
	}

	fmt.Println("==========")

	for j := 0; j < 15; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke - ", j)
	}
}
