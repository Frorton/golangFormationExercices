package main

import "testing"

func TestDoMath(t *testing.T) {
	got := doMath(1, 1, add)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestAdd(t *testing.T) {
	got := add(1, 1)
	want := 2
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}

func TestSubtract(t *testing.T) {
	got := subtract(1, 1)
	want := 0
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", got, want)
	}
}
