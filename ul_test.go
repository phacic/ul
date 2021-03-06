package main

import (
	"testing"
)

type upperLowerTest struct {
	val   string
	count int
	want  string
}

var upperTests = []upperLowerTest{
	{"something", 0, "SOMETHING"},
	{"something", -1, "SOMETHING"},
	{"something", 3, "SOMething"},
	{"something", 25, "SOMETHING"},
}

func TestToUpper(t *testing.T) {
	for _, test := range upperTests {
		got := ToUpper(test.val, test.count)
		if got != test.want {
			t.Errorf("toUpper(%s, %v) = %s, want %s", test.val, test.count, got, test.want)
		}
	}
}

var lowerTests = []upperLowerTest{
	{"SOMETHING", 0, "something"},
	{"SOMETHING", -1, "something"},
	{"SOMETHING", 3, "somETHING"},
	{"SOMETHING", 5, "sometHING"},
	{"SOMETHING", 25, "something"},
}

func TestToLower(t *testing.T) {
	for _, test := range lowerTests {
		got := ToLower(test.val, test.count)
		if got != test.want {
			t.Errorf("ToLower(%s, %v) = %s, want %s", test.val, test.count, got, test.want)
		}
	}
}
