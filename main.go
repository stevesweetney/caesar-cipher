package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// Lines is used to store the content of a file
type Lines []string

// check if a rune is a letter in the range (a...z) or (A...Z)
func isLetter(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}
	return true
}

func (l Lines) cipherAll(key int) {
	nLines := len(l)
	ch := make(chan int)

	for i := 0; i < nLines; i++ {
		go applyCipher(l, i, ch, key)
	}

	for i := 0; i < nLines; i++ {
		<-ch
	}

}

func applyCipher(lines Lines, i int, c chan int, key int) {
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
	c <- 1

}

// readFile opens a file at the specified path and returns
// a slice of strings that represents the lines of its content
func readFile(path string) (Lines, error) {
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

func main() {
	var fileName string

	app := cli.NewApp()
	app.Name = "Caesar-Cipher"
	app.Usage = "Encrypts/Decrypts a text file."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "encrypt, e, decrypt, d",
			Usage:       "`File` to encrypt/decrypt",
			Destination: &fileName,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}

		lines, err := readFile(fileName)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		fmt.Println(lines)
		return nil
	}

	app.Run(os.Args)
}
