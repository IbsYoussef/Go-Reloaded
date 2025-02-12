package textmod

import "strings"

func ReplaceAWithAn(s string) string {
	// Define a local function 'firstRune' to return the first rune (character) of a string
	firstRune := func(s string) string {
		vocal := []rune(s)
		return string(vocal[0])
	}

	// Initialize a strings.Builder to efficiently build the modified string
	var str strings.Builder
	result := strings.Fields(s)

	// Loop through each word (v) and its index (i) in the 'result' slice
	for i, v := range result {
		// Append the current word and a space to the strings.Builder 'str'
		str.WriteString(v + " ")

		// Check if the current word is "a" and if the first character of the next word is a vowel or "h"
		if v == "a" && (firstRune(result[i+1]) == "a" || firstRune(result[i+1]) == "e" || firstRune(result[i+1]) == "i" || firstRune(result[i+1]) == "o" || firstRune(result[i+1]) == "u" || firstRune(result[i+1]) == "h") {
			// Replace the current word "a" with "An" to correct the grammar
			result[i] = "An"
		}
	}

	// Join the modified 'result' slice back into a string with spaces between words and return the final string
	return strings.Join(result, " ")
}
