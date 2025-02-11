package main

import (
	"fmt"
	"go-reloaded/internal/textmod"
	"go-reloaded/internal/utils"
	"os"
	"path/filepath"
)

func main() {
	// Get command-line arguments
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf("Invalid input format, usage: go run . <input_filename.txt> <output_filename.txt>")
		return
	}

	// Get the filenames from arguements
	inputFile := args[0]
	outputFile := args[1]

	// Automatically prepend the default folder paths
	// Input files are expected to be in the "text-files" directory.
	// Output files will be written in the "output" directory.
	inputPath := filepath.Join("text-files", inputFile)
	outputPath := filepath.Join("outputs", outputFile)

	// Resolve absolute paths for both input and output
	inputPath, err := filepath.Abs(inputPath)
	if err != nil {
		fmt.Printf("Error resolving input file path: %v\n", err)
		return
	}
	outputPath, err = filepath.Abs(outputPath)
	if err != nil {
		fmt.Printf("Error resolving output file path: %v\n", err)
		return
	}

	// Read the content from the input file.
	inputContent, err := utils.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputPath, err)
		return
	}

	// Validate that the input file is not empty.
	if inputContent == "" {
		fmt.Println("Input file is empty or the format is invalid.")
		return
	}

	// Apply modifications to text content
	inputContent = textmod.ModifyText(inputContent)

	// Write the processed content to the output file.
	err = utils.WriteFile(outputPath, inputContent)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", outputPath, err)
		return
	}

	fmt.Println("File processed successfully.")

}
