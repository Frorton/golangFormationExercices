package main

import "fmt"

func main() {
	cs := make(chan int) // Removed <-
	// to make the channel bidirectionnal so we can send/receive

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Println("------")
	fmt.Printf("cs\t%T\n", cs)

	cr := make(chan int) // Removed <-
	// to make the channel bidirectionnal so we can send/receive

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Println("------")
	fmt.Printf("cr\t%T\n", cr)
}
