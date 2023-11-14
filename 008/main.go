// créer hors func main
package main

import "fmt"

var x string = "Hey Ho, Let's GO!"

const y int = 21

func main() {

	z := 20

	soustr := y - z

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(soustr)
}
