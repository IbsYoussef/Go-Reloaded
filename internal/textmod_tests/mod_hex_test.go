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
		{
			input:    "1E (hex) files were added and 3A (hex) were removed.",
			expected: "30 files were added and 58 were removed.",
		},
		{
			input:    "FF (hex) colors were selected.",
			expected: "255 colors were selected.",
		},
		{
			input:    "The code 0A (hex) caused an error.",
			expected: "The code 10 caused an error.",
		},
		{
			input:    "No hexadecimal numbers here.",
			expected: "No hexadecimal numbers here.",
		},
		{
			input:    "Multiple values: A (hex), B (hex), and C (hex).",
			expected: "Multiple values: 10, 11, and 12.",
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := textmod.ConvertHexToDecimal(test.input)
			if result != test.expected {
				t.Errorf("ConvertHexToDecimal(%q) = %q; want %q", test.input, result, test.expected)
			}
		})
	}
}
