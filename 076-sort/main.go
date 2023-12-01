package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{1, 4, 7, 2, 5, 8, 3, 6, 9}
	xs := []string{"A", "C", "F", "B", "I", "H", "E", "G", "D"}

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("-------")

	sort.Ints(xi)
	fmt.Println(xi)
	sort.Strings(xs)
	fmt.Println(xs)
}
