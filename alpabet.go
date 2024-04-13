package main

import "fmt"

func main() {
	// Membuat variabel string untuk menyimpan alphabet
	var alphabet string

	// Loop dari A sampai Z menggunakan ASCII code
	for i := 'A'; i <= 'Z'; i++ {
		// Mengkonversi ASCII code ke karakter dan menambahkannya ke variabel alphabet
		alphabet += string(i)
	}

	// Menampilkan alphabet dari A sampai Z
	fmt.Println("Alphabet: ", alphabet)
}
