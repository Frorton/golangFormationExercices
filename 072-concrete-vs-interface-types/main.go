/*
Concrete type :
	Type you can directly create a value from.
Interface type :
	Type you cannot directly create a value from
	You have to create instances of concrete types
	to satisfy the interface.
*/

package main

import "fmt"

// Concrete type
type book struct {
	title string
	year  int
}

func (b book) listing() {
	fmt.Printf("Title : %v\t First publication : %v\n", b.title, b.year)
}

//Interface type
type author interface {
	listing()
}

func list(a author) {
	a.listing()
}
func main() {

	b1 := book{
		title: "La recherche de l'absolu",
		year:  1834,
	}
	list(b1)
}
