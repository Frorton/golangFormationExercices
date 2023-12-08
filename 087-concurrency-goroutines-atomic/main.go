package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	var counter int64

	const geez = 100

	var wg sync.WaitGroup
	wg.Add(geez)

	for i := 0; i < geez; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Count :", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("count :", counter)
}
