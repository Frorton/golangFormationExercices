package main

import (
	"fmt"

	"github.com/GoesToEleven/puppy"
)

func main() {
	fmt.Println("Puppy time")

	x1 := puppy.Bark()
	x2 := puppy.Barks()
	x3 := puppy.BigBark()

	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)
}
