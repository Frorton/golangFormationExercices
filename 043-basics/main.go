/*
- Create a func with the identifier foo that returns an int
- Create a func with the identifier bar that returns an int and a string
- Call both funcs
- Print out their results
*/

package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x)

	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "James Bond"
}
