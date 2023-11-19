/*
Go provides a built-in framework that simplifies the process of testing.

File structure and naming conventions :

	- Test files : They live in the same package as the code they test. They are named with the following convention :
		"filename_test.go" where filename is the name of the file that you want to test.

	- Test functions : They must start with the word "Test" followed by a word that starts with a capitale letter :
		"func TestXxx(*testing.T)", Xxx must start with a capital letter.

	Example :
*/
//main.go
package main

import "fmt"

func main() {
	fmt.Println(Add(3, 4))
}
func Add(a int, b int) int {
	return a + b
}
