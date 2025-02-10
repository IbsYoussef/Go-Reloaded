package main

import (
	"fmt"
	"go-reloaded/internal/textmod"
	"go-reloaded/internal/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Printf("Invalid input format, usage: go run . <input_filename.txt> <output_filename.txt>")
		return
	}

	inputFile := args[0]
	outputFile := args[1]

	inputContent, err := utils.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	if inputContent == "" {
		fmt.Println("Input file is empty or the format is invalid.")
		return
	}

	inputContent = textmod.ModifyText(inputContent)

	err = utils.WriteFile(outputFile, inputContent)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}

	fmt.Println("File processed successfully.")

}
