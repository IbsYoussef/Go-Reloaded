package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// Configuration represents a modification configuration.
type Configuration struct {
	Search  string
	Replace string
	Number  int
}

var (
	configurations = []Configuration{
		{Search: "(hex)", Replace: "${dec}"},
		{Search: "(bin)", Replace: "${dec}"},
		{Search: "(up)", Replace: "${up}"},
		{Search: "(low)", Replace: "${low}"},
		{Search: "(cap)", Replace: "${cap}"},
	}

	punctuationRegex = regexp.MustCompile(`([.,!?;:])\s*`)
	ellipsisRegex    = regexp.MustCompile(`\.{3}`)
	exclamationRegex = regexp.MustCompile(`[!?]`)
	quoteRegex       = regexp.MustCompile(`'([^']*)'`)
	vowelRegex       = regexp.MustCompile(`(?i)\ba\s([aeiouh])`)
)

// ReplaceHex replaces hexadecimal numbers with their decimal equivalent.
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

// ReplaceBin replaces binary numbers with their decimal equivalent.
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

// capitalizeWord capitalizes a word based on the specified case.
func capitalizeWord(word string, caseType rune) string {
	switch caseType {
	case 'u':
		return strings.ToUpper(word)
	case 'l':
		return strings.ToLower(word)
	case 'c':
		return strings.Title(strings.ToLower(word))
	default:
		return word
	}
}

// ReplaceCase replaces words based on case modification configurations.
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

// FormatPunctuation formats punctuation marks in the text.
func FormatPunctuation(text string) string {
	text = punctuationRegex.ReplaceAllString(text, "$1 ")
	text = ellipsisRegex.ReplaceAllString(text, "...")
	text = exclamationRegex.ReplaceAllString(text, "")
	return text
}

// ReplaceVowelA replaces "a" with "an" if the next word begins with a vowel or "h".
func ReplaceVowelA(text string) string {
	return vowelRegex.ReplaceAllString(text, " an $1")
}

// ApplyModifications applies all the required modifications to the given text.
func ApplyModifications(text string) string {
	text = ReplaceHex(text)
	text = ReplaceBin(text)
	text = ReplaceCase(text)
	text = FormatPunctuation(text)
	text = ReplaceVowelA(text)
	return text
}

// ProcessFile reads the input file, applies modifications, and writes the result to the output file.
func ProcessFile(inputFile, outputFile string) error {
	input, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer input.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer output.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := ApplyModifications(line)
		_, err := fmt.Fprintln(output, modifiedLine)
		if err != nil {
			return fmt.Errorf("failed to write to output file: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error scanning input file: %w", err)
	}

	return nil
}

// TestProcessFile tests the ProcessFile function.
func TestProcessFile(t *testing.T) {
	inputFile := "sample.txt"
	outputFile := "result.txt"

	err := ProcessFile(inputFile, outputFile)
	if err != nil {
		t.Errorf("ProcessFile returned an error: %v", err)
	}

	expectedFile := "expected.txt"
	expectedLines, err := readLines(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read expected file: %v", err)
	}

	resultLines, err := readLines(outputFile)
	if err != nil {
		t.Fatalf("Failed to read result file: %v", err)
	}

	if len(expectedLines) != len(resultLines) {
		t.Errorf("Mismatched line count. Expected: %d, Got: %d", len(expectedLines), len(resultLines))
	}

	for i := range expectedLines {
		if expectedLines[i] != resultLines[i] {
			t.Errorf("Mismatched content at line %d. Expected: %s, Got: %s", i+1, expectedLines[i], resultLines[i])
		}
	}
}

// readLines reads all lines from a file and returns them as a slice of strings.
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return lines, nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input-file> <output-file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := ProcessFile(inputFile, outputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modifications complete.")
}
