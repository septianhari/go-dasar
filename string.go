package main

import "fmt"

func main() {

	var kata = "hallo Sabarno"
	fmt.Println(kata)
	fmt.Println(panjangKata(kata))

}

func panjangKata(kata string) int {
	return len(kata)
}
