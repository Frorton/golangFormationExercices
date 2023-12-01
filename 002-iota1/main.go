// Tour of Go - Effective go - iota
package main

import "fmt"

/*
const (
	c0 = iota - c0 == 0
	c1 = iota - c1 == 1
	c2 = iota - c2 == 2
	c3 = iota - etc
	c4 = iota
	c5 = iota
	c6 = iota
)
*/
const (
	_ = iota // c0 == 0
	a
	b
	c
	d
	e
	f
)

func main() {
	//fmt.Println(c0, c1, c2)
	//fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b\n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b\n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b\n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b\n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b\n", 1<<f, 1<<f)
}
