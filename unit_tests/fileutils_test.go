package unit_tests

import (
	utils "go-reloaded/internal/utils"
	"io"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testinput*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Ensure the file is deleted after the test

	// Write test content to the temporary file
	testContent := "This is a test\n"
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	// Now test the ReadFile function
	content, err := utils.ReadFile(tmpFile.Name())
	if err != nil {
		t.Errorf("ReadFile returned an error: %v", err)
	}

	if content != testContent {
		t.Errorf("Expected content %q but got %q", testContent, content)
	}
}

func TestWriteFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testoutput*.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testContent := "Writing test content to file\n"

	// Test the WriteFile function
	err = utils.WriteFile(tmpFile.Name(), testContent)
	if err != nil {
		t.Fatalf("WriteFile returned an error: %v", err)
	}

	// Read the file back and verify the content
	file, err := os.Open(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to open file after writing: %v", err)
	}
	defer file.Close()

	// Read the content
	writtenContent, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file content: %v", err)
	}

	if string(writtenContent) != testContent {
		t.Errorf("Expected content %q but got %q", testContent, string(writtenContent))
	}
}
