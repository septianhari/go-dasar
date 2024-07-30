package main

import "fmt"

func main() {

	day := 9
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	fmt.Println("=========================")

	switch length := len("Helloo"); length > 5 {
	case true:
		fmt.Println("Longer than 5")
	case false:
		fmt.Println("Shorter than 5")
	}

	fmt.Println("=========================")

	switch {
	case day > 5:
		fmt.Println("Weekend")
	case day < 5:
		fmt.Println("Weekday")
	}
}
