package benchtesting

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "It is only with the heart that one can see rightly; what is essential is invisible to the eye."
	xs := strings.Split(s, " ")
	s = Cat(xs)
	if s != "It is only with the heart that one can see rightly; what is essential is invisible to the eye." {
		t.Error("got", s, "want", "It is only with the heart that one can see rightly; what is essential is invisible to the eye.")
	}
}

func TestJoin(t *testing.T) {
	s := "It is only with the heart that one can see rightly; what is essential is invisible to the eye."
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "It is only with the heart that one can see rightly; what is essential is invisible to the eye." {
		t.Error("got", s, "want", "It is only with the heart that one can see rightly; what is essential is invisible to the eye.")
	}
}

func ExampleCat() {
	s := "It is only with the heart that one can see rightly; what is essential is invisible to the eye."
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output:
	// It is only with the heart that one can see rightly; what is essential is invisible to the eye.
}

func ExampleJoin() {
	s := "It is only with the heart that one can see rightly; what is essential is invisible to the eye."
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output:
	// It is only with the heart that one can see rightly; what is essential is invisible to the eye.
}

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
