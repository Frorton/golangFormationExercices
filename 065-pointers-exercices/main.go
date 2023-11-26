/*
- Create a value and assign it to a variable.
- Print the adress of the value.
	-------
- Print the value stored in each of the values f,g,h,i
- Print the type of each value
- Print the data stored at memory locations :
	dereference the stored memory adress
*/

package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	f := "L'aïeul mourait froid et rigide."
	g := "Aux clairs matins de grand soleil"
	h := "Comme la vie est douce et brève !"
	i := 113
	a = &f
	b = &g
	c = &h
	d = &i
}

func main() {
	e := 21
	fmt.Printf("Value is %v\t Adress is %v\n", e, &e)

	fmt.Println("-------")

	fmt.Printf("Value is %v\t type is %T\n", a, a)
	fmt.Printf("Value is %v\t type is %T\n", b, b)
	fmt.Printf("Value is %v\t type is %T\n", c, c)
	fmt.Printf("Value is %v\t type is %T\n", d, d)
	fmt.Println(*a) //dereferencing
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
