package textmod

import "strings"

func ReplacePunt(s string) string {
	// quotes is a function to handle single quotes
	quotes := func(s string) string {
		str := ""
		var removeSpace bool // default false

		// Loop through each character in the input string
		for i, char := range s {
			// Check if the character is a single quote (ASCII 39) and the previous character is a space
			if char == 39 && s[i-1] == ' ' {
				if removeSpace {
					// If removeSpace is true, remove the last character (space) from the result
					str = str[:len(str)-1]
					// Add the current character (single quote) to the result
					str = str + string(char)
					removeSpace = false // Reset removeSpace to false
				} else {
					// If removeSpace is false, add the current character (single quote) to the result
					str = str + string(char)
					removeSpace = true // Set removeSpace to true
				}
			} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
				// Check if the previous two characters are a single quote (ASCII 39) and a space
				if removeSpace {
					// If removeSpace is true, remove the last character (space) from the result
					str = str[:len(str)-1]
					// Add the current character to the result
					str = str + string(char)
				} else {
					// If removeSpace is false, add the current character to the result
					str = str + string(char)
				}
			} else {
				// For all other cases, add the current character to the result
				str = str + string(char)
			}
		}
		return str
	}

	strr := ""
	result := strings.Fields(s)
	for _, v := range result {
		strr += v + " "
	}

	// The code above removes the additional spaces from the input string and
	// prepares it for further processing.

	result1 := strings.Fields(strr)
	str := ""
	for _, v := range result1 {
		str += v + " "
	}

	// The code above removes the additional spaces again and
	// prepares the string for punctuation replacement.

	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			// Check if it's the last character in the string
			if char == '.' || char == ',' || char == ':' || char == ';' || char == '!' || char == '?' {
				// Check if the last character is a punctuation mark
				if str[i-1] == ' ' {
					// If there is a space before the punctuation mark, remove the space from the result
					word = word[:len(word)-1]
					word = word + string(char)
				} else {
					// If there is no space before the punctuation mark, add the punctuation mark to the result
					word = word + string(char)
				}
			} else {
				// For non-punctuation marks, add the character to the result
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == ':' || char == ';' || char == '!' || char == '?' {
			// Check if the character is a punctuation mark
			if str[i-1] == ' ' {
				// If there is a space before the punctuation mark, remove the space from the result
				word = word[:len(word)-1]
				word = word + string(char)
			} else {
				// If there is no space before the punctuation mark, add the punctuation mark to the result
				word = word + string(char)
			}

			// Check if the next character is not a space or another punctuation mark
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				// If it is not, add a space after the punctuation mark to separate words
				word = word + " "
			}
		} else {
			// For non-punctuation marks, add the character to the result
			word = word + string(char)
		}
	}

	// Call the quotes function to handle single quotes and return the final processed string
	return quotes(word)
}
