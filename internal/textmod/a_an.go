package textmod

import "strings"

func ConvertAtoAn(text string) string {
	words := strings.Fields(text)

	for i := 0; i < len(words)-1; i++ {
		if words[i] == "a" || words[i] == "A" {
			if len(words[i+1]) > 0 && strings.Contains("aeiouAEIOU", string(words[i+1][0])) {
				words[i] = words[i] + "n"
			}
		}
	}

	text = strings.Join(words, " ")
	return text
}
