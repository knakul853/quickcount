package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// Log the start of the program
	log.Println("Starting ccwc tool")

	// Define command-line flags
	bytesFlag := flag.Bool("c", false, "Output the number of bytes in a file")
	linesFlag := flag.Bool("l", false, "Output the number of lines in a file")
	wordsFlag := flag.Bool("w", false, "Output the number of words in a file")
	charsFlag := flag.Bool("m", false, "Output the number of characters in a file")

	// Log the parsing of command-line flags
	log.Println("Parsing command-line flags")
	flag.Parse()

	// Get the filename from the command-line arguments
	var filename string
	if len(flag.Args()) > 0 {
		filename = flag.Arg(0)
		// Log the filename if provided
		log.Printf("Filename provided: %s", filename)
	} else {
		// Log that no filename was provided and input will be read from stdin
		log.Println("No filename provided, reading from stdin")
	}

	// Read from stdin if no filename is provided
	var reader io.Reader
	if filename == "" {
		reader = os.Stdin
	} else {
		// Log the opening of the file
		log.Printf("Opening file: %s", filename)
		file, err := os.Open(filename)
		if err != nil {
			// Log the error if the file cannot be opened
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			log.Printf("Error opening file: %v", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	// Log the start of the counting process
	log.Println("Performing counting")
	// Perform counting
	lineCount, wordCount, byteCount, charCount := count(reader)

	// Log the output of the results
	log.Println("Outputting results")
	// Output the results
	if *bytesFlag {
		fmt.Printf("%d ", byteCount)
	} else if *linesFlag {
		fmt.Printf("%d ", lineCount)
	} else if *wordsFlag {
		fmt.Printf("%d ", wordCount)
	} else if *charsFlag {
		fmt.Printf("%d ", charCount)
	} else {
		fmt.Printf("%d %d %d ", lineCount, wordCount, byteCount)
	}

	if filename != "" {
		fmt.Printf("%s\n", filename)
	} else {
		fmt.Println()
	}
	// Log the end of the program
	log.Println("ccwc tool finished")
}

// count reads from the provided io.Reader and returns the line, word, byte, and character counts.
func count(reader io.Reader) (int, int, int, int) {
	log.Println("Starting count function")
	lineCount := 0
	wordCount := 0
	byteCount := 0
	charCount := 0

	// Read the entire input
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("Error reading file: %v", err)
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Count bytes directly from the content
	byteCount = len(content)

	// Count characters (runes)
	charCount = utf8.RuneCount(content)

	// Split content into lines
	text := string(content)
	if text == "" {
		log.Println("Count function finished")
		return 0, 0, 0, 0
	}

	// Count lines and words
	lines := strings.Split(text, "\n")
	lineCount = len(lines)
	if len(content) > 0 && content[len(content)-1] == '\n' {
		lineCount--
	}

	// Count words
	for _, line := range lines {
		fields := strings.Fields(line)
		wordCount += len(fields)
	}

	log.Println("Count function finished")
	return lineCount, wordCount, byteCount, charCount
}
