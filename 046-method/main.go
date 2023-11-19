/*
Create a user defined struct with the identifier “person”, the fields:
	first
	age
- Attach a method to type person with the identifier “speak”
- The method should have the person say their name and age
- Create a value of type person
- Call the method from the value of type person
*/

package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) Speak() {
	fmt.Println("My name is", p.first, "and my age is", p.age)
}

func main() {

	p1 := person{
		first: "Elie",
		age:   0,
	}

	p1.Speak()
}
