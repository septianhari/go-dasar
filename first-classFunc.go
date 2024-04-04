package main

import "fmt"

func inc(x int) int {
	return x + 1
}

func main() {
	dec := func(x int) int {
		return x - 1
	}
	fmt.Println(inc(1))
	fmt.Println(dec(3))
}
