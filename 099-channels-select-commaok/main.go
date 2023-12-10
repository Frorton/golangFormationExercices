package main

func main() {
	// Using bool
	/*
			func main() {
			even := make(chan int)
			odd := make(chan int)
			quit := make(chan bool)

			go send(even, odd, quit)

			receive(even, odd, quit)

			fmt.Println("about to exit")
		}

		// send channel
		func send(even, odd chan<- int, quit chan<- bool) {
			for i := 0; i < 100; i++ {
				if i%2 == 0 {
					even <- i
				} else {
					odd <- i
				}
			}
			close(quit)
		}

		// receive channel
		func receive(even, odd <-chan int, quit <-chan bool) {
			for {
				select {
				case v := <-even:
					fmt.Println("the value received from the even channel:", v)
				case v := <-odd:
					fmt.Println("the value received from the odd channel:", v)
				case i, ok := <-quit:
					if !ok {
						fmt.Println("from comma ok bit", i)
						return
					} else {
						fmt.Println("from comma ok bit", i)
					}
					or
				case _, _ = <-quit:
					return
					or
				case <-quit:
					return
				}
			}
		}
	*/

	// Using int
	/*
	   func main() {
	   	even := make(chan int)
	   	odd := make(chan int)
	   	quit := make(chan int)

	   	go send(even, odd, quit)

	   	receive(even, odd, quit)

	   	fmt.Println("about to exit")
	   }

	   // send channel
	   func send(even, odd, quit chan<- int) {
	   	for i := 0; i < 10; i++ {
	   		if i%2 == 0 {
	   			even <- i
	   		} else {
	   			odd <- i
	   		}
	   	}
	   	close(quit)
	   }

	   // receive channel
	   func receive(even, odd, quit <-chan int) {
	   	for {
	   		select {
	   		case v := <-even:
	   			fmt.Println("the value received from the even channel:", v)
	   		case v := <-odd:
	   			fmt.Println("the value received from the odd channel:", v)
	   		case i, ok := <-quit:
	   			if !ok {
	   				fmt.Println("from comma ok", i, ok)
	   				return
	   			} else {
	   				fmt.Println("from comma ok", i)
	   			}
	   		}
	   	}
	   }
	*/

	// Comma ok alone
	/*

	   func main() {
	   	c := make(chan int)
	   	go func() {
	   		c <- 42
	   		close(c)
	   	}()

	   	v, ok := <-c

	   	fmt.Println(v, ok)

	   	v, ok = <-c

	   	fmt.Println(v, ok)
	   }
	*/

}
