package main

import "fmt"

func endApplication() {
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApplication()
	if error {
		panic("Error Application")
	}
	fmt.Println("Run Application")
}

func main() {
	app := true
	runApp(app)
}
