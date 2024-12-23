package main

import (
	"fmt"
)

func main() {

	var number int8 = 10
	number2 := 111
	number = 20

	fmt.Println(number2)
	fmt.Println(number)

	//var name string

	// var name = "Har Sabarno"
	// fmt.Println(name)

	name := "Har Septian"
	fmt.Println(name)

	name = "Hari Sabarno"
	fmt.Println(name)

	age := 22
	fmt.Println(age)

	address := "Jakarta"
	fmt.Println(address)

	pariabel()
}

func pariabel() {
	var (
		nama   = "Hari"
		usia   = 16
		alamat = "Malang"
	)

	fmt.Println(nama, usia, alamat)

}
