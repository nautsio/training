package main

import (
	"github.com/codegangsta/cli"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	_, err := exec.LookPath("terraform")
	if err != nil {
		fmt.Println("FATAL - Terraform not found")
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = "training"
	app.Version = Version
	app.Usage = "Create training images."
	app.Author = "nauts.io"
	app.Commands = Commands

	app.Run(os.Args)
}
