package main

import "fmt"

func transtlateIdToEN(word string) string {
	indonesia := []string{"satu", "dua", "tiga"}
	english := []string{"one", "two", "three"}

	for i, id := range indonesia {
		if id == word {
			return english[i]
		}
	}
	return ""
	// idToEn := map[string]string{

	// 	"satu": "one",
	// 	"dua":  "two",
	// 	"tiga": "three",
	// }
	// return idToEn[word]
}

func main() {
	fmt.Println(transtlateIdToEN("satu"))
}
