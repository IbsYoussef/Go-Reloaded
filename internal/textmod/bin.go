package textmod

import (
	"strconv"
	"strings"
)

// ConvertBinToDecimal replaces instances of a word preceding the (bin) tag with the decimal equivalent of the binary number.
func ConvertBinToDecimal(text string) string {
	if text == "" {
		return text
	}

	// Split the text into fields (words)
	splitText := strings.Fields(text)

	// Iterate over the tokens
	for i := 0; i < len(splitText); i++ {
		// Remove a trailing period from the token if present
		trimmedToken := strings.TrimSuffix(splitText[i], ".")
		// If the token (after trimming) equals "(bin)", then do the conversion.
		if trimmedToken == "(bin)" {
			// Check that there is a preceding token
			if i > 0 {
				// Determine if the (bin) tag had punctuation attached.
				// We'll use it later to re-append the punctuation if needed.
				// (In your test cases, sometimes you want to drop the punctuation and sometimes preserve it.)
				hasPeriod := splitText[i] != trimmedToken

				binValue := splitText[i-1]
				// Attempt to parse the binary value.
				decValue, err := strconv.ParseInt(binValue, 2, 64)
				if err != nil {
					// Return error message without extra spaces
					return "Error could not convert bin value to string"
				}
				// Convert the decimal number to a string.
				strValue := strconv.Itoa(int(decValue))
				// If there was punctuation on the (bin) tag, re-append it to the converted number.
				if hasPeriod {
					strValue += "."
				}
				// Replace the preceding token with the converted value.
				splitText[i-1] = strValue
				// Remove the (bin) tag token.
				splitText = append(splitText[:i], splitText[i+1:]...)
				// Adjust the index because we removed an element.
				i--
			} else {
				return "Error could not convert bin value to string"
			}
		}
	}

	return strings.Join(splitText, " ")
}
