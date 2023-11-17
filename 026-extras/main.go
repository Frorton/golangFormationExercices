package main

import "fmt"

func main() {

	// fmt.Printf("true and true is : %v\n", (true && true)) // redundant but works
	fmt.Printf("true and false is : %v\n", (true && false))
	// fmt.Printf("true or true is : %v\n", (true || true)) // commented to clear the yellow stuff
	fmt.Printf("true or false is : %v\n", (true || false))
	fmt.Printf("not true is : %v\n", !true)

	fmt.Println("-------")

	//switch string
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("Living dangerously!")
	}

	fmt.Println("-------")

	//switch no default
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}

	fmt.Println("-------")

	//if else string
	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}

	fmt.Println("-------")

	//if string
	y := "James Bond"

	if y == "James Bond" {
		fmt.Println(x)
	}
}
