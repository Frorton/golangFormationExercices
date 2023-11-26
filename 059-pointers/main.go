/*
& is an operator and gives the memory adress of the value stored in a variable
*/

package main

import "fmt"

func main() {

	x := 5
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("------- Types -------")

	y := 0.732
	z := "Golfing without drinking"
	fmt.Printf("%T\n", &x) // is of type pointer to an int
	fmt.Printf("%T\n", &y) // is of type pointer to a float64
	fmt.Printf("%T\n", &z) // is of type pointer to a string

	/*
		*any-type = pointer to said-type, it's different from a type
		ex : *int =/= int
	*/

	fmt.Println("------- Dereferencing -------")

	d := &x // d is of type pointer to an int

	fmt.Printf("%T\n", d)  // if you print the value of d, you get the adress of the value x
	fmt.Printf("%v\n", *d) // if you print the value of pointer d, you get the value of x

	*d = 2

	fmt.Printf("%v\n", *d)       // if you change the value of pointer d, you change the value of x
	fmt.Printf("%v %v\n", &x, x) // adress of &x doesnt change but the value does
}
