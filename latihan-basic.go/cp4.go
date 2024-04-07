package main

import "fmt"

/*
Jika tanggal pembelian ticket tersebut adalah tanggal ganjil
Jika ticket yang dibeli kurang dari 5 ticket, maka mendapat diskon 15% dari total pembelian
Jika ticket yang dibeli minimal 5 ticket keatas, maka mendapat diskon 25% dari total pembelian
Jika tanggal pembelian ticket tersebut adalah tanggal genap
Jika ticket yang dibeli kurang dari 5 ticket, maka mendapat diskon 10% dari total pembelian
Jika ticket yang dibeli minimal 5 ticket keatas, maka mendapat diskon 20% dari total pembelian
*/

func GetTicketPrice(vip int, regular int, student int, day int) float32 {
	var totalTicket = vip + regular + student
	var totalPrice = (vip * 30) + (regular * 20) + (student * 10)

	if totalPrice >= 100 {
		if day%2 != 0 {
			if totalTicket < 5 {
				return float32(totalPrice) - (float32(totalPrice) * 0.15)
			} else {
				return float32(totalPrice) - (float32(totalPrice) * 0.25)
			}

		} else {
			if totalTicket < 5 {
				return float32(totalPrice) - (float32(totalPrice) * 0.10)
			} else {
				return float32(totalPrice) - (float32(totalPrice) * 0.20)
			}
		}
	} else {
		return float32(totalPrice)
	}
}
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
	fmt.Println(GetTicketPrice(6, 6, 6, 21))
	fmt.Println(GetTicketPrice(3, 3, 3, 27))

}
