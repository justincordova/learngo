package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Processing files...")
	fmt.Println()

	// Try to process different file types
	files := []struct {
		name    string
		context string
	}{
		{"config.json", "configuration"},
		{"data.txt", "data file"},
		{"settings.yaml", "settings"},
	}

	for _, f := range files {
		if err := processFile(f.name, f.context); err != nil {
			fmt.Printf("Error processing %s:\n%v\n\n", f.name, err)
		}
	}
}

// readFile attempts to read a file and wraps any error with the filename
func readFile(filename string) error {
	_, err := os.ReadFile(filename)
	if err != nil {
		// Wrap the error with context about which file failed
		return fmt.Errorf("failed to read file %q: %w", filename, err)
	}
	return nil
}

// processFile processes a file and adds additional context to any errors
func processFile(filename, context string) error {
	if err := readFile(filename); err != nil {
		// Add another layer of context about what was being processed
		return fmt.Errorf("failed to process %s: %w", context, err)
	}
	return nil
}
