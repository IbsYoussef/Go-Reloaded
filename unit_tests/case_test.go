package unit_tests

import (
	"go-reloaded/internal/textmod"
	"testing"
)

func TestChangeCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"go (up)", "GO"},
		{"harold (cap) gundersan (cap)", "Harold Gundersan"},
		{"GO (low)", "go"},
		{"one two three lets go (up, 5)", "ONE TWO THREE LETS GO"},
		{"ONE TWO THREE LETS GO (low, 5)", "one two three lets go"},
		{"one two three lets go (cap, 5)", "One Two Three Lets Go"},
		{"(cap) no change", "Tag cannot be at the start"},
		{"(low) no change", "Tag cannot be at the start"},
		{"(cap) no change", "Tag cannot be at the start"},
	}

	for _, test := range tests {
		output := textmod.ChangeCase(test.input)
		if output != test.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, output)
		}
	}
}
