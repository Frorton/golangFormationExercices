package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
	// remember when imports go "No Required Module Provides Package" etc
	// 		➡ go mod tidy
)

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNums interface {
	// All the ints and floats, more info : https://pkg.go.dev/golang.org/x/exp/constraints
	constraints.Integer | constraints.Float
}

func addT[T myNums](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 22))
}
