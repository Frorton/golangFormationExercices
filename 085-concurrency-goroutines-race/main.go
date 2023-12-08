package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	counter := 0

	const geez = 100

	var wg sync.WaitGroup
	wg.Add(geez)

	for i := 0; i < geez; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count :", counter)
}
