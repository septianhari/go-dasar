package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "John",
		"address": "123 Main St",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println("=========================")

	book := make(map[string]string)
	book["title"] = "Harry Potter"
	book["author"] = "J.K. Rowling"
	book["year"] = "1997"
	delete(book, "year")
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "year")
	fmt.Println(book)
	fmt.Println(len(book))
}
