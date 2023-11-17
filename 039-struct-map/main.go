/*
Take the code from the previous exercise.
- Store the VALUES of type person in a map with the KEY of last name.
- Access each value in the map.
- Print out the values.
- Range over the slice.
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

	fmt.Println("-------") // L'exercice commence ici

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.ficf {
			fmt.Println(v.first, v.last, v2)
		}
	}
}
