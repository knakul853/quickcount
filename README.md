# quickcount

## Project Description

`quickcount` is a command-line tool that mimics the functionality of the Unix `wc` command. It counts the number of lines, words, bytes, and characters in a file or from standard input.

## Usage Instructions

The `ccwc` command-line tool accepts the following options:

-   `-c`: Outputs the number of bytes in a file.
-   `-l`: Outputs the number of lines in a file.
-   `-w`: Outputs the number of words in a file.
-   `-m`: Outputs the number of characters in a file (supports multibyte characters if the locale allows).
-   No option: Outputs all counts (lines, words, and bytes).

The tool can read input from a file or from standard input (when no filename is provided).

## Examples

### Counting lines in a file:

```bash
ccwc -l test.txt
```

### Counting words in a file:

```bash
ccwc -w test.txt
```

### Counting bytes in a file:

```bash
ccwc -c test.txt
```

### Counting characters in a file:

```bash
ccwc -m test.txt
```

### Counting lines, words, and bytes in a file:

```bash
ccwc test.txt
```

### Reading from standard input:

```bash
cat test.txt | ccwc -l
```

```bash
echo "hello world" | ccwc
```

## Building and Running

To build the project, run:

```bash
go build main.go
```

To run the tool, use:

```bash
./main [options] [filename]
