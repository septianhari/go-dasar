package main

import (
	"fmt"
)

func endApplication() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println("Error", message)
}

func runApplication(error bool) {
	defer endApplication()
	if error {
		panic("Error Application")
	}
	fmt.Println("Run Application")

}

func main() {
	runApplication(true)
	fmt.Println("Septian")
}
