// Adding test/example/benchmark to starting code
package main

import (
	chien "Golang/golangFormationExercices/125-test-ex-bench/dog"
	"fmt"
)

type canine struct {
	name string
	age  int
}

func main() {
	klyde := canine{
		name: "Klyde",
		age:  chien.Years(10),
	}
	fmt.Println(klyde)
	fmt.Println(chien.YearsTwo(20))
}
