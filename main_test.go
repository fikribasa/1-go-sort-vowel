package main

import (
	"testing"
)

var inputTest = []struct {
	input  string
	output string
}{
	{"omama", "aaomm"},
	{"osama", "aaoms"},
	{"HelloWorld", "eooHWdlllr"},
}

func TestCount(t *testing.T) {
	for _, val := range inputTest {
		s := SortString(val.input)
		if s != val.output {
			t.Errorf("SortString(%s) => %s, Should be %s, /n", val.input, s, val.output)
		}
	}
}
