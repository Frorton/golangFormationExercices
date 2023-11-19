package main

import "fmt"

func main() {

	fmt.Println(newborn("dad"))
}

func newborn(b string) string {
	return fmt.Sprint("I'm about to be a ", b)
}
