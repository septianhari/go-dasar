package main

import "fmt"

func main() {

	var students [2][3]string
	students[0][0] = "Har"
	students[0][1] = "Bobi"
	students[0][2] = "Rudi"
	students[1][0] = "Har"
	students[1][1] = "Bobi"
	students[1][2] = "Rudi"

	fmt.Println(students[0][0])

	fmt.Println(students[1][2])

	fmt.Println(students[0])

	fmt.Println(students[1])

	fmt.Println(students[0][0], students[0][1], students[0][2])

	fmt.Println(students)

	var names [3]string
	names[0] = "Har"
	names[1] = "Bobi"
	names[2] = "Rudi"

	for i, name := range names {
		fmt.Println(i, name)
	}

	// var students [4]string
	// students[0] = "Har"
	// students[1] = "Bobi"
	// students[2] = "Rudi"

	// fmt.Println(students)

	// var numbers [6]int (//arr 1 dimensi)
	// for i := 0; i <= 5; i++ {
	// 	numbers[i] = i
	// }
	// fmt.Println(numbers)
	//====================================
	// students := [5]string{"Har", "Bobi", "Rudi"}

	// fmt.Println(students)
	//====================================
	// var student1, student2, student3 string
	// student1 = "Har"
	// student2 = "Bobi"
	// student3 = "Rudi"

	// fmt.Println(student1, student2, student3)
	//======================================
	// var students [3]string
	// students[0] = "Har"
	// students[1] = "Bobi"
	// students[2] = "Rudi"

	// fmt.Println(students[0], students[1], students[2])

}
