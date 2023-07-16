package main

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	firstLetter := rune(word[0])
	capitalized := string(unicode.ToUpper(firstLetter))
	return capitalized + word[1:]
}
func cap(sentence string) string {
	words := strings.Split(sentence, " ")
	capCount := 1 // Default value of 1 if no number is specified

	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(cap") && strings.HasSuffix(words[i], ")") {
			parts := strings.Split(words[i], ",")
			if len(parts) == 2 {
				wordCount := strings.TrimSpace(parts[1])
				var count int
				if _, err := fmt.Sscanf(wordCount, "%d", &count); err == nil && count > 0 {
					capCount = count
				}
			}

			startIndex := max(i-capCount, 0)
			for j := i - 1; j >= startIndex; j-- {
				words[j] = capitalize(words[j])
			}

			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to recheck the current index
		}
	}

	return strings.Join(words, " ")
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Lowercase(word string) string {
	return strings.ToLower(word)
}
func low(sentence string) string {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		if words[i] == "(low)" {
			words[i-1] = Lowercase(words[i-1])
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
		words[i] = capitalize(words[i])
		count--
	}
}
func Count(sentence string) string {
	words := strings.Fields(sentence)
	conversionMap := make(map[int]string)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			conversion := strings.Trim(words[i], "()")
			parts := strings.Split(conversion, ",")
			if len(parts) == 2 {
				wordCount := strings.TrimSpace(parts[1])
				var count int
				if _, err := fmt.Sscanf(wordCount, "%d", &count); err == nil && count > 0 {
					conversionType := strings.TrimSpace(parts[0])
					conversionMap[i] = conversionType
					if i-count >= 0 {
						delete(conversionMap, i-count)
					}
				}
			}
		}
	}
	for index, conversionType := range conversionMap {
		switch conversionType {
		case "low":
			convertToLowercaseByCount(words[index:], index-len(conversionMap)+1)
		case "up":
			convertToUppercaseByCount(words[index:], index-len(conversionMap)+1)
		case "cap":
			capitalizeWordsByCount(words[index:], index-len(conversionMap)+1)
		}
	}
	return strings.Join(words, " ")
}

func main() {

}
