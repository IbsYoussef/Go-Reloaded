package main

import "testing"

func TestApplyModifications(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1E (hex) files were added",
			expected: "30 files were added",
		},
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
			expected: "This is SO EXCITING",
		},
		{
			input:    "I was sitting over there ,and then BAMM !!",
			expected: "I was sitting over there, and then BAMM!!",
		},
		{
			input:    "I was thinking ... You were right",
			expected: "I was thinking... You were right",
		},
		{
			input:    "I am exactly how they describe me: ' awesome '",
			expected: "I am exactly how they describe me: 'awesome'",
		},
		{
			input:    "As Elton John said: ' I am the most well-known homosexual in the world '",
			expected: "As Elton John said: 'I am the most well-known homosexual in the world'",
		},
		{
			input:    "There it was. A amazing rock!",
			expected: "There it was. An amazing rock!",
		},
	}

	for _, test := range tests {
		result := applyModifications(test.input)
		if result != test.expected {
			t.Errorf("Expected: %s, Got: %s", test.expected, result)
		}
	}
}
