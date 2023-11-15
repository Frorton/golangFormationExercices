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

	fmt.Println("------")

	age := m["Q"]
	fmt.Println("Age of Q :", age)
}
