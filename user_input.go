package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Printf("Hello %s!\n", name)

	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Printf("You are %d years old.\n", age)

	//tambahkan prediksi umur untuk 5 tahun ke depan
	age += 5
	fmt.Printf("In 5 years, you will be %d years old.\n", age)

}
