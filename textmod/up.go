package textmod

import (
	"strconv"
	"strings"
)

func Up(s string) string {
	sentence := strings.Fields(s)
	for i, v := range sentence {
		if v == "(up)" {
			sentence[i-1] = strings.ToUpper(sentence[i-1])
			sentence[i] = "" // Removes (cap)
		}
		if v == "(up," {
			// Converts word before to Upper Case
			sentence[i-1] = strings.ToUpper(sentence[i-1])
			boundary := len(sentence[i+1])
			number := sentence[i+1][:boundary-1] // 2
			num, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}

			for j := 1; j <= num; j++ {
				sentence[i-j] = strings.ToUpper(sentence[i-j])
			}
			sentence[i], sentence[i+1] = "", "" // Remove (cap) and number.
		}
	}
	return strings.Join(sentence, " ")
}
