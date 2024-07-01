package main

import "fmt"

func main() {
	hasEaten := !false
	hasStudy := !true
	//canPlayGame := hasEaten && hasStudy
	canPlayGame := hasEaten || hasStudy
	fmt.Println(canPlayGame)

}
