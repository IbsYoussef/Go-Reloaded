package text_mod_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestConvertBinToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"101 (bin) is a small number", "5 is a small number"},
		{"I paid 1001 (bin) dollars", "I paid 9 dollars"},
		{"110 (bin), 111 (bin), and 1000 (bin) are all binary numbers", "6, 7, and 8 are all binary numbers"},
	}

	for _, test := range tests {
		result := textmod.ConvertBinToDecimal(test.input)
		if result != test.expected {
			t.Errorf("For input %s, expected: %s, but got: %s", test.input, test.expected, result)
		}
	}
}
