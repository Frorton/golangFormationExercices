package main

import "fmt"

func main() {
	c := make(chan int)

	// Send
	go foo(c)

	// Receive
	bar(c)
}

// Send
func foo(c chan<- int) {
	c <- 21
}

// Receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
