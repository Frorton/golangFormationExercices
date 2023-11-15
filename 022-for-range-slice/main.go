package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46}

	for i, v := range x {
		fmt.Printf("Index %v \t valeur %v\n", i, v)
	}
}
