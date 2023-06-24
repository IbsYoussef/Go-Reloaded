package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Hexadecimal to Decimal conversion
func hexToDec(hex string) string {
	dec := 0
	base := 1

	for i := len(hex) - 1; i >= 0; i-- {
		if hex[i] >= '0' && hex[i] <= '9' {
			dec += int(hex[i]-'0') * base
		} else if hex[i] >= 'A' && hex[i] <= 'F' {
			dec += int(hex[i]-'A'+10) * base
		} else if hex[i] >= 'a' && hex[i] <= 'f' {
			dec += int(hex[i]-'a'+10) * base
		}
		base *= 16
	}

	return fmt.Sprintf("%d", dec)
}

// Binary to Decimal conversion
func binToDec(bin string) string {
	dec := 0
	base := 1

	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			dec += base
		}
		base *= 2
	}

	return fmt.Sprintf("%d", dec)
}

// Applies the requested modifications to the text
func applyModifications(text string) string {
	re := regexp.MustCompile(`\(([a-zA-Z]+)(?:,\s*(\d+))?\)`)

	words := strings.Fields(text)
	modifiedWords := make([]string, 0, len(words))

	for _, word := range words {
		matches := re.FindStringSubmatch(word)

		if len(matches) > 0 {
			switch matches[1] {
			case "hex":
				modifiedWords = append(modifiedWords, hexToDec(matches[2]))
			case "bin":
				modifiedWords = append(modifiedWords, binToDec(matches[2]))
			case "up":
				if len(matches) == 3 {
					n := min(len(modifiedWords), atoi(matches[2]))
					modifiedWords = append(modifiedWords[:len(modifiedWords)-n], toUpper(modifiedWords[len(modifiedWords)-n:])...)
				} else {
					modifiedWords = append(modifiedWords, toUpper([]string{modifiedWords[len(modifiedWords)-1]})...)
				}
			case "low":
				if len(matches) == 3 {
					n := min(len(modifiedWords), atoi(matches[2]))
					modifiedWords = append(modifiedWords[:len(modifiedWords)-n], toLower(modifiedWords[len(modifiedWords)-n:])...)
				} else {
					modifiedWords = append(modifiedWords, toLower([]string{modifiedWords[len(modifiedWords)-1]})...)
				}
			case "cap":
				if len(matches) == 3 {
					n := min(len(modifiedWords), atoi(matches[2]))
					modifiedWords = append(modifiedWords[:len(modifiedWords)-n], capitalize(modifiedWords[len(modifiedWords)-n:])...)
				} else {
					modifiedWords = append(modifiedWords, capitalize([]string{modifiedWords[len(modifiedWords)-1]})...)
				}
			}
		} else if strings.HasPrefix(word, "'") && strings.HasSuffix(word, "'") {
			word = strings.Trim(word, "'")
			modifiedWords = append(modifiedWords, "'"+word+"'")
		} else if strings.HasSuffix(word, "...") || strings.HasSuffix(word, "!?") {
			modifiedWords = append(modifiedWords, word)
		} else if word == "a" && len(modifiedWords) < len(words)-1 && isVowelOrH(words[len(modifiedWords)+1]) {
			modifiedWords = append(modifiedWords, "an")
		} else {
			modifiedWords = append(modifiedWords, word)
		}
	}

	return strings.Join(modifiedWords, " ")
}

// Helper functions
func toUpper(words []string) []string {
	modifiedWords := make([]string, len(words))
	for i, word := range words {
		modifiedWords[i] = strings.ToUpper(word)
	}
	return modifiedWords
}

func toLower(words []string) []string {
	modifiedWords := make([]string, len(words))
	for i, word := range words {
		modifiedWords[i] = strings.ToLower(word)
	}
	return modifiedWords
}

func capitalize(words []string) []string {
	modifiedWords := make([]string, len(words))
	for i, word := range words {
		modifiedWords[i] = strings.Title(word)
	}
	return modifiedWords
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func atoi(s string) int {
	val := 0
	for _, ch := range s {
		val = val*10 + int(ch-'0')
	}
	return val
}

func isVowelOrH(ch string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	for _, vowel := range vowels {
		if ch == vowel {
			return true
		}
	}
	return false
}

// Read input file, apply modifications, and write output file
func processFile(inputFile, outputFile string) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := applyModifications(line)
		fmt.Fprintln(output, modifiedLine)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := processFile(inputFile, outputFile)
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Println("Text modification completed successfully.")
}
