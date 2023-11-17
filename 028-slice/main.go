/*
Using a COMPOSITE LITERAL:
- Create a SLICE of TYPE int.
- Assign these 10 VALUES :
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51
- Range over the slice and print the values out.
- Print out the VALUE and the TYPE.
*/

package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range x {
		fmt.Printf("%v \t %T \t %v\n", v, v, i)
	}

	fmt.Printf("%T \t %#v\n", x, x)
	fmt.Printf("%T \t %v\n", x, x)

}
