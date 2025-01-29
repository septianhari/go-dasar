package main

import (
	"fmt"
)

func main() {

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	var correctPassword string = "juned321"
	var input string

	for input != correctPassword {
		fmt.Print("Enter password: ")
		fmt.Scanln(&input)

		if input != correctPassword {
			fmt.Println("Incorrect password!")
		}
	}

	fmt.Println("Correct password!")

}
