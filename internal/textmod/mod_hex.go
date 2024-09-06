package textmod

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func ConvertHexToDecimal(text string) string {
	// Regular expression to find hexadecimal numbers followed by "(hex)"
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+)\b\s*\(hex\)`)

	// Replace function to convert hex to decimal
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		parts := strings.Fields(match)
		hexValue := parts[0]
		decimalValue, err := strconv.ParseInt(hexValue, 16, 64)
		if err != nil {
			return match // Return the original match if there's an error
		}
		return fmt.Sprintf("%d", decimalValue)
	})

	return result
}
