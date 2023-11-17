/*
- “defer” multiple functions in main
- Show that a deferred func runs after the func containing it exits.
- Determine the order in which the multiple defer funcs run
*/

package main

import "fmt"

func main() {

	defer fmt.Println(3)
	defer fmt.Println(2)
	defer fmt.Println(1)
	fmt.Println(0)

}

// deferred functions run in LIFO order
// last in first out LIFO
// often used to close files etc
