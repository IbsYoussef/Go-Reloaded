package main

import (
	"strconv"
	"strings"
	"unicode"
)

func hexToDec(hex string) string {
	dec, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(dec, 10)
}

func binToDec(bin string) string {
	dec, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(dec, 10)
}

func capitalizeFirstLetter(word string) string {
	if len(word) == 0 {
		return word
	}

	firstLetter := rune(word[0])
	capitalized := string(unicode.ToUpper(firstLetter))
	return capitalized + word[1:]
}

func cap(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = capitalizeFirstLetter(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to recheck the current index
		}
	}
	return strings.Join(words, " ")
}

func convertToLowercase(word string) string {
	return strings.ToLower(word)
}
func low(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = convertToLowercase(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func convertToUppercase(word string) string {
	return strings.ToUpper(word)
}
func up(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = convertToUppercase(words[i-1])
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func main() {

}
