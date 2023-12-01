package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Quentin","Last":"BOURGEAIS","Age":27},
{"First":"Elie","Last":"BOURGEAIS","Age":0},
{"First":"Léa","Last":"BOURGEAIS","Age":27}]`

	sb := []byte(s)

	// sb is of type uint8 (alias for byte)
	fmt.Printf("s type : %T\t\t sb type : %T\n", s, sb)
	fmt.Println("-------")

	people := []person{} // or : var people []person whichever is more readable

	// If err already declared above, use = instead of :=
	err := json.Unmarshal(sb, &people)
	if err != nil {
		fmt.Println("error :", err)
	}
	// fmt.Println("All the data : ", people)
	for i, v := range people {
		fmt.Println("\n--- PERSON NUMBER ", i, "---")
		fmt.Println(v.First, v.Last, v.Age)

	}
}
