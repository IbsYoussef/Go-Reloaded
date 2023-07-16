package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cap(s string) string {
	capitalizeWord := func(word string) string {
		if len(word) == 0 {
			return word
		}
		return strings.ToUpper(string(word[0])) + word[1:]
	}

	words := strings.Fields(s) // Split the string into individual words

	for i := 0; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			prevWord := words[i-1]
			if len(prevWord) > 0 {
				words[i-1] = capitalizeWord(prevWord)
			}

			// Remove the instance of (cap)
			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to recheck the current position
		} else if strings.HasPrefix(words[i], "(cap,") && strings.HasSuffix(words[i], ")") {
			// Handle (cap,<number>) instance
			// Extracting the number from the string
			numStr := words[i][5 : len(words[i])-1]
			num := 0
			if parsedNum, err := strconv.Atoi(numStr); err == nil {
				num = parsedNum
			}

			// Capitalize the specified number of words before (cap)
			for j := 0; j < num && i-j-1 >= 0; j++ {
				prevWord := words[i-j-1]
				if len(prevWord) > 0 {
					words[i-j-1] = capitalizeWord(prevWord)
				}
			}

			// Remove the instance of (cap,<number>)
			words = append(words[:i], words[i+1:]...)
			i-- // Decrement i to recheck the current position
		}
	}

	// Reconstruct the string with modified words
	return strings.Join(words, " ")
}
func up(s string) string {
	result := strings.Fields(s)
	for i, v := range result {
		if v == "(up)" {
			result[i-1] = strings.ToUpper(result[i-1])
			result[i] = ""
		}
		if v == "(up," {
			result[i-1] = strings.ToUpper(result[i-1])
			twov := len(result[i+1])
			numb := result[i+1][:twov-1] // 2
			nu, err := strconv.Atoi(numb)
			if err != nil {
				panic(err)
			}

			for j := 1; j <= nu; j++ {
				result[i-j] = strings.ToUpper(result[i-j])
			}
			result[i], result[i+1] = "", ""
		}
	}
	return strings.Join(result, " ")
}
func low(s string) string {
	result := strings.Fields(s)
	for i, v := range result {
		if v == "(low)" {
			result[i-1] = strings.ToLower(result[i-1])
			result[i] = ""
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
			result[i], result[i+1] = "", ""
		}
	}
	return strings.Join(result, " ")
}

func HexandBin(s string) string {
	result := strings.Fields(s)
	for i, v := range result {
		if v == "(hex)" {
			h, err := strconv.ParseInt(result[i-1], 16, 64)
			result[i-1] = fmt.Sprint(h)
			result[i] = ""
			if err != nil {
				panic(err)
			}
		}
		if v == "(bin)" {
			j, err := strconv.ParseInt(result[i-1], 2, 64)
			result[i-1] = fmt.Sprint(j)
			result[i] = ""
			if err != nil {
				panic(err)
			}
		}
	}
	return strings.Join(result, " ")
}

func quotes(s string) string {
	str := ""
	var removeSpace bool // default false
	for i, char := range s {
		if char == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
				removeSpace = false
			} else {
				str = str + string(char)
				removeSpace = true
			}
		} else if i > 1 && s[i-2] == 39 && s[i-1] == ' ' {
			if removeSpace {
				str = str[:len(str)-1]
				str = str + string(char)
			} else {
				str = str + string(char)
			}
		} else {
			str = str + string(char)
		}
	}
	return str
}
func ReplacePunt(s string) string {
	strr := ""
	result := strings.Fields(s)
	for _, v := range result {
		strr += v + " "
	}
	result1 := strings.Fields(strr)
	str := ""
	for _, v := range result1 {
		str += v + " "
	}
	word := ""
	for i, char := range str {
		if i == len(str)-1 {
			if char == '.' || char == ',' || char == ':' || char == ';' || char == '!' || char == '?' {
				if str[i-1] == ' ' {
					word = word[:len(word)-1] // end of paragraph avoidance of space after the full stop
					word = word + string(char)
				} else {
					word = word + string(char)
				}
			} else {
				word = word + string(char)
			}
		} else if char == '.' || char == ',' || char == ':' || char == ';' || char == '!' || char == '?' {
			if str[i-1] == ' ' {
				word = word[:len(word)-1] // removes blank space prior to character
				word = word + string(char)
			} else {
				word = word + string(char)
			}
			if str[i+1] != ' ' && str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ';' && str[i+1] != ':' {
				word = word + " " // adds a space after character
			}
		} else {
			word = word + string(char)
		}
	}
	return quotes(word)
}

func first_rune(s string) string {
	vocal := []rune(s)
	return string(vocal[0])
}

func ReplaceAwhitAn(s string) string {
	str := ""
	result := strings.Fields(s)
	for i, v := range result {
		str += v + " "
		if v == "a" && first_rune(result[i+1]) == "a" || v == "a" && first_rune(result[i+1]) == "e" || v == "a" && first_rune(result[i+1]) == "i" || v == "a" && first_rune(result[i+1]) == "o" || v == "a" && first_rune(result[i+1]) == "u" || v == "a" && first_rune(result[i+1]) == "h" {
			result[i] = "An"
		}
	}
	return strings.Join(result, " ")
}

func main() {

	sentence := "There it was. A amazing rock!"
	fmt.Println(ReplaceAwhitAn(sentence))

}
