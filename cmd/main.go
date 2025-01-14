package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ConvertHexToDecimal(text string) string {
	// Compile the expression to search for (hex) instances
	re := regexp.MustCompile(`\b([0-9A-Fa-f]+) \(hex\)`)

	result := re.ReplaceAllStringFunc(text, func(match string) string {
		// Extract hexadecimal number using a submatch
		submatches := re.FindStringSubmatch(match)
		if len(submatches) < 2 {
			return match // Return the original match if no submatch
		}
		hexMatch := submatches[1]

		// Convert the hex string to a decimal integer
		decimalValue, err := strconv.ParseInt(hexMatch, 16, 64)
		if err != nil {
			return match // Return the original match if conversion fails
		}

		// Return the decimal value as a string
		return fmt.Sprintf("%d", decimalValue)
	})

	return result
}

func main() {
	text := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	fmt.Println(ConvertHexToDecimal(text)) // Output: Simply add 66 and 10 (bin) and you will see the result is 68.
}
