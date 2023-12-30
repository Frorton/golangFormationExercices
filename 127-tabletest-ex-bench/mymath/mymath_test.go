package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data []int
		avg  float64
	}
	// The "table test part"
	tests := []test{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5},
		{[]int{1, 1, 1, 1, 1, 2}, 1},
		{[]int{10, 20, 1020, 100}, 60},
		{[]int{1, 11, 11, 3}, 7},
	}

	// Ranging over the table test
	for _, v := range tests {
		x := CenteredAvg(v.data)
		if x != v.avg {
			t.Error("Expected", v.avg, "Got", x)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 2, 4, 6, 9}))
	// Output :
	// 4
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 2, 3, 1, 2, 3})
	}
}
