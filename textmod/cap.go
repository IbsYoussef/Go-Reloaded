package textmod

import (
	"strconv"
	"strings"
	"unicode"
)

func cap(s string) string {
	capitalizeWord := func(word string) string {
		// Convert the word to a rune slice to process individual characters
		runes := []rune(word)
		if len(runes) > 0 {
			// Capitalize the first character using Unicode.ToUpper
			runes[0] = unicode.ToUpper(runes[0])
		}

		// Convert the rune slice back to a string and return the capitalized word
		return string(runes)
	}

	// Split the input string into individual words using spaces as delimiters
	sentence := strings.Fields(s)

	// Iterate through each word in the result slice using its index and value
	for i, v := range sentence {
		// Check if the current word is "(cap)"
		if v == "(cap)" {
			// If it is "(cap)", capitalize the previous word (i-1) and set the current word to an empty string to remove it
			sentence[i-1] = capitalizeWord(sentence[i-1])
			sentence[i] = ""
		}

		// Check if the current word is "(cap,"
		if v == "(cap," {
			// If it is "(cap,", capitalize the previous word (i-1) and extract the number after "(cap," from the next word (i+1)
			sentence[i-1] = capitalizeWord(sentence[i-1])
			boundary := len(sentence[i+1])           //Finds the number by finding length and making sure not to gobeyond boundary.
			cap_number := sentence[i+1][:boundary-1] // Extracts Number from (cap)
			num, err := strconv.Atoi(cap_number)
			if err != nil {
				panic(err)
			}

			// Loop through the specified number (num) of previous words and capitalize each of them
			for j := 1; j <= num; j++ {
				sentence[i-j] = capitalizeWord(sentence[i-j])
			}

			// Remove ("(cap,") and number after
			sentence[i], sentence[i+1] = "", ""
		}
	}

	// Join the modified result slice back into a string with spaces between words and return the final string
	return strings.Join(sentence, " ")
}
