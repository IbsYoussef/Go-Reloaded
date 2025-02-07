package unit_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestConvertHexToDecimal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"The value is 3A (hex) in the report", "The value is 58 in the report"},
		{"No hex value here", "No hex value here"},
		{"7F (hex) is the maximum value", "127 is the maximum value"},
		{"Maximum value is FF (hex)", "Maximum value is 255"},
		{"The value of A3 (hex) (important)", "The value of 163 (important)"},
		{"1G (hex) is not valid", "Error Could not convert hex value to string"},
		{"", ""},
		{"(hex) was found", "Error Could not convert hex value to string"},
		{"1E (hex) and 3A (hex) were added", "30 and 58 were added"},
		{"The value is 2f (hex)", "The value is 47"},
		{"0001 (hex) is a small number", "1 is a small number"},
	}

	for _, test := range tests {
		output := textmod.ConvertHexToDecimal(test.input)
		if output != test.expected {
			t.Errorf("For input %q, expected %q but got %q", test.input, test.expected, output)
		}
	}

}
