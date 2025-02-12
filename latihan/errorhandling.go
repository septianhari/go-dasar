package main

import (
	"fmt"
	"strconv"
)

func ErrorFunc(str interface{}) int {
	angka, err := strconv.Atoi(str.(string))
	if err != nil {
		panic(err)
	}
	return angka
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi error : ", r)
		}
	}()

	angka := ErrorFunc("SABAR")
	fmt.Println(angka)
}
