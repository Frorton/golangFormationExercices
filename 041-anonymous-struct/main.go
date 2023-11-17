/*
Create and use an anonymous struct with these fields:
	first     string
	friends   map[string]int
	favDrinks []string
- Print things.
*/

package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Jenny": 27,
			"Q":     87,
			"003":   47,
		},
		favDrinks: []string{
			"Whisky",
			"Coffee",
		},
	}

	for k, v := range p1.friends {
		fmt.Println(p1.first, "- friends - ", k, v)
	}

	for _, v := range p1.favDrinks {
		fmt.Println(p1.first, "- drinks - ", v)
	}

}
