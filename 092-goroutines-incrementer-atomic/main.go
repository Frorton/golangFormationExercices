// Same base code as in 091
// Removed mutex part added atomic.AddInt64/LoadInt64
// Nees incrementer to become int64

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	var incrementer int64

	const max = 150

	var wg sync.WaitGroup
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			runtime.Gosched()
			fmt.Println("Count :", atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("incrementer :", incrementer)
}
