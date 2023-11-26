package main

import "fmt"

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is", d.first, "and I walk.")
}

func (d dog) run() {
	fmt.Println("My name is", d.first, "and I run very fast.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()
}

func main() {
	d1 := dog{"Chien"}
	youngRun(d1)

	d2 := dog{"Midget"}
	youngRun(d2)
	d2 = d2.changeName("Dog")
	youngRun(d2)
}

func (d dog) changeName(s string) dog {
	d.first = s
	return d
}

// No pointer semantics because it's safer and data is small
