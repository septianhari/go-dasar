package main

import "fmt"

func main() {
	// hasCreative := true
	// if hasCreative {
	// 	fmt.Println("I have creative")
	// }
	// fmt.Println("End")

	// grade := 0
	// if grade >= 80 {
	// 	fmt.Println("cakep")
	// } else if grade <= 80 {
	// 	fmt.Println("lumayan")
	// } else {
	// 	fmt.Println("hmmm")
	// }
	// fmt.Println("selesai")

	// Deklarasi variabel
	age := 15
	salary := 4000

	// Nested condition
	if age > 18 {
		fmt.Println("Usia Anda melebihi 18 tahun.")

		if salary > 2000 {
			fmt.Println("Gaji Anda melebihi 2000.")
		} else {
			fmt.Println("Gaji Anda tidak melebihi 2000.")
		}
	} else {
		fmt.Println("Usia Anda kurang dari atau sama dengan 18 tahun.")
	}
}
