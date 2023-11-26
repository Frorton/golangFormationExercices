//Can be called lambda func or func literal

package main

import "fmt"

func main() {
	func() {
		a := "Through the ages of time"
		b := "I've been known for my hate"
		c := "But I'm a dealer of simple choices"
		d := "For me it's never too late"
		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(d)
	}()
}
