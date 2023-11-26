package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "Quentin",
	}
	fmt.Println(p)

	fmt.Println("-------")

	p = changename(p, "Léa")
	fmt.Println(p)

	fmt.Println("-------")

	changenamep(&p, "Elie")
	fmt.Println(p)
}

func changename(p person, s string) person {
	p.first = s
	return p
}

func changenamep(p *person, s string) {
	p.first = s
}
