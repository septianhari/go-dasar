package main

import "fmt"

type FilterNumber func([]int) ([]int, []int)

func main() {
	var filterGanjil FilterNumber = func(listNumber []int) (ganjil, genap []int) {
		for _, v := range listNumber {
			switch {
			case v%2 == 0:
				genap = append(genap, v)
			case v%2 == 1:
				ganjil = append(ganjil, v)
			}
		}
		return ganjil, genap
	}
	var listNumber = []int{1, 2, 3, 4, 5, 6, 7}
	var ganjil, genap = filterGanjil(listNumber)
	fmt.Println("Ganjil :", ganjil)
	fmt.Println("Genap :", genap)
}
