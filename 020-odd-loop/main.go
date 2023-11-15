package main

import "fmt"

func main() {
	x := 52
	for x > 0 {
		if x%2 != 0 {
			fmt.Println(x)
		}
		x--
	}
}
