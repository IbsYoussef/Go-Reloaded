package textmod

import (
	"fmt"
	"strconv"
	"strings"
)

func HexandBin(s string) string {
	// Split the input string into individual words using spaces as delimiters
	sentence := strings.Fields(s)

	// Iterate through each word in the result slice using its index and value
	for i, v := range sentence {
		// Check if the current word is "(hex)"
		if v == "(hex)" {
			// Attempt to parse the previous word as a hexadecimal integer (base 16)
			hex, err := strconv.ParseInt(sentence[i-1], 16, 64)

			// Replace the previous word with the parsed hexadecimal value as a decimal string
			sentence[i-1] = fmt.Sprint(hex)

			// Set the current word to an empty string to remove it from the result, i.e Remove (hex)
			sentence[i] = ""

			// Handle any error that occurred during the conversion
			if err != nil {
				panic(err)
			}
		}

		// Check if the current word is "(bin)"
		if v == "(bin)" {
			// Attempt to parse the previous word as a binary integer (base 2)
			bin, err := strconv.ParseInt(sentence[i-1], 2, 64)

			// Replace the previous word with the parsed binary value as a decimal string
			sentence[i-1] = fmt.Sprint(bin)

			// Set the current word to an empty string to remove it from the result, i.e Remove (bin)
			sentence[i] = ""

			// Handle any error that occurred during the conversion
			if err != nil {
				panic(err)
			}
		}
	}

	// Join the modified result slice back into a string with spaces between words and return the final string
	return strings.Join(sentence, " ")
}
