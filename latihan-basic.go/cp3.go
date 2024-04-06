package main

import "fmt"

// Tedapat sekolah yang memberikan predikat pada murid berdasarkan nilai rata-rata yang didapat :

// Jika nilai yang didapat adalah 100, murid mendapat predikat 'Sempurna'
// Jika nilai yang didapat 90 keatas, murid mendapat predikat 'Sangat Baik'
// Jika nilai yang didapat 80 keatas, murid mendapat predikat 'Baik'
// Jika nilai yang didapat 70 keatas, murid mendapat predikat 'Cukup'
// Jika nilai yang didapat 60 keatas, murid mendapat predikat 'Kurang'
// Jika nilai yang didapat kurang dari 60, murid mendapat predikat 'Sangat kurang'
// Murid mendapat 4 nilai dari 4 tugas yang berbeda, yaitu math, science, english, dan indonesia.

// Terdapat fungsi bernama GetPredicate yang menerima 4 parameter,
// yaitu math, science, english, dan indonesia bertipe int yang berisi nilai dari masing-masing mata pelajaran.
// Fungsi ini akan mengembalikan nilai bertipe string yang berisi predikat yang didapat murid.

func GetPredicate(math, science, english, indonesia int) string {
	if math == 100 && science == 100 && english == 100 && indonesia == 100 {
		return "Sempurna"
	} else if math >= 90 && science >= 90 && english >= 90 && indonesia >= 90 {
		return "Sangat Baik"
	} else if math >= 80 && science >= 80 && english >= 80 && indonesia >= 80 {
		return "Baik"
	} else if math >= 70 && science >= 70 && english >= 70 && indonesia >= 70 {
		return "Cukup"
	} else if math >= 60 && science >= 60 && english >= 60 && indonesia >= 60 {
		return "Sangat Kurang"
	} else {
		return "banyakin belajar"
	}
}

func main() {
	fmt.Println(GetPredicate(100, 100, 100, 100))
	fmt.Println(GetPredicate(90, 90, 90, 90))
	fmt.Println(GetPredicate(80, 80, 80, 80))
	fmt.Println(GetPredicate(70, 70, 70, 70))
	fmt.Println(GetPredicate(60, 60, 60, 60))
	fmt.Println(GetPredicate(20, 20, 20, 20))
}
