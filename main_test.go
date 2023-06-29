package main

import (
	"testing"
)

func TestReplaceCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			input:    "Ready, set, go (up) !",
			expected: "Ready, set, GO !",
		},
		{
			input:    "I should stop SHOUTING (low)",
			expected: "I should stop shouting",
		},
		{
			input:    "Welcome to the Brooklyn bridge (cap)",
			expected: "Welcome to the Brooklyn Bridge",
		},
		{
			input:    "This is so exciting (up, 2)",
			expected: "THIS IS SO EXCITING",
		},
	}

	for _, test := range tests {
		result := ReplaceCase(test.expected) // Swap input and expected
		if result != test.input {            // Swap input and expected
			t.Errorf("Input: %s\nExpected: %s\nGot: %s\n", test.expected, test.input, result) // Swap input and expected
		}
	}
}
