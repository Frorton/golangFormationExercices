package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {

	p1 := person{
		first: "Quentin",
		age:   27,
	}
	p2 := person{
		first: "Léa",
		age:   27,
	}
	p3 := person{
		first: "Elie",
		age:   0,
	}
	p4 := person{
		first: "Amour",
		age:   11,
	}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)

	fmt.Println("-------")

	sort.Sort(byAge(people))
	fmt.Println(people)
}

// Using sort.Slice, simpler, but doesn't go through Sort interface and user defined type
/*
	sort.Slice(people, func(i, j int) bool { return people[i].first < people[j].first })
		fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].age < people[j].age })
		fmt.Println("By age:", people)
*/

/*
The stability of a sorting algorithm refers to preserving the order of duplicate values.
In a stable sort by age:

	{"James",32},{"Joe",20},{"Mark",20}

becomes,

	{"Joe",20},{"Mark",20},{"James",32}

Joe and Mark keep their relative position


In an unstable sort, this can happen:

	{"James",32},{"Joe",20},{"Mark",20}

becomes,

	{"Mark",20},{"Joe",20},{"James",32}

Mark and Joe switched positions, but the list is still sorted.
*/
