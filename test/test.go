package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func convertToLowercaseByCount(words []string, count int) {
	for i := len(words) - 1; i >= 0 && count > 0; i-- {
		words[i] = strings.ToLower(words[i])
		count--
	}
}

func convertToUppercaseByCount(words []string, count int) {
	for i := len(words) - 1; i >= 0 && count > 0; i-- {
		words[i] = strings.ToUpper(words[i])
		count--
	}
}

func capitalizeWordsByCount(words []string, count int) {
	for i := len(words) - 1; i >= 0 && count > 0; i-- {
		words[i] = capitalizeFirstLetter(words[i])
		count--
	}
}

func Count(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			conversion := strings.Trim(words[i], "()")
			parts := strings.Split(conversion, ",")
			if len(parts) == 2 {
				wordCount := strings.TrimSpace(parts[1])
				var count int
				if _, err := fmt.Sscanf(wordCount, "%d", &count); err == nil && count > 0 {
					switch strings.TrimSpace(parts[0]) {
					case "low":
						convertToLowercaseByCount(words[i+1:], count)
						words = append(words[:i], words[i+count+1:]...)
					case "up":
						convertToUppercaseByCount(words[i+1:], count)
						words = append(words[:i], words[i+count+1:]...)
					case "cap":
						capitalizeWordsByCount(words[i+1:], count)
						words = append(words[:i], words[i+count+1:]...)
					}
				}
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {

	text := "Hello there (cap) my name (cap) is ibraheem (cap)"
	text2 := "Hello there (cap) my name (cap) is miguel (cap)"
	text3 := "i (cap) am testing the (cap) cap (cap) function (cap) to see if, (cap) it works. (cap)"
	fmt.Println(cap(text))
	fmt.Println(cap(text2))
	fmt.Println(cap(text3))

	text4 := "CONVERTING (low) to converting."
	text5 := "TESTING, (low) the low function TO (low) see if it WORKS. (low)"
	fmt.Println(low(text4))
	fmt.Println(low(text5))

	text6 := "Testing (up) the upper case (up) function now (up)"
	fmt.Println(up(text6))

	text7 := "HI THERE, (low, 2)"
	fmt.Println(Count(text7))

}
