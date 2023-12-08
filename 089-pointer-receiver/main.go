/*
Create a type person struct
- Attach a method speak to type person using a pointer receiver
	*person
- Create a type human interface
	To implicitly implement the interface, a human must have the speak method
- Create func “saySomething”
	Have it take in a human as a parameter
	Have it call the speak method
- Show the following in your code
	You CAN pass a value of type *person into saySomething
	You CANNOT pass a value of type person into saySomething
*/

package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak() string
}

func (p *person) speak() string {
	return "I am human, I can speak."
}

func saySomething(h human) {
	fmt.Println("Are you human ? Can you speak ?\n", h.speak())
}

func main() {
	p := person{
		first: "Quentin",
	}
	fmt.Println(p)

	fmt.Println("-------")

	//saySomething(p)
	saySomething(&p)
}
