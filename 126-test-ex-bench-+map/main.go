package main

import (
	"Golang/golangFormationExercices/126-test-ex-bench-+map/quote"
	"Golang/golangFormationExercices/126-test-ex-bench-+map/word"
	"fmt"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
