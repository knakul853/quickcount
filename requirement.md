# Project Name: quickcount

## Functional Requirements

### Command-Line Interface (CLI):

- Implement the `ccwc` CLI tool.
- Accept the following options:
  - `-c`: Outputs the number of bytes in a file.
  - `-l`: Outputs the number of lines in a file.
  - `-w`: Outputs the number of words in a file.
  - `-m`: Outputs the number of characters in a file (supports multibyte characters if the locale allows).
  - No option: Outputs all counts (lines, words, and bytes).
- Support reading input from a file or standard input (when no filename is provided).

### Output Format:

Match the format of Unix `wc`:

```
<count> <filename>
```

For no options:

```
<lines> <words> <bytes> <filename>
```

For standard input:

```
<count>
```

## Technical Requirements

### Programming Language:

- Use Golang as the implementation language.

### File Handling:

- Open and read files using efficient file I/O (`os` and `bufio` packages).

### Input Validation:

- Validate input arguments and provide meaningful error messages for invalid input.
- Ensure the file exists and is readable.
- Handle cases when no arguments or invalid options are provided.

### Counting Logic:

- Bytes (`-c`): Use the `len()` function on the file content.
- Lines (`-l`): Count the number of newline (`\n`) characters in the file.
- Words (`-w`): Split text by whitespace (e.g., using `strings.Fields()`).
- Characters (`-m`): Count characters considering multibyte characters (use `utf8.RuneCountInString` from the `unicode/utf8` package).

### Locale Support (`-m`):

- Ensure the tool respects the system's locale settings for multibyte characters.
- Use `golang.org/x/text` for enhanced locale handling if needed.

### Standard Input:

- When no filename is provided, read from `os.Stdin`.
- Ensure compatibility with piped input, such as:

```bash
cat test.txt | ccwc -l
```

### Testing:

- Validate functionality against the expected outputs of the Unix `wc` tool.
- Test edge cases:
  - Empty file.
  - Large files.
  - Non-ASCII characters (e.g., UTF-8 encoded text).

### Performance:

- Process files line by line to handle large files efficiently.
- Avoid loading the entire file into memory unless absolutely necessary.

### Documentation:

- Provide a `README.md` file with:
  - Project description.
  - Usage instructions.
  - Examples of input and output.
- Include inline comments for clarity in code.

## Optional Enhancements

### Multi-File Support:

- Support processing multiple files in a single command:

```bash
ccwc -l file1.txt file2.txt
```

- Display totals when multiple files are processed.

### Error Handling:

- Gracefully handle scenarios like missing files, permission errors, or invalid input.

### Flags for Help and Version:

- Add `--help` to display usage instructions.
- Add `--version` to display the tool's version.
