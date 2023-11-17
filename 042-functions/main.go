/*
func sum(ii []int) (total int) {
	total = 0
	for _, v := range ii {
		total += v
	}
	return
}
Name returns ➡ pas ouf

DRY = Do not repeat yourself
KISS = Keep it stupid simple
*/

package main

import "fmt"

func main() {

	x := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(x)
}

func sum(ii []int) int { // (total int) disparaît, total = 0 ➡ t := 0
	t := 0 // Short scope ➡ short variable names 🤏
	for _, v := range ii {
		t += v
	}
	return t
}
