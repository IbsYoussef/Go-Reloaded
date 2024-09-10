package textmod

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ConvertBinToDecimal(text string) string {
	// Regular expression to find binary numbers followed by "(bin)"
	re := regexp.MustCompile(`\b([01]+)\b\s*\(bin\)`)

	// Replace function to convert binary to decimal
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		// Split the match into binary value and "(bin)"
		parts := strings.Fields(match)
		binValue := parts[0]

		// Convert the binary value to decimal
		decimalValue, err := strconv.ParseInt(binValue, 2, 64)
		if err != nil {
			return match // Return the original match if there's an error
		}

		// Return the decimal value as a string
		return fmt.Sprintf("%d", decimalValue)
	})

	return result
}
