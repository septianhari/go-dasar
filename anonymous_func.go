package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("you are Welcome", name)
	}

}

func main() {

	registerUser("Har", func(name string) bool {
		return name == "Babi"
	})

}
