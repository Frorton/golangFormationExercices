package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := rand.Intn(250)
	fmt.Printf("x is of value %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("x is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("x is between 101 and 200")
	default:
		fmt.Println("x is between 201 and 250")
	}
}
