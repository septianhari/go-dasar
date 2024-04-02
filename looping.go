package main

import "fmt"

func main() {

	// nested loop
	// Mendefinisikan variabel untuk jumlah baris dan kolom
	rows := 1
	cols := 5

	// Nested loop untuk mencetak pola
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print("* ")
		}
		fmt.Println() // Pindah ke baris berikutnya setelah mencetak kolom selesai
	}

	// usia := 0
	// for {
	// 	fmt.Print("Masukan Usia: ")
	// 	fmt.Scan(&usia)
	// 	if usia > 150 {
	// 		fmt.Println("Usia tidak valid")
	// 	} else {
	// 		break
	// 	}
	// }
	// fmt.Println("Usia Anda : ", usia)

	//i := 1
	// for i := 1; i <= 20; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("End looping")

	// for i := 1; i < 5; i++ {
	// 	fmt.Println(i)
	// }
	// fmt.Println("End looping")

	// i := 1
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println("End looping")

	// i := 1 (infinity loop)
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }
}
