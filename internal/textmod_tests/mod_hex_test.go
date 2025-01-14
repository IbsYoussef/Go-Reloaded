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

func TestConvertHexToDecimalEdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		//Edge Cases
		// Edge cases
		{
			input:    "(hex) at the start of the sentence.",
			expected: "(hex) at the start of the sentence.", // No preceding word to convert
		},
		{
			input:    "Hex value with a lowercase tag: 1e (hex).",
			expected: "30.", // Should work with lowercase letters in the hex
		},
		{
			input:    "Hex value with no word: (hex).",
			expected: "(hex).", // Should not crash or modify text
		},
		{
			input:    "Mixed-case hex: AbC (hex) and 12aB (hex).",
			expected: "2748 and 4779.",
		},
		{
			input:    "Invalid hex format: G1 (hex) and HZ (hex).",
			expected: "G1 (hex) and HZ (hex).", // Invalid hex values should remain unchanged
		},
		{
			input:    "0 (hex) is the same as zero.",
			expected: "0 is the same as zero.", // Edge case for zero
		},
		{
			input:    "Trailing space after hex: 1E (hex) .",
			expected: "30 .", // Ensure correct handling of trailing spaces
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
