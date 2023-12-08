/*
Race conditions occur when multiple goroutines access a shared variable.
Mutex = like a book in a library, one at a time !
*/

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
	var mu sync.Mutex

	for i := 0; i < geez; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count :", counter)
}
