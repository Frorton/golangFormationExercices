/*
Sets of type T or *T
*/
package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I'm walking.")
}

func (d *dog) run() {
	d.first = "Rover"
	fmt.Println("My name is", d.first, "and I'm running.")
}

func main() {
	d1 := dog{"Henry"}
	d1.walk()
	d1.run() // Can use different types (d1 = dog; run = *dog) as long as it's not to implement an interface

	d2 := &dog{"Padget"}
	d2.walk()
	d2.run()
}
