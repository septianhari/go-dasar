package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")

	go func() {
		fmt.Println("Hello from new goroutine")
	}()

	fmt.Println("Before sleep")
	time.Sleep(10 * time.Millisecond) // menunggu 10 milli-second agar go routine berjalan
	fmt.Println("main stoped")

}
