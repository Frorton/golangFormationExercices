// https://go.dev/doc/tutorial/generics

package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

// Generic constraint : accepts either types in [T x|x]
func addT[T int | float64](a, b T) T {
	return a + b
}

func main() {
	// Normal
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))
	// Generic constraint
	fmt.Println(addT(1, 2))
	fmt.Println(addT(1.2, 2.2))
	fmt.Println(addT(2, 2.2))
}
