package greet

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Quentin")
	if s != "Hello there Quentin" {
		t.Error("got", s, "want", "Hello there Quentin")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Quentin"))
	// Output :
	// Hello there Quentin
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Quentin")
	}
}
