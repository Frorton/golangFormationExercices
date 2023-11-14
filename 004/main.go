/*
var for zero value
short declaration operator
multiple initializations
var when specificity is required
blank identifier
*/
package main

import "fmt"

var x int

func main() {
	fmt.Println(x)

	y := 42
	fmt.Println(y)

	a, b := 43, "Vinzou"
	fmt.Println(a, b)

	var z float32 = 42.42
	fmt.Printf("%v is of this type %T\n", z, z)

	c, d, _ := 45, 46, 47
	fmt.Println(c, d)

}
