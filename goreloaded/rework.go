package main

// import (
// 	"bufio"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// func hexToDec(hex string) string {
// 	dec, err := strconv.ParseInt(hex, 16, 64)
// 	if err != nil {
// 		return ""
// 	}
// 	return strconv.FormatInt(dec, 10)
// }

// func binToDec(bin string) string {
// 	dec, err := strconv.ParseInt(bin, 2, 64)
// 	if err != nil {
// 		return ""
// 	}
// 	return strconv.FormatInt(dec, 10)
// }

// func modifyByNumber(word, action, num string) string {
// 	n, err := strconv.Atoi(num)
// 	if err != nil || n <= 0 {
// 		return ""
// 	}

// 	words := strings.Fields(word)
// 	if len(words) < n {
// 		return ""
// 	}

// 	switch action {
// 	case "low":
// 		words[n-1] = strings.ToLower(words[n-1])
// 	case "up":
// 		for i := n - 1; i < len(words) && i < n-1+n; i++ {

// 			words[i] = strings.ToUpper(words[i])
// 		}
// 	case "cap":
// 		for i := n - 1; i < len(words) && i < n-1+n; i++ {
// 			words[i] = capitalizeWord(words[i])
// 		}
// 	}

// 	copy(words[n-1:], words[n:])
// 	words = words[:len(words)-1]

// 	return strings.Join(words, " ")
// }

// func formatPunctuations(text string) string {
// 	punctuations := []string{".", ",", "!", "?", ":", ";"}

// 	for _, punctuation := range punctuations {
// 		// Add a space after punctuation, if it is not followed by a space
// 		text = strings.ReplaceAll(text, punctuation+" ", punctuation)
// 		text = strings.ReplaceAll(text, " "+punctuation, punctuation+" ")
// 	}

// 	// Adjust spacing for ellipsis (...)
// 	text = strings.ReplaceAll(text, ". . .", "...")
// 	text = strings.ReplaceAll(text, ". .", "..")
// 	text = strings.ReplaceAll(text, " . .", "...")

// 	// Remove excessive spaces
// 	text = strings.ReplaceAll(text, "  ", " ")

// 	text = strings.ReplaceAll(text, " '", "'")
// 	text = strings.ReplaceAll(text, "' ", "'")

// 	text = strings.ReplaceAll(text, "!! ", "!!")
// 	text = strings.ReplaceAll(text, " ...", "...")

// 	return text
// }

// func capitalizeWord(word string) string {
// 	if len(word) == 0 {
// 		return word
// 	}
// 	first := string(unicode.ToUpper(rune(word[0])))
// 	rest := word[1:]
// 	return first + rest
// }

// func modifyText(input string) string {
// 	lines := strings.Split(input, "\n")
// 	var output []string

// 	for _, line := range lines {
// 		words := strings.Fields(line)
// 		var modifiedWords []string
// 		for i := 0; i < len(words); i++ {
// 			word := words[i]
// 			switch {
// 			case word == "(hex)":
// 				if i > 0 {
// 					words[i] = hexToDec(words[i-1])
// 					words[i-1] = ""
// 				}
// 			case word == "(bin)":
// 				if i > 0 {
// 					words[i] = binToDec(words[i-1])
// 					words[i-1] = ""
// 				}
// 			case word == "(up)":
// 				if i > 0 {
// 					words[i] = strings.ToUpper(words[i-1])
// 					words[i-1] = ""
// 				} else {
// 					words[i] = strings.ToUpper(word)
// 				}
// 			case word == "(low)":
// 				if i > 0 {
// 					words[i] = strings.ToLower(words[i-1])
// 					words[i-1] = ""
// 				}
// 			case word == "(cap)":
// 				if i > 0 {
// 					words[i] = capitalizeWord(words[i-1])
// 					words[i-1] = ""
// 				}
// 			case strings.HasPrefix(word, "(low,"):
// 				parts := strings.SplitN(word, ",", 2)
// 				if len(parts) == 2 {
// 					words[i] = modifyByNumber(words[i-1], "(low)", parts[1])
// 					words[i-1] = ""
// 				}
// 			case strings.HasPrefix(word, "(up,"):
// 				parts := strings.SplitN(word, ",", 2)
// 				if len(parts) == 2 {
// 					words[i] = modifyByNumber(words[i-1], "(up)", parts[1])
// 					words[i-1] = ""
// 				}
// 			case strings.HasPrefix(word, "(cap,"):
// 				parts := strings.SplitN(word, ",", 2)
// 				if len(parts) == 2 {
// 					words[i] = modifyByNumber(words[i-1], "(cap)", parts[1])
// 					words[i-1] = ""
// 				}
// 			}
// 		}

// 		// Remove empty words and join modified words
// 		for _, word := range words {
// 			if word != "" {
// 				modifiedWords = append(modifiedWords, word)
// 			}
// 		}

// 		line = strings.Join(modifiedWords, " ")
// 		line = formatPunctuations(line)

// 		output = append(output, line)
// 	}

// 	return strings.Join(output, "\n")
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		log.Fatal("Please provide an input file and an output file.")
// 	}

// 	inputFile := args[0]
// 	outputFile := args[1]

// 	file, err := os.Open(inputFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var lines []string
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	modifiedText := modifyText(strings.Join(lines, "\n"))

// 	output, err := os.Create(outputFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer output.Close()

// 	writer := bufio.NewWriter(output)
// 	for _, line := range strings.Split(modifiedText, "\n") {
// 		_, err := writer.WriteString(line + "\n")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	writer.Flush()

// 	log.Println("Modified text has been written to", outputFile)
// }
