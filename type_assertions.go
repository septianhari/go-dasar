package main

import "fmt"

func random() any {
	return "ok"
}

func main() {

	var result any = random()

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}

// result := random()
// resultString := result.(string)
// fmt.Println(resultString)

// resultInt := result.(int) //panic
// fmt.Println(resultInt)
