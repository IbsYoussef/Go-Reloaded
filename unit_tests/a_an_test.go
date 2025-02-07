package unit_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestConvertAtoAn(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"A apple", "An apple"},
		{"There is a elephant in a zoo", "There is an elephant in a zoo"},
		{"He has a dog and a cat.", "He has a dog and a cat."},
		{"She found a umbrella, and a attic.", "She found an umbrella, and an attic."},
		{"a", "a"},
		{"A apple a banana a orange a grape", "An apple a banana an orange a grape"},
	}

	for _, test := range tests {
		output := textmod.ConvertAtoAn(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected %s, but got %s", test.input, test.expected, output)
		}
	}
}
