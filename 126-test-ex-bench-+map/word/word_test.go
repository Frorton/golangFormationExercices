package word

import (
	"Golang/golangFormationExercices/126-test-ex-bench-+map/quote"
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	s := Count(quote.SaintEx)
	if s != 19 {
		t.Error("got", s, "want", 49)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("a b c")
	for k, v := range m {
		switch k {
		case "a":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "b":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "c":
			if v != 1 {
				t.Error("got", v, "want", 3)
			}
		}

	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SaintEx))
	// Output:
	// 19
}

// No example for UseCount because it's a map and it won't result in the same order everytime

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SaintEx)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SaintEx)
	}
}
