/*
Callback functions are very common in many programming languages.
Callbacks are essentially functions that are passed as arguments to others functions and are
	intended to be called (or executed) later in the program.
*/

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(printSquare(square, i))
	}
	fmt.Println("Tadaaa")
}

//Printsquare returns a string of a func that takes an int and returns an int
func printSquare(f func(int) int, a int) string {
	x := f(a)
	return fmt.Sprintf("the number %v squared is %v\n", a, x)
}

//Square takes an int called n and returns an int n*n
func square(n int) int {
	return n * n
}
