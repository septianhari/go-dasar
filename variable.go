package main

import "fmt"

func main() {

	// var age uint
	// age = 30
	// fmt.Println(age)

	nama, umur, negara := "Har", 22, "Indonesia"
	fmt.Println(nama)
	fmt.Println(umur)
	fmt.Println(negara)

	var isOnline bool

	isOnline = true
	fmt.Println(isOnline)

	fmt.Println("========================")

	var (
		name  string
		usia  int
		isGod bool
	)

	name = "Bobi"
	usia = 32
	isGod = true

	fmt.Println(name, usia, isGod)

}
