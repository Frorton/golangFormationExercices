// Same base code as in 090

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	incrementer := 0

	const max = 150

	// Creating a variable mutex to lock down the readability of the incrementer
	var lock sync.Mutex

	var wg sync.WaitGroup
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			lock.Lock() // Locked
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			lock.Unlock() //Unlocked
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("incrementer :", incrementer)
}
