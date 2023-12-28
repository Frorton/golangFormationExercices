package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("ERROR !! ERROR !! : %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "Needs more sleep",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("foo fell asleep - ", e)
}
