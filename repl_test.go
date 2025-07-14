package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
}{
	{
		input: "Hello World",
		expected: []string{"hello", "world"},
	},
	{
		input: "   HEllo WorlD",
		expected: []string{"hello", "world"},
	},
	{
		input: "  I love SEALS       FrIendS",
		expected: []string{"i", "love", "seals", "friends"},
	},
}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lengths do not match actual %s vs expected %s", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("CleanInput: %s == %s vs expected %s", c.input, actual, c.expected)
			}
		}
		
	}
}
