package textmod

import (
	"strconv"
	"strings"
)

func Low(s string) string {
	result := strings.Fields(s)
	for i, v := range result {
		if v == "(low)" {
			result[i-1] = strings.ToLower(result[i-1])
			result[i] = "" // Removes (low)
		}
		if v == "(low," {
			result[i-1] = strings.ToLower(result[i-1])
			twov := len(result[i+1])
			numb := result[i+1][:twov-1]
			nu, err := strconv.Atoi(numb)
			if err != nil {
				panic(err)
			}
			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToLower(result[i-j])
			}
			result[i], result[i+1] = "", "" // Removes (low) and number.
		}
	}
	return strings.Join(result, " ")
}
