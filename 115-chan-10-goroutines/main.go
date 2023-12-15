package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- i
		}()
		close(ch)
	}
	for x := 0; x > 100; x++ {
		fmt.Println(x, <-ch)
	}
}
