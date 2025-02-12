package main

import (
	"go-reloaded/textmod"
	"os"
)

func main() {

	// Reads input file
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	text := string(file)
	text = textmod.Up(text)
	text = textmod.Cap(text)
	text = textmod.Low(text)
	text = textmod.HexandBin(text)
	text = textmod.ReplaceAWithAn(text)
	text = textmod.ReplacePunt(text)
	/*
	   The file permission 0o644 (or 420 in decimal) means that the file can be read and written by the owner of the file,
	   and only read by others.
	   Text written to new file.
	*/
	err = os.WriteFile("output.txt", []byte(text), 0o644)
	if err != nil {
		panic(err)
	}

}
