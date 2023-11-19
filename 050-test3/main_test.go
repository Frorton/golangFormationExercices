package main

import "testing"

func TestNewborn(t *testing.T) {
	got := newborn("dad")
	want := "I'm about to be a dad"
	if got != want {
		t.Errorf("Text was incorrect, got: %s, want: %s.", got, want) // %s = string
	}
}
