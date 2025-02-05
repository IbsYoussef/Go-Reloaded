package textmod

import (
	"strconv"
	"strings"
)

func ChangeCase(text string) string {
	if text == "" {
		return text
	}

	strSlice := strings.Fields(text)

	switch strSlice[0] {
	case "(cap,", "(low,", "(up,", "(cap)", "(low)", "(up)":
		return "Tag cannot be at the start"
	}

	for i := 0; i < len(strSlice); i++ {
		if strSlice[i] == "(up)" || strSlice[i] == "(low)" || strSlice[i] == "(cap)" {
			tag := strSlice[i]
			if i > 0 {
				word := strSlice[i-1]
				switch tag {
				case "(cap)":
					word = strings.Title(word)
				case "(low)":
					word = strings.ToLower(word)
				case "(up)":
					word = strings.ToUpper(word)
				}
				strSlice[i-1] = word
				strSlice = append(strSlice[:i], strSlice[i+1:]...)
				i--
			}
		}

		if strSlice[i] == "(cap," || strSlice[i] == "(low," || strSlice[i] == "(up," {
			tagNum := strSlice[i+1]
			numStr := ""
			if strings.HasSuffix(tagNum, ")") {
				numStr = strings.TrimSuffix(tagNum, ")")
			}

			if i > 0 {
				number, err := strconv.Atoi(numStr)
				if err != nil {
					return "Could not convert string number to integer"
				}
				switch strSlice[i] {
				case "(cap,":

					for j := 1; j <= number; j++ {
						strSlice[i-j] = strings.Title(strSlice[i-j])
					}
				case "(low,":

					for j := 1; j <= number; j++ {
						strSlice[i-j] = strings.ToLower(strSlice[i-j])
					}

				case "(up,":

					for j := 1; j <= number; j++ {
						strSlice[i-j] = strings.ToUpper(strSlice[i-j])
					}
				}
				strSlice[i], strSlice[i+1] = "", ""
			}
		}
	}
	return strings.Join(strSlice, " ")
}
