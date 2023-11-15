package main

import "fmt"

func main() {

	m := map[string]int{
		"James": 42,
		"Jenny": 37,
	}

	for k, v := range m {
		fmt.Printf("Name : %v \t age : %v\n", k, v)
	}
}
