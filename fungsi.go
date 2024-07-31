package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

// return multiple value
func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

// named return values

func main() {
	result := getHello("Bagas")
	fmt.Println(result)

	var firstName = "Otong"
	var lastName = "Surotong"
	result = getFullName(firstName, lastName)
	fmt.Println(result)
}
