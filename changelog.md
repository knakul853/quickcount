# Changelog

## Initial Commit

- Created `main.go` with basic structure for command-line argument parsing and file input.
- Implemented basic counting logic for lines, words, bytes, and characters.
- Added support for reading from standard input when no filename is provided.
- Created `changelog.md` to track changes.

## Added Logging and Comments

- Added logging statements to `main.go` to track the program's execution.
- Added comments to `main.go` to explain the code's functionality.

## Fixed Counting Logic

- Corrected the `count` function to accurately count lines, words, bytes, and characters.
- Fixed issues with byte and character counts due to newline characters.
- Fixed issues with line counts due to empty lines.
- Fixed issues with standard input handling in tests.
- Addressed issues with incorrect counts in various test cases.
