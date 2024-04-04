package main

import "fmt"

//secara basic function = memiliki name-parameter dan return type

func add(a int, b int) int {
	//fmt.Println(a + b)
	return a + b
}

func sayHello(name string) {
	fmt.Printf("Hallo, %s\n", name)
}

func addNumbers(numbers ...int) int { //varidaic parameter
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// bagian-defer = mengakhirkan/menunda suatu ekseskusi dari func
func readFile(fileName string) {
	fmt.Println("Open file")
	defer fmt.Println("Close file")

	fileContent := "Hello"
	fmt.Println(fileContent)

	if fileContent == "Hello" {
		return
	}
}

func main() {

	readFile("test.txt")

	// a := 2
	// b := 6
	fmt.Println(addNumbers(10, 3)) //variadic parameter
	fmt.Println(addNumbers(10, 3, 1, 1))

	fmt.Println(add(2, 4))
	fmt.Println(add(1, 1))
	fmt.Println(add(6, 2))

	sayHello("moriis")
	sayHello("John")
	sayHello("Miu")

}
