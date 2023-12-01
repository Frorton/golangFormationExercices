package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// Generic type set : accepts either types using x | x
type myNumber interface {
	int | float64
}

func addT[T myNumber](a, b T) T {
	return a + b
}

func main() {
	// Normal
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	// Generic type set
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))
	fmt.Println(addT(2, 2.2))
}
