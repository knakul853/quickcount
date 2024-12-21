package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestCount(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		wantLines int
		wantWords int
		wantBytes int
		wantChars int
	}{
		{
			name:      "empty file",
			input:     "",
			wantLines: 0,
			wantWords: 0,
			wantBytes: 0,
			wantChars: 0,
		},
		{
			name:      "single line",
			input:     "hello world",
			wantLines: 1,
			wantWords: 2,
			wantBytes: 11,
			wantChars: 11,
		},
		{
			name:      "multiple lines",
			input:     "hello\nworld\n",
			wantLines: 2,
			wantWords: 2,
			wantBytes: 12,
			wantChars: 12,
		},
		{
			name:      "multiple_lines_with_leading_and_trailing_spaces",
			input:     "  hello  \n world  \n  ",
			wantLines: 3,
			wantWords: 2,
			wantBytes: 21,
			wantChars: 21,
		},
		{
			name:      "non-ascii_characters",
			input:     "你好，世界\nこんにちは、世界",
			wantLines: 2,
			wantWords: 2,
			wantBytes: 40,
			wantChars: 14,
		},
		{
			name:      "mixed_characters",
			input:     "hello\n你好",
			wantLines: 2,
			wantWords: 2,
			wantBytes: 12,
			wantChars: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := bytes.NewReader([]byte(tc.input))
			lines, words, bytes, chars := count(reader)
			if lines != tc.wantLines {
				t.Errorf("got lines %d, want %d", lines, tc.wantLines)
			}
			if words != tc.wantWords {
				t.Errorf("got words %d, want %d", words, tc.wantWords)
			}
			if bytes != tc.wantBytes {
				t.Errorf("got bytes %d, want %d", bytes, tc.wantBytes)
			}
			if chars != tc.wantChars {
				t.Errorf("got chars %d, want %d", chars, tc.wantChars)
			}
		})
	}
}

func TestCount_Stdin(t *testing.T) {
	input := "test input\nsecond line"
	fmt.Printf("Input bytes: %v\n", []byte(input))

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	w.Write([]byte(input))
	w.Close()

	lines, words, bytes, chars := count(os.Stdin)

	// Expected values:
	// "test input\nsecond line" = 22 bytes
	// t e s t   i n p u t \n s e c o n d   l i n e
	// Total: 22 characters/bytes (10 chars + newline + 11 chars)

	if lines != 2 {
		t.Errorf("got lines %d, want %d", lines, 2)
	}
	if words != 4 {
		t.Errorf("got words %d, want %d", words, 4)
	}
	if bytes != 22 {
		t.Errorf("got bytes %d, want %d", bytes, 22)
	}
	if chars != 22 {
		t.Errorf("got chars %d, want %d", chars, 22)
	}
}
