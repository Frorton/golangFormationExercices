/*
A unit test is a type of software testing where individual components or units of a software are tested.
The purpose is to validate that each unit of the software performs as designed.
A unit is the smallest testable part of any software, it usually has one or a few inputs and usually a single output.
In Go unit testing represents the testing of a single function, method or struct.
It is typically written using the "testing" package.
	Note : unit testing is a subset of tests. Tests can be of many forms :
		- Unit tests
		- Integration tests
		- Functional tests
		- System tests, etc...

Summary :
	Unit testing is one of many kind of test available in Go. It's purpose is to test the behaviour of small,
	isolated pieces of functionality. Over types of tests have different focuses and may require different tools and packages.
*/

package main

import "fmt"

func main() {

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(1, 1, add)
	fmt.Println(x)

	y := doMath(1, 1, subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
