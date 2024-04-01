package main

import "fmt"

func main() {
	day := 13

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekdays")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Nothing days")

	}

	// switch day {
	// case 1:
	// 	fmt.Println("Senin")
	// case 2:
	// 	fmt.Println("Selasa")
	// case 3:
	// 	fmt.Println("Rabu")
	// case 4:
	// 	fmt.Println("Kamis")
	// case 5:
	// 	fmt.Println("Jumat")
	// case 6:
	// 	fmt.Println("Sabtu")
	// case 7:
	// 	fmt.Println("Minggu")

	// }

}
