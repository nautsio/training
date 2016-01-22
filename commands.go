package main

import (
	"github.com/codegangsta/cli"
)

// Commands is an array containing the available commands.
var Commands = []cli.Command{
	commandSetup,
	commandList,
}

var commandSetup = cli.Command{
	Name:        "setup",
	Usage:       "Setup training images on the specified cloud provider",
	Description: "Setup training images on the specified cloud provider.",
	Action:      doSetup,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "label, l",
			Usage: "The label of the created instances.",
		},
		cli.IntFlag{
			Name:  "count, c",
			Usage: "The number of instances to create.",
		},
		cli.StringFlag{
			Name:  "provider, p",
			Usage: "The cloud provider to use.",
		},
	},
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "List created instances",
	Description: "Show a list of ip addresses of the created instances.",
}

func doSetup(c *cli.Context) {
}
