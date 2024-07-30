package main

import "fmt"

func main() {

	if name := "Jokoy"; name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Har" {
		fmt.Println("Hello Har")
	} else {
		fmt.Println("Hi, unknown")
	}

	fmt.Println("=============================")

	name := "Saeful"
	if length := len(name); length > 5 {
		fmt.Println("Nama kepanjang")
	} else {
		fmt.Println("Nama kependek")
	}
}
