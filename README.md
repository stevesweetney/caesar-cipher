# Caesar Cipher

Caesar cipher implementation in go, encodes and decodes text files.

### Command-line Usage

First, [install Go](https://golang.org/doc/install).

    go get -u github.com/stevesweetney/caesar-cipher
    caesar-cipher -i input.txt -o output.txt -k 15


| Flag | Default | Description |
| --- | --- | --- |
| `i` | n/a | input file |
| `o` | n/a | output file |
| `k` | 0 | number of letters to shift |
| `d` | false | if set, shifts letters to the left by k. |
