# Go Reloaded

Go Reloaded is a text processing program written in Go.
It reads a text file, applies different formatting and correction rules, and writes the processed result to a new file.

## Features

The program supports:

* Converting hexadecimal numbers to decimal.
* Converting binary numbers to decimal.
* Converting text to uppercase.
* Converting text to lowercase.
* Capitalizing words.
* Applying transformations to a specific number of words.
* Correcting punctuation spacing.
* Replacing `a` with `an` when required.
* Removing unnecessary spaces around punctuation.

## Usage

Run the program with:

```bash
go run . input.txt output.txt
```

Where:

* `input.txt` is the input file containing the text to process.
* `output.txt` is the file where the processed text will be saved.

### Example

Input:

```text
I have 2 (hex) apples and 10 (bin) oranges.
this is a test. (up)
```

Output:

```text
I have 2 apples and 2 oranges.
THIS IS A TEST.
```

## Supported Commands

| Command    | Description                                         |
| ---------- | --------------------------------------------------- |
| `(hex)`    | Converts the previous hexadecimal number to decimal |
| `(bin)`    | Converts the previous binary number to decimal      |
| `(up)`     | Converts the previous word to uppercase             |
| `(low)`    | Converts the previous word to lowercase             |
| `(cap)`    | Capitalizes the previous word                       |
| `(up, n)`  | Converts the previous `n` words to uppercase        |
| `(low, n)` | Converts the previous `n` words to lowercase        |
| `(cap, n)` | Capitalizes the previous `n` words                  |

## Requirements

* Go 1.XX or later

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Navigate to the project directory:

```bash
cd go-reloaded
```

Run the program:

```bash
go run . input.txt output.txt
```

## Project Structure

```text
go-reloaded/
├── main.go
├── input.txt
├── output.txt
├── go.mod
└── README.md
```

## Author

**Shady Elmougy**
