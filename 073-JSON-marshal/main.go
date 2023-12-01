package main

import (
	"encoding/json"
	"fmt"
)

// First, last & age capital to make it work with json
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Quentin",
		Last:  "BOURGEAIS",
		Age:   27,
	}
	p2 := person{
		First: "Elie",
		Last:  "BOURGEAIS",
		Age:   0,
	}
	p3 := person{
		First: "Léa",
		Last:  "BOURGEAIS",
		Age:   27,
	}
	people := []person{
		p1,
		p2,
		p3,
	}
	fmt.Println(people)

	//marshal
	sb, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error :", err)
	}

	// without capital letters in the type will return empty slices
	fmt.Println(string(sb))
}
