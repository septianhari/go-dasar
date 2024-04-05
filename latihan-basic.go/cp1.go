package main

import "fmt"

func GraduateStudent(score int, absent int) {
	// score := 75
	// absent := 6

	if score >= 70 && absent <= 5 {
		fmt.Println("lulus")
	} else if score <= 70 || absent >= 5 {
		fmt.Println("tidak lulus")
	}
}

func main() {
	GraduateStudent(75, 3)
}
