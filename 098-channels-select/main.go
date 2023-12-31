package main

import "fmt"

func main() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// Send
	go send(even, odd, quit)

	// Receive
	receive(even, odd, quit)
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even channel:", v)
		case v := <-o:
			fmt.Println("From odd channel:", v)
		case v := <-q:
			fmt.Println("From quit channel:", v)
			return
		}
	}
}

// If you receive from a closed channel == 0
// "After the last value has been received from a closed channel c,
// 		any receive from c will succeed without blocking,
// 		returning the zero value for the channel element."
