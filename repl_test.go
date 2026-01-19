package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "this is a string",
			expected: []string{"this", "is", "a", "string"},
		},
		{
			input:    "CAPITAL And lowercase",
			expected: []string{"capital", "and", "lowercase"},
		},
		{
			input:    "   whitespace before and after string   ",
			expected: []string{"whitespace", "before", "and", "after", "string"},
		},
		{
			input:    "big           space",
			expected: []string{"big", "space"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v, got: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v, got %v", expectedWord, word)
			}
		}
	}
}
