package main

import "fmt"

func countDown(n int) {
	//lebih simpel dari menggunakan loop
	fmt.Println(n)
	// base condition
	if n == 0 {
		return
	}
	countDown(n - 1)

	//infinity loop
	// for i := n; i >= 0; i-- {
	// 	fmt.Println(i)
	// }
}

func factorial(n int) int {
	if n == 1 {
		//base condition
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	countDown(3)

	fmt.Println(factorial(4))
}
