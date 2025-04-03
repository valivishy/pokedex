package tests

import (
	"github.com/valivishy/pokedex/util"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HeLp",
			expected: []string{"help"},
		},
		{
			input:    "EXIT ",
			expected: []string{"exit"},
		},
		{
			input:    "  help  exit ",
			expected: []string{"help", "exit"},
		},
		{
			input:    "\n\tHelp\n",
			expected: []string{"help"},
		},
	}

	for _, c := range cases {
		actual := util.CleanInput(c.input)
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
