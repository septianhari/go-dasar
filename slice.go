package main

import "fmt"

func main() {
	//buat slice array
	// names := [5]string{"Har", "Bobi", "Rudi", "Gibran", "Owi"}
	// fmt.Println(names)

	// bestNames := names[0:3]
	// fmt.Println(bestNames)

	names := []string{"Har", "Bobi", "Rudi", "Gibran", "Owi"}
	fmt.Println(names)

	//pake make
	newNames := make([]string, 3)
	newNames[0] = "Har"
	newNames[1] = "Bobi"
	newNames[2] = "Rudi"
	newNames = append(newNames, "Owi")
	fmt.Println(newNames)

	//ubah yang Bobi
	for i, name := range names {
		if name == "Bobi" {
			names[i] = "Har"
		} else {
			fmt.Println(name)
		}
	}

}
