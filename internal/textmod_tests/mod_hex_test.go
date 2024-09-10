package text_mod_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestConvertHexToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"I have 1E (hex) apples", "I have 30 apples"},
		{"You owe me 2A (hex) dollars for lunch", "You owe me 42 dollars for lunch"},
		{"I have no hexadecimal here", "I have no hexadecimal here"},
		{"1F (hex) is my number", "31 is my number"},
		{"My score is 3C (hex)", "My score is 60"},
		{"I owe 1E (hex) and 2B (hex) dollars", "I owe 30 and 43 dollars"},
		{"I got 00A (hex) candies", "I got 10 candies"},
		{"", ""},
		{"I bought 1E (hex), 2F (hex), and 3A (hex) items", "I bought 30, 47, and 58 items"},
	}

	for _, test := range tests {
		result := textmod.ConvertHexToDecimal(test.input)
		if result != test.expected {
			t.Errorf("For input: %s, expected: %s, but got: %s", test.input, test.expected, result)
		}
	}
}
