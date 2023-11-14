package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

/*
Each source file can define its own niladic init function to set up whatever state is required.
(Actually each file can have multiple init functions.)
And finally means finally: init is called after all the variable declarations in the package have evaluated their initializers,
and those are evaluated only after all the imported packages have been initialized.
Besides initializations that cannot be expressed as declarations,
a common use of init functions is to verify or repair correctness of the program state before real execution begins.
*/

func main() {
	x := rand.Intn(250)
	fmt.Printf("x is of value %v\n", x)

	switch {
	case x <= 100:
		fmt.Println("x is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("x is between 101 and 200")
	default:
		fmt.Println("x is between 201 and 250")
	}
}

// niladic : function or program that has no arguments (ex : func init)
