package chapter1

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {

	in, want := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics.", []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	if got := WordCount(in); reflect.DeepEqual(in, got) {
		t.Errorf("Reverse(%q) == %q, want %q", in, got, want)
	}
}