package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

// interface kosong
func Sayable() any {
	return 1
}

func main() {

	var kosong any = Sayable()
	fmt.Println(kosong)

	person := Person{Name: "Har"}

	sayHello(person)

	animal := Animal{Name: "Kucing"}
	sayHello(animal)
}
