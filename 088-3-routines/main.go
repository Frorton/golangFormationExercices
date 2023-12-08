/*
In addition to the main goroutine, launch two additional goroutines
- Each additional goroutine should print something out
- Use waitgroups to make sure each goroutine finishes before your program exist
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 25; i > 0; i-- {
		fmt.Println("main :", i)
	}
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo :", i)
	}
	wg.Done()
}

func bar() {
	for i := 5; i > 0; i-- {
		fmt.Println("bar :", i)
	}
	wg.Done()
}
