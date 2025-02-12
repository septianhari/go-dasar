package modul

import "fmt"

func functionModul() {
	fmt.Println("Function dari modul")
}

package main

import ("ramdan/modul")

func main() {
	modul.functionModul()
}