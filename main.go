package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

func main() {
	if _, err := exec.LookPath("terraform"); err != nil {
		log.Fatal("[ERROR] Terraform not found")
	}

	app := cli.NewApp()
	app.Name = "training"
	app.Version = Version
	app.Usage = "Create training environments."
	app.Author = "nauts.io"
	app.Commands = Commands

	app.Run(os.Args)
}
