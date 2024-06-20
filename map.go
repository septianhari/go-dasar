package main

import "fmt"


scoreMap := map[string]string {
    "A": "100-86",
    "B": "75-86",
    "C": "50-74",
    "D": "30-49",
    "E": "0-29",
}

score := scoreMap["A"]
fmt.Println(score)

func transtlateIdToEN(word string) string {
	// indonesia := []string{"satu", "dua", "tiga"}
	// english := []string{"one", "two", "three"}

	// for i, id := range indonesia {
	// 	if id == word {
	// 		return english[i]
	// 	}
	// }
	// return ""

	idToEn := map[string]string{

		"satu": "one",
		"dua":  "two",
		"tiga": "three",
	}
	return idToEn[word]
}

func main() {
	//fmt.Println(transtlateIdToEN("satu"))
	fmt.Println(transtlateIdToEN("tiga"))
}
