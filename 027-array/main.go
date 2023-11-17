/*
Using a COMPOSITE LITERAL:
- Create an ARRAY which holds 5 VALUES of TYPE int.
- Assign VALUES to each index position.
- Range over the array and print the values out.
- Print out the VALUE and the TYPE.
*/

package main

import "fmt"

func main() {

	// create an array
	x := [5]int{}

	// assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	// print out
	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}

}
