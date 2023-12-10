/*
Channels can be unidirectional :
	- Only receive : make (<-chan int)
	- Only send : make (chan<- int)
Channels can be bidirectional :
	- Send & receive : make (chan int)
Bidirectionnal channels can be assigned to unidirectionnal channels, not the opposite :
	cha = ch won't work
	cha = c won't work
	ch = c won't work
	c = cha will work
	ch = cha will work
*/

package main

import "fmt"

func main() {

	cha := make(chan int)

	cha <- 23

	fmt.Println(<-cha)

	ch := make(<-chan int, 2)

	// ch <- 12 // Can only receive, can't send ch <- is impossible

	fmt.Println(<-ch) // Receiving works

	c := make(chan<- int, 1)

	c <- 21 // Sending works

	// fmt.Println(<-c) // Can only send, can't receive <-c is impossible
}
