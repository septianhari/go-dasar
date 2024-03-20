package main

import "fmt"

func main() {
	var number = 10
	var pointer *int = &number // Mengambil alamat dari variabel number

	*pointer = 20 // Mengubah nilai yang ditunjuk oleh pointer

	fmt.Println("number:", number)
	fmt.Println("pointer:", *pointer) // Mencetak nilai yang ditunjuk oleh pointer
}
