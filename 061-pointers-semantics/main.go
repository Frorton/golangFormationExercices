package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// pointer semantics
func addOneP(v *int) {
	*v += 1
}

func main() {

	// value semantics
	a := 1
	fmt.Println(a)         // 1
	fmt.Println(addOne(a)) // 2
	fmt.Println(a)         // 1 -> copies data to use with func addOne, value unchanged, not efficient but safe

	// pointer semantics
	b := 1
	fmt.Println("-------")
	fmt.Println(b) // 1
	addOneP(&b)
	fmt.Println(b) // 2 -> takes the adress of the value and changes it, more efficient but at risk of side effects

	// VALUE SEMANTICS : COPIED VALUES, if no need for memory management or modifying input, small data structures
	// POINTER SEMANTICS : SHARED MEMORY, large data structures, if needed to modify the receiver or input parameter, 64 bytes or +

	//"The majority of bugs that we get in software have to do with the mutation of memory" - Bill KENEDY

	// Reminder : KISS !
}
