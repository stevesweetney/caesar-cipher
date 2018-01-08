package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/urfave/cli"
)

// Lines is used to store the content of a file
type Lines []string

func saveFile(path string, content Lines) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("%s already exists", path)
		return
	}

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error occurred opening file %s", err)
		return
	}
	defer file.Close()

	for _, l := range content {
		formatted := fmt.Sprintf("%s\r\n", l)
		file.WriteString(formatted)
	}
}

// check if a rune is a letter in the range (a...z) or (A...Z)
func isLetter(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}
	return true
}

// CipherAll shifts all the letters in each line by the key arguement
// shift right for numbers > 0, shift left for numbers < 0
func (l Lines) CipherAll(key int) {
	numCPU := runtime.NumCPU()
	nLines := len(l)
	ch := make(chan int, nLines)

	for i := 0; i < numCPU; i++ {
		go l.doSome(i*nLines/numCPU, (i+1)*nLines/numCPU, ch, key)
	}

	for i := 0; i < numCPU; i++ {
		<-ch
	}
}

func (l Lines) doSome(i, n int, ch chan int, key int) {
	for ; i < n; i++ {
		applyCipher(l, i, key)
	}
	ch <- 1
}

func applyCipher(lines Lines, i int, key int) {
	line := lines[i]
	translated := make([]rune, 0, len(line))

	for _, r := range line {
		if isLetter(r) {
			shifted := int(r) + key
			if r >= 'a' && r <= 'z' {
				if shifted > 'z' {
					shifted -= 26
				} else if shifted < 'a' {
					shifted += 26
				}
			} else {
				if shifted > 'Z' {
					shifted -= 26
				} else if shifted < 'A' {
					shifted += 26
				}
			}
			translated = append(translated, rune(shifted))
		} else {
			translated = append(translated, r)
		}
	}

	lines[i] = string(translated)
}

// ReadFile opens a file at the specified path and returns
// a slice of strings that represents the lines of its content
func ReadFile(path string) (Lines, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 1)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func errorMessage(message string) bool {
	fmt.Println(message)
	return false
}

func main() {
	var decrypt bool
	var input string
	var output string
	var key int

	app := cli.NewApp()
	app.Name = "Caesar-Cipher"
	app.Usage = "Encrypts/Decrypts a text file."
	app.UsageText = "Usage: caesar-cipher [OPTIONS] -i input -o output -k key"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "decrypt, d",
			Usage:       "Decrypts the input file if set. Defaults to encrypting",
			Destination: &decrypt,
		},
		cli.StringFlag{
			Name:        "input, i",
			Usage:       "`File` to encrypt/decrypt",
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "output, o",
			Usage:       "output file",
			Destination: &output,
		},
		cli.IntFlag{
			Name:        "key, k",
			Usage:       "Number between 1-25 to use when shifting",
			Destination: &key,
		},
	}

	app.Action = func(c *cli.Context) error {
		ok := true
		if input == "" {
			ok = errorMessage("ERROR: input argument required")
		}

		if output == "" {
			ok = errorMessage("ERROR: output argument required")
		}

		if key < 1 || key > 25 {
			ok = errorMessage("ERROR: enter a number between 1-25 to use as key")
		}

		if !ok {
			return cli.ShowAppHelp(c)
		}

		lines, err := ReadFile(input)
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		if decrypt {
			key = -key
		}
		lines.CipherAll(key)

		saveFile(output, lines)
		return nil
	}

	app.Run(os.Args)
}
