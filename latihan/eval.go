package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	fmt.Println("Start")
	start := time.Now()

	time.AfterFunc(1*time.Second, func() {
		duration := time.Since(start)
		fmt.Println("Done", duration)
		wg.Done()
	})

	wg.Wait()

	fmt.Println("program closed")
}
