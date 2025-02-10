package unit_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestModifyPunctuationAndQuotes(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
	}

	for _, test := range tests {
		output := textmod.ModifyPunctuationAndQuotes(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, output)
		}
	}
}
