package textmod

import (
	"fmt"
	"regexp"
	"strconv"
)

// ConvertBinToDecimal replaces instances of (bin) with the decimal equivalent of the preceding binary number.
func ConvertBinToDecimal(text string) string {
	// Regular expression to match a binary number followed by (bin)
	re := regexp.MustCompile(`\b([01]+)\s\(bin\)`)

	// Replace all matches with their decimal equivalents
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		// Extract the binary number from the match using capture groups
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // Return the original match if no valid binary number found
		}

		binaryNumber := submatches[1]

		// Convert the binary number to a decimal integer
		decimalValue, err := strconv.ParseInt(binaryNumber, 2, 64)
		if err != nil {
			return match // Return the original match if conversion fails
		}

		// Return the decimal value as a string
		return fmt.Sprintf("%d", decimalValue)
	})

	return result
}
