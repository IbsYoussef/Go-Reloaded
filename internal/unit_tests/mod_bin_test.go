package unit_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestConvertBinToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			input:    "There are 101 (bin) ways to solve this.",
			expected: "There are 5 ways to solve this.",
		},
		// Edge cases
		{
			input:    "(bin) with no preceding binary number.",
			expected: "(bin) with no preceding binary number.",
		},
		{
			input:    "This binary 2 (bin) is invalid.",
			expected: "This binary 2 (bin) is invalid.",
		},
		{
			input:    "Multiple values: 1 (bin), 10 (bin), and 11 (bin).",
			expected: "Multiple values: 1, 2, and 3.",
		},
		{
			input:    "Binary at the start: 110 (bin) is valid.",
			expected: "Binary at the start: 6 is valid.",
		},
		{
			input:    "Trailing space after binary: 1010 (bin) .",
			expected: "Trailing space after binary: 10 .", // Correct handling of trailing spaces
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := textmod.ConvertBinToDecimal(test.input)
			if result != test.expected {
				t.Errorf("ConvertBinToDecimal(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}

}
