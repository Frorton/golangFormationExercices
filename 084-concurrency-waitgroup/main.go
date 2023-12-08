/*
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to
set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished.
Writing concurrent code: put “go” in front of a function or method call.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t", runtime.GOOS)
	lisy()
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	lisy()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	lisy()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	lisy()
	wg.Add(1)
	go foo()
	lisy()
	bar()
	lisy()
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	lisy()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	lisy()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo :", i)
	}
	wg.Done()
}

func bar() {
	for i := 10; i > 0; i-- {
		fmt.Println("bar :", i)
	}
}

func lisy() {
	fmt.Println("-------")
}
