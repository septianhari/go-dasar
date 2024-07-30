package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("done")

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("muter ke", counter)
	}

	fmt.Println("===========================")

	names := []string{"Har", "Bobi", "Rudi"}

	for _, name := range names {
		fmt.Println(name)
	}

}
