package main

import "fmt"

func main() {
	c := make(chan int)

	/*
		c <- 21

		fmt.Println(<-c)
	*/
	// This won't run, blocked cuz needs to send & receive at the same time but we only receive
	// Like a relay (athleticism), we need something to launch the data : goroutines

	go func() {
		c <- 21
	}()

	fmt.Println(<-c)
	// Easy as that
	// The 2nd guy on a track field during a relay won't ever get to the finish line
	// if there's no one to run the first part of the race, until then he's just standing there
	// It's a channel block !

	// Can also create a buffer channel : allows a certain number of int the stay in the channel
	ch := make(chan int, 1)

	ch <- 12

	fmt.Println(<-ch)
	// It works because we allowed 1 value to stay in the channel
	// Won't work if we add on line 30
	// 	ch <- 13
	// Because we added ", 1" in make(chan int) not 2
	// Buffer channel have limits because of memory usage and reaching threshold
	// using goroutines is a better habit than using buffers
}
