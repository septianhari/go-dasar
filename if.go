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

	fmt.Println("=============================")

	var point = 10
	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
