package main

import "fmt"

func main() {
	var student1, student2, student3 string
	student1 = "Har"
	student2 = "Bobi"
	student3 = "Rudi"

	fmt.Println(student1, student2, student3)

	var students [3]string
	students[0] = "Har"
	students[1] = "Bobi"
	students[2] = "Rudi"

	fmt.Println(students[0], students[1], students[2])

}
