package main

import "fmt"

func main() {
	var point int

	fmt.Print("Input bilanganmu : ")
	fmt.Scanln(&point)

	// if point > 75 {
	// 	fmt.Println("Selamat, kamu lulus")
	// } else {
	// 	fmt.Println("Maaf, kamu tidak lulus")
	// }

	if 100 >= point && point >= 90 {
		fmt.Println("Selamat, kamu mendapatkan nilai A")
	} else if 90 > point && point >= 80 {
		fmt.Println("Selamat, kamu mendapatkan nilai B")
	} else if 80 > point && point >= 70 {
		fmt.Println("Selamat, kamu mendapatkan nilai C")
	} else if 70 > point && point >= 60 {
		fmt.Println("Selamat, kamu mendapatkan nilai D")
	} else if 60 > point && point >= 0 {
		fmt.Println("Maaf, kamu mendapatkan nilai E")
	} else {
		fmt.Println("Maaf, nilai yang kamu masukkan tidak valid")
	}
}
