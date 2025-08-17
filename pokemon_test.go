package main

import "testing"

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
			input:    "Charmander Bulbasaur Pikachu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lenghth mismatch")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word!= expectedWord{
				t.Errorf("Error")
			}
		}
	}

}
