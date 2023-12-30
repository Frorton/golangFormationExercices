package chien

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	s := Years(7)
	if s != 49 {
		t.Error("got", s, "want", 49)
	}
}

func TestYearsTwo(t *testing.T) {
	s := YearsTwo(10)
	if s != 70 {
		t.Error("got", s, "want", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
