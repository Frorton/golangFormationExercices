package main

import "fmt"

func main() {
	fmt.Println("3 + 2 =", myAdd(3, 2))
	fmt.Println("7 + 4 =", myAdd(7, 4))
	fmt.Println("13 + 5 =", myAdd(13, 5))
}

func myAdd(xi ...int) int {
	var add int
	for _, v := range xi {
		add += v
	}
	return add
}
