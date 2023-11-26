/*
To document :
	The convention for documenting code uses comments directly preceding the declaration or definition of the item being documented.
	This style of documentation is used by the godoc tool to automatically generate API documentation.

- Package documentation :
	Should be a brief one-liner that begins with "Package <packagename>" and provides a high level overview of the package's purpose.
	It is typically located in a separate "doc.go" file.

- Function/method documentation :
	Comments should be written directly above the function, starting with the function's name.

- Type documentation :
	Idem, starting with the type's name.

- Constantsand variables documentation :
	Idem, starting with the constant's or variable's name.

- Other Go conventions :
	Use complete sentences.
		The first sentence should be a brief summary starting with the name being declared.
	Prefer breaking long lines after punctuation or operators, keeping operands together.
	Do not use URLs in package-level comments, it's better to create a README for that purpose.
	Remember : KISS Keep It Stupid Simple

"go help doc" for in-terminal information
*/

package main

import "fmt"

// main calls newborn which returns a string + "dad"
func main() {
	fmt.Println(newborn("dad"))
}

// newborn returns a string
func newborn(b string) string {
	return fmt.Sprint("I'm about to be a ", b)
}
