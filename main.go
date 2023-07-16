package main

import (
	"fmt"
	"os"
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

func up(s string) string {
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

func low(s string) string {
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

func HexandBin(s string) string {
	// Split the input string into individual words using spaces as delimiters
	sentence := strings.Fields(s)

	// Iterate through each word in the result slice using its index and value
	for i, v := range sentence {
		// Check if the current word is "(hex)"
		if v == "(hex)" {
			// Attempt to parse the previous word as a hexadecimal integer (base 16)
			hex, err := strconv.ParseInt(sentence[i-1], 16, 64)

			// Replace the previous word with the parsed hexadecimal value as a decimal string
			sentence[i-1] = fmt.Sprint(hex)

			// Set the current word to an empty string to remove it from the result, i.e Remove (hex)
			sentence[i] = ""

			// Handle any error that occurred during the conversion
			if err != nil {
				panic(err)
			}
		}

		// Check if the current word is "(bin)"
		if v == "(bin)" {
			// Attempt to parse the previous word as a binary integer (base 2)
			bin, err := strconv.ParseInt(sentence[i-1], 2, 64)

			// Replace the previous word with the parsed binary value as a decimal string
			sentence[i-1] = fmt.Sprint(bin)

			// Set the current word to an empty string to remove it from the result, i.e Remove (bin)
			sentence[i] = ""

			// Handle any error that occurred during the conversion
			if err != nil {
				panic(err)
			}
		}
	}

	// Join the modified result slice back into a string with spaces between words and return the final string
	return strings.Join(sentence, " ")
}

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

func main() {

	// Reads input file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(file)
	text = up(text)
	text = cap(text)
	text = low(text)
	text = HexandBin(text)
	text = ReplaceAWithAn(text)
	text = ReplacePunt(text)
	/*
	   The file permission 0o644 (or 420 in decimal) means that the file can be read and written by the owner of the file,
	   and only read by others.
	   Text written to new file.
	*/
	err = os.WriteFile("output.txt", []byte(text), 0o644)
	if err != nil {
		panic(err)
	}

}
