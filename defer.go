package main

import "fmt"

func looging() {
	fmt.Println("Selesai")
}

func runApplication() {
	fmt.Println("Run Application")
	defer looging()
	fmt.Println("Start Application")
}

func main() {
	runApplication()

	fmt.Println("End Application")
}
