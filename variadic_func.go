package main

import (
	"fmt"
)

func sumAll(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2))
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumAll(number...))
}
