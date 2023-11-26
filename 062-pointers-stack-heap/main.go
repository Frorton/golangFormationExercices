package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)  // Value semantics
	fmt.Println(&a) // Pointer semantics
}

/*
Line 7 : escapes to the heap
variable shared between main() and Println()

Line 8 : moved to heap
variable moved to the heap
*/
