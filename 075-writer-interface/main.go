package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Am bout be a dad")

	// What Println calls for :
	// fmt.Fprintln(w io.Writer, a ...any) (n int, err error)

	// os.Stdout is a pointer to a file which is of type Writer
	fmt.Fprintln(os.Stdout, "Am bout be a dad")

	// io.WriteString(w io.Writer, s string) (n int, err error)
	io.WriteString(os.Stdout, "Am bout be a dad")
}

/*
Interface : inter (“between”) + face (“shape, figure, form”)
	The place at which independent and often unrelated systems meet and act on or communicate with each other.
*/
