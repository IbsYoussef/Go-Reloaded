package textmod

import (
	"fmt"
	"strconv"
	"strings"
)

// ConvertBinToDecimal replaces instances of a word preceding the (bin) tag with the decimal equivalent of the binary number.
func ConvertBinToDecimal(text string) string {
	if text == "" {
		return text
	}

	slicedText := strings.Fields(text)

	for i := 0; i < len(slicedText); i++ {
		if slicedText[i] == "(bin)" {
			if i > 0 {
				word := slicedText[i-1]
				binValue, err := strconv.ParseInt(word, 2, 64)
				if err != nil {
					fmt.Println("Error could not convert bin value to string")
				}
				strValue := strconv.Itoa(int(binValue))
				slicedText[i-1] = strValue
				slicedText = append(slicedText[:i], slicedText[i+1:]...)
				i--
			} else {
				return "Error could not convert bin value to string"
			}
		}
	}

	return strings.Join(slicedText, " ")
}
