package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myType interface {
	~int | ~float64 // ~ take into account all values of type in (including myAlias)
}

func addT[T myType](a, b T) T {
	return a + b
}

// Underlying type constraint
type myAlias int

func main() {
	// Normal
	fmt.Println(addI(1, 2))
	fmt.Println(addF(1.2, 2.2))

	var x myAlias = 5

	fmt.Println(addT(1, x))
	fmt.Println(addT(x, 22))
}
