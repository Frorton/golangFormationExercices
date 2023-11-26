/*
A mutable value is a value that can be changed.
Slices, maps and pointers are mutable data types.
Even though they are passed by value, they behave as if they were passed by reference because
	the "value" that is passed is the reference to the underlying data, not the actual data.
*/

package main

import "fmt"

func intDelta(n *int) {
	*n = 43 // Changes the value stored at the adress*n
}

func sliceDelta(ii []int) {
	ii[0] = 99 // Changes the value at the 0 position of the ii[] slice
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33 // Changes the value of the int stored in the md map[string]int
}

func main() {
	a := 42
	fmt.Println(a) // Prints 42
	intDelta(&a)   // Mutating a to make it 43
	fmt.Println(a) // Prints 43

	xi := []int{1, 2, 3, 4}
	fmt.Println(xi) // Prints 1,2,3,4
	sliceDelta(xi)  // Mutating to make position 0 = 99
	fmt.Println(xi) // Prints 99,2,3,4

	m := make(map[string]int)
	m["James"] = 32
	fmt.Println(m["James"]) // Prints 32
	mapDelta(m, "James")    //Mutating "James" to make it 33
	fmt.Println(m["James"]) // Prints 33
}
