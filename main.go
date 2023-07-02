package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var punctuationRegex = regexp.MustCompile(`([.,!?;:])\s*`)
var ellipsisRegex = regexp.MustCompile(`\.{3}`)
var exclamationRegex = regexp.MustCompile(`[!?]`)
var quoteRegex = regexp.MustCompile(`'([^']*)'`)
var vowelRegex = regexp.MustCompile(`(?i)\ba\s([aeiouh])`)

func ReplaceHex(text string) string {
	hexRegex := regexp.MustCompile(`\b(0x[0-9A-Fa-f]+)\b`)
	return hexRegex.ReplaceAllStringFunc(text, func(hex string) string {
		dec, err := strconv.ParseInt(hex[2:], 16, 64)
		if err != nil {
			log.Println("Error converting hex to decimal:", err)
			return hex
		}
		return strconv.FormatInt(dec, 10)
	})
}

func ReplaceBin(text string) string {
	binRegex := regexp.MustCompile(`\b(0b[01]+)\b`)
	return binRegex.ReplaceAllStringFunc(text, func(bin string) string {
		dec, err := strconv.ParseInt(bin[2:], 2, 64)
		if err != nil {
			log.Println("Error converting bin to decimal:", err)
			return bin
		}
		return strconv.FormatInt(dec, 10)
	})
}

func capitalizeWord(word string, caseType rune) string {
	switch caseType {
	case 'u':
		return strings.ToUpper(word)
	case 'l':
		return strings.ToLower(word)
	case 'c':
		return strings.ToUpper(string(word[0])) + word[1:]
	default:
		return word
	}
}

func ReplaceCase(text string) string {
	return quoteRegex.ReplaceAllStringFunc(text, func(quote string) string {
		quote = strings.Trim(quote, "'")
		parts := strings.Split(quote, " ")
		if len(parts) != 2 {
			return quote
		}

		mod, word := parts[0], parts[1]
		if len(mod) == 0 || len(word) == 0 {
			return quote
		}

		// Extract the case type and word count (if provided)
		caseType := rune(mod[0])
		wordCount := -1
		if len(mod) > 1 && unicode.IsDigit(rune(mod[1])) {
			wordCount = int(mod[1] - '0')
		}

		// Capitalize the word
		words := strings.Fields(word)
		if wordCount > len(words) || wordCount < 0 {
			wordCount = len(words)
		}

		for i := 0; i < wordCount; i++ {
			words[i] = capitalizeWord(words[i], caseType)
		}

		return strings.Join(words, " ")
	})
}

func formatPunctuation(text string) string {
	text = punctuationRegex.ReplaceAllString(text, "$1 ")
	text = ellipsisRegex.ReplaceAllString(text, "...")
	text = exclamationRegex.ReplaceAllStringFunc(text, func(exclamation string) string {
		if len(exclamation) > 1 {
			return exclamation
		}
		return exclamation + exclamation
	})
	return text
}

func formatQuotes(text string) string {
	return quoteRegex.ReplaceAllString(text, "'$1'")
}

func replaceVowels(text string) string {
	return vowelRegex.ReplaceAllStringFunc(text, func(match string) string {
		return strings.Replace(match, " a ", " an ", -1)
	})
}

func ProcessText(text string) string {
	text = ReplaceCase(text)
	text = ReplaceHex(text)
	text = ReplaceBin(text)
	text = formatPunctuation(text)
	text = formatQuotes(text)
	text = replaceVowels(text)
	return text
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-file> <output-file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading input file:", err)
	}

	// Process text
	var processedLines []string
	for _, line := range lines {
		processedLines = append(processedLines, ProcessText(line))
	}

	// Write output file
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range processedLines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatal("Error writing to output file:", err)
		}
	}
	writer.Flush()

	fmt.Println("Text processed and saved to", outputFile)
}
