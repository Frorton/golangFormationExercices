/*
Create a type “person” which will have an underlying type of “struct” so that it can store the following data:
	first name
	last name
	favorite ice cream flavors
- Create two VALUES of TYPE person.
- Print out the values.
- Range over the elements in the slice which stores the favorite flavors.
*/

package main

import "fmt"

type person struct {
	first string
	last  string
	ficf  []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		ficf:  []string{"Chocolate", "Mango", "Passion fruit", "Guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		ficf:  []string{"Strawberry", "Chocolate"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p2.ficf)
	fmt.Println(p1.ficf)

	for _, v := range p1.ficf {
		fmt.Println(p1.first, "favorite is", v)
	}

	for _, v := range p2.ficf {
		fmt.Println(p2.first, "favorite is", v)
	}
}
