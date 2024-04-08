package main

import "fmt"

/*Jika Umur di bawah 5 tahun maka tidak dapat membeli tiket wahana dengan menampilkan -1 (repersentasi dari error)
Jika anak umur 5-7 tahun atau tinggi lebih dari 120 cm maka akan dikenakan tarif Rp 15.000.
Jika anak umur 8-9 tahun atau tinggi lebih dari 135 cm maka akan dikenakan tarif Rp 25.000.
Jika anak umur 10-11 tahun atau tinggi lebih dari 150 cm maka akan dikenakan tarif Rp 40.000.
Jika anak umur 12 tahun atau tinggi lebih dari 160 cm, akan dikenakan tarif Rp 60.000.
Jika di atas 12 tahun, akan mendapat tiket khusus Remaja yaitu sebesar Rp 100.000*/

func TicketPlayground(height, age int) int {
	//var price int

	if age < 5 {
		return -1
	}
	var price = 10000

	if age >= 5 || height > 120 {
		price += 5000
	}

	if age >= 8 || height > 135 {
		price += 10000
	}

	if age >= 10 || height > 150 {
		price += 15000
	}

	if age >= 12 || height > 160 {
		price += 20000
	} else {
		price = 100000
	}

	return price // TODO: replace this
}

func main() {
	fmt.Println(TicketPlayground(120, 10))
}
