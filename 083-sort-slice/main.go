package main

import (
	"fmt"
	"sort"
)

type user struct {
	first   string
	last    string
	age     int
	sayings []string
}

func main() {
	u1 := user{
		first: "James",
		last:  "Bond",
		age:   32,
		sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
		sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		first: "M",
		last:  "Hmmmm",
		age:   54,
		sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("-------")

	sort.Slice(users, func(i, j int) bool { return users[i].age < users[j].age })
	fmt.Println("By age:", users)

	fmt.Println("-------")

	sort.Slice(users, func(i, j int) bool { return users[i].last < users[j].last })
	fmt.Println("By name:", users)
}
