package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		for j := 5; j > 0; j-- {
			fmt.Printf("Outer loop %v \t Inner loop %v\n", i, j)
		}
	}
}
