package main

import "testing"

func TestMyAdd(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}
	// The "table test part"
	tests := []test{
		{[]int{-21, 21}, 0},
		{[]int{8, -4, 2}, 6},
		{[]int{10, 1}, 11},
		{[]int{-1, 0, 1}, 0},
	}

	// Ranging over the table test
	for _, v := range tests {
		x := myAdd(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
