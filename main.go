package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func readFile(path string) ([]string, error) {
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
