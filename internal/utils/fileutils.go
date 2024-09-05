package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(filepath string) (string, error) {
	// Open file and handle any errors if unsuccessful
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return "", err
	}
	defer file.Close() // Close file after

	// Create a new scanner to read the contents of the file
	// Create a string builder to accumulate the content read from file
	scanner := bufio.NewScanner(file)
	var content strings.Builder

	// As scanner reads the file we append the content to the string builder in a neat format
	for scanner.Scan() {
		content.WriteString(scanner.Text())
		content.WriteString("\n")
	}

	// Check if the scanner read the file successfully
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading the file: %w", err)
	}

	// Return the content read and a nil error indicating success
	return content.String(), nil
}

func WriteFile(filepath string, content string) error {
	// Open and create the file
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Created a buffered writer that will write content to the file
	writer := bufio.NewWriter(file)

	// Write content to the buffer, not yet to file
	_, err = writer.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	// Flush buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// Return a nil error indicating success
	return nil
}
