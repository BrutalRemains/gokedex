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
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "help me bro",
			expected: []string{"help", "me", "bro"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("legnth mismatch: got %d, want %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words mismatch: got %s, want %s", word, expectedWord)
			}
		}
	}
}
