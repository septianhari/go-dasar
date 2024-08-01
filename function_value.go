package main

import "fmt"

func getGoodBye(name string) string {
	return "Good bye " + name
}

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("Hari"))
	fmt.Println(getHello("Hari"))
}
