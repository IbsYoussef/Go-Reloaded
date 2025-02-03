package textmod

import (
	"strings"
)

func ChangeCase(text string) string {
	if text == "" {
		return text
	}

	strSlice := strings.Fields(text)

	for i := 0; i < len(strSlice); i++ {
		if strSlice[i] == "(up)" || strSlice[i] == "(low)" || strSlice[i] == "(cap)" {
			tag := strSlice[i]
			if i > 0 {
				word := strSlice[i-1]
				switch tag {
				case "(up)":
					word = strings.ToUpper(word)
				case "(low)":
					word = strings.ToLower(word)
				case "(cap)":
					word = strings.Title(word)
				}
				strSlice[i-1] = word
				strSlice = append(strSlice[:i], strSlice[i+1:]...)
				i--
			} else {
				return "Error could perform string operation"
			}
		}
	}

	return strings.Join(strSlice, " ")
}
