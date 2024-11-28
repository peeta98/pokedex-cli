package main

import "testing"

func TestGetCommands(t *testing.T) {
	cases := []struct{
		got string
		expected string
	}{
			{
				got: getCommands()["help"].name,
				expected: "help",
			},
			{
				got: getCommands()["help"].description,
				expected: "Displays a help message",
			},
			{
				got: getCommands()["exit"].name,
				expected: "exit",
			},
			{
				got: getCommands()["exit"].description,
				expected: "Exit the Pokedex",
			},
	}

	for _, c := range cases {
		if c.got != c.expected {
			t.Errorf("expected %q, got %q", c.expected, c.got)
		}
	}
}

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  ",
			expected: []string{},
		},
		{
			input: "  hello  ",
			expected: []string{"hello"},
		},
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}