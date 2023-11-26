package main

import "fmt"

func main() {

	f := outer()
	fmt.Println(f())
}

//Outer returns a func that returns an int
func outer() func() int {
	return func() int {
		return 21
	}
}
