package main

import "fmt"

func main() {

	// const pi float32 = 3.14
	// //pi = 3.14  (error)
	// fmt.Println(pi)

	// casting===========

	var age int8 = -10
	largeAge := uint8(age)
	fmt.Println(age, largeAge)
}
