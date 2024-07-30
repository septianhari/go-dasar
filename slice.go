package main

import "fmt"

func main() {
	names := [...]string{"Har", "Bobi", "Rudi", "Hari"}

	slice1 := names[1:3]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	fmt.Println("=================================")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] //sabtu - minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Hari Baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("=================================")

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Har"
	newSlice[1] = "Har"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}
