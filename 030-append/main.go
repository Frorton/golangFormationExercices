/*
Start with this slice:
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
- Append to that slice this value:
	52
- Print out the slice.
- In ONE STATEMENT append to that slice these values:
	53
	54
	55
- Print out the slice.
- Append to the slice this slice:
	y := []int{56, 57, 58, 59, 60}
- Print out the slice.
*/

package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // y... ➡ takes everything inside y
	fmt.Println(x)
}
