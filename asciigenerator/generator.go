package asciigenerator

import (
	"errors"
	"strings"
)

// GenerateAsciiArt takes input text and a banner name, then returns the ASCII representation.
func GenerateAsciiArt(input string, banner string) (string, error) {
	// Construct the path and load the banner file content
	bannerPath := "banners/" + banner + ".txt"
	bannerText, err := readFile(bannerPath) // Assumes readFile is implemented in utils.go
	if err != nil {
		return "", errors.New("500: Could not read banner file")
	}

	// Normalize line endings for consistent processing
	bannerText = strings.ReplaceAll(bannerText, "\r", "")
	bannerLines := strings.Split(bannerText, "\n")

	// Validate input length and handle empty inputs
	if input == "" {
		return "", nil
	}
	if len(input) > 10000 {
		return "", errors.New("400: Input size exceeds limit")
	}

	// Handle input normalization
	input = strings.ReplaceAll(input, "\r", "")

	// Helper logic for inputs that are only newline characters
	if isOnlyNewline(input) { // Assumes isOnlyNewline is implemented in utils.go
		return input, nil
	}

	inputLines := strings.Split(input, "\n")
	var result string
	const (
		AsciiOffset = 32 // ASCII printable characters start at 32
		CharHeight  = 9  // Height of each ASCII character (8 lines + 1 separator)
	)

	for _, line := range inputLines {
		// Ensure only allowed ASCII characters are processed
		if !isValidInput(line) { // Assumes isValidInput is implemented in utils.go
			return "", errors.New("400: Only ASCII characters 32-126 are allowed")
		}

		// Handle empty lines within multi-line input
		if line == "" {
			result += "\n"
			continue
		}

		// Process the 8 lines height for each character in the current line
		for height := 1; height < 9; height++ {
			for _, char := range line {
				// Calculate index in the banner slice
				asciiStartLine := (int(char) - AsciiOffset) * CharHeight
				lineIndex := asciiStartLine + height
				result += bannerLines[lineIndex]
			}
			result += "\n"
		}
	}
	return result, nil
}
