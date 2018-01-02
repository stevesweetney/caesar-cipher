package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Caesar-Cipher"
	app.Usage = "Encrypts/Decrypts a text file."

	app.Run(os.Args)
}
