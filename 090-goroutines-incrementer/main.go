package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	incrementer := 0

	const max = 150

	var wg sync.WaitGroup
	wg.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			wg.Done()
		}()
		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("incrementer :", incrementer)
}

// Created a race condition, incrementer =/= 150 sometimes
// Concurrency can be unexpected so sometimes incrementer may be = 150
// Use go run -race main.go to cleanly run the program or use mutex as seen in 086
