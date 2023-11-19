/*
Create a type SQUARE
	length int
	width int
- Create a type CIRCLE
	radius float64
- Attach a method to each that calculates AREA and returns it
	circle area= π r 2
	math.Pi
	math.Pow
	square area = L * W
- Create a type SHAPE that defines an interface as anything that has the AREA method
- Create a func INFO which takes type shape and then prints the area
- Create a value of type square and of type circle
- Use func info to print the area of square and the area of circle
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64 // pas int par praticité
	width  float64 // rappel int = whole numbers, float64 = IEEE-754 64-bit floating-point numbers
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64 // any type that has the area method is also of type shape
}

func info(s shape) float64 {
	return s.area()
}

func main() {

	s1 := square{
		length: 7,
		width:  13,
	}

	c1 := circle{
		radius: 21,
	}

	fmt.Println("Area of circle is : ", info(c1)) // Polymorphism
	fmt.Println("Area of square is : ", info(s1)) // a value can be of more than 1 type
}
