package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "training"
	app.Version = Version
	app.Usage = "Create training images."
	app.Author = "nauts.io"
	app.Commands = Commands

	app.Run(os.Args)
}
