/*
A wrapper func wraps or modifies another function's behavior,
while a callback func is used as an argument to be executed at a specific point or event.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

// doWork is callback of timeFunc
func doWork() {
	for i := 0; i < 2_000; i++ {
		fmt.Println(i)
	}
}

// timeFunc is wrapper of f()
func timeFunc(f func()) {
	start := time.Now()
	f() // here
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
