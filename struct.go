package main

import "fmt"

type Student struct {
	Name    string
	Age     int
	Address string
}

func (s Student) SayHello() {
	fmt.Println("Hello", s.Name)
}

func (s Student) SayAge() {
	fmt.Println("Age", s.Age)
}

func (s Student) SayAddress() {
	fmt.Println("Address", s.Address)
}

func (s Student) SayName() {
	fmt.Println("Name", s.Name)
}

func main() {

	// var student Student
	var student Student
	student.Name = "Har"
	student.Age = 22
	student.Address = "Jakarta"
	fmt.Println(student)

	student.SayHello()
	student.SayAge()
	student.SayAddress()

	student.SayName()
}
