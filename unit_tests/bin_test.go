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
		// Test Cases
		{"", ""}, // Passes
		{"The binary value is 1010 (bin)", "The binary value is 10"},
		{"There are no binary tags here.", "There are no binary tags here."},
		{"(bin) 1010 is not valid.", "Error could not convert bin value to string"},
		{"NotBinary (bin) but 110 (bin).", "Error could not convert bin value to string"},
		{"(bin) is invalid usage.", "Error could not convert bin value to string"},
		{"0010 (bin) and 0001 (bin)", "2 and 1"},
		{"Spaces   110 (bin)   are   tricky.", "Spaces 6 are tricky."},
		{"1101101010101010101010 (bin) is huge.", "3582634 is huge."},
	}

	for _, test := range tests {
		output := textmod.ConvertBinToDecimal(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, output)
		}
	}
}
