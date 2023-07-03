package main

import (
	"bufio"
	"log"
	"os"
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

func formatPunctuations(text string) string {
	punctuations := []string{".", ",", "!", "?", ":", ";"}

	for _, punctuation := range punctuations {
		text = strings.ReplaceAll(text, " "+punctuation, punctuation)
	}

	text = strings.ReplaceAll(text, " '", "'")
	text = strings.ReplaceAll(text, "' ", "'")

	text = strings.ReplaceAll(text, "...", "...")
	text = strings.ReplaceAll(text, "!", "!")
	text = strings.ReplaceAll(text, "?", "?")
	text = strings.ReplaceAll(text, ": ", ":")
	text = strings.ReplaceAll(text, "; ", ";")

	return text
}

func modifyByNumber(word, action, num string) string {
	n, err := strconv.Atoi(num)
	if err != nil {
		return ""
	}

	words := strings.Fields(word)
	if len(words) == 0 || n >= len(words) {
		return ""
	}

	switch action {
	case "low":
		words[n-1] = strings.ToLower(words[n-1])
	case "up":
		words[n-1] = strings.ToUpper(words[n-1])
	case "cap":
		words[n-1] = capitalizeWord(words[n-1])

	}

	return strings.Join(words, " ")
}

func capitalizeWord(word string) string {
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func modifyText(input string) string {
	lines := strings.Split(input, "\n")
	var output []string

	for _, line := range lines {
		words := strings.Fields(line)
		for i, word := range words {
			switch {
			case word == "(hex)":
				words[i] = hexToDec(words[i-1])
				words[i-1] = ""
			case word == "(bin)":
				words[i] = binToDec(words[i-1])
				words[i-1] = ""
			case word == "(up)":
				words[i] = strings.ToUpper(words[i-1])
				words[i-1] = ""
			case word == "(low)":
				words[i] = strings.ToLower(words[i-1])
				words[i-1] = ""
			case word == "(cap)":
				words[i] = capitalizeWord(words[i-1])
				words[i-1] = ""
			case strings.HasPrefix(word, "(low,"):
				parts := strings.SplitN(word, ",", 2)
				if len(parts) == 2 {
					num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil && num > 0 && num <= i {
						words[i-num] = strings.ToLower(words[i-num])
						words[i-1] = ""
					}
				}
			case strings.HasPrefix(word, "(up,"):
				parts := strings.SplitN(word, ",", 2)
				if len(parts) == 2 {
					num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil && num > 0 && num <= i {
						words[i-num] = strings.ToUpper(words[i-num])
						words[i-1] = ""
					}
				}
			case strings.HasPrefix(word, "(cap,"):
				parts := strings.SplitN(word, ",", 2)
				if len(parts) == 2 {
					num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil && num > 0 && num <= i {
						words[i-num] = capitalize(words[i-num])
						words[i-1] = ""
					}
				}
			}
		}

		line = strings.Join(words, " ")
		line = formatPunctuations(line)

		output = append(output, line)
	}

	return strings.Join(output, "\n")
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return string(unicode.ToUpper(rune(word[0]))) + word[1:]
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("Please provide an input file and an output file.")
	}

	inputFile := args[0]
	outputFile := args[1]

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	modifiedText := modifyText(strings.Join(lines, "\n"))

	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	writer := bufio.NewWriter(output)
	for _, line := range strings.Split(modifiedText, "\n") {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	writer.Flush()

	log.Println("Modified text has been written to", outputFile)
}
