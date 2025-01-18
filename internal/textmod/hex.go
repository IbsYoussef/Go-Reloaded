package textmod

import (
	"strconv"
	"strings"
)

func ConvertHexToDecimal(text string) string {
	if text == "" {
		return text
	}

	slicedText := strings.Fields(text)

	for i := 0; i < len(slicedText); i++ {
		if slicedText[i] == "(hex)" {
			// Check if there is a valid word before "(hex)"
			if i > 0 {
				word := slicedText[i-1]
				hexValue, err := strconv.ParseInt(word, 16, 64)
				if err != nil {
					return "Error Could not convert hex value to string"
				}
				strValue := strconv.Itoa(int(hexValue))
				slicedText[i-1] = strValue
				slicedText = append(slicedText[:i], slicedText[i+1:]...)
				i--
			} else {
				// Handle invalid case where "(hex)" is at the beginning
				return "Error Could not convert hex value to string"
			}
		}
	}

	return strings.Join(slicedText, " ")
}
