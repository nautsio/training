package main

import (
	"github.com/codegangsta/cli"
	"github.com/nautsio/training/command"
)

// Commands is an array containing the available commands.
var Commands = []cli.Command{
	commandSetup,
	commandList,
}

var commandSetup = cli.Command{
	Name:        "setup",
	Usage:       "Setup training environments on the specified cloud provider",
	Description: "Setup training environments on the specified cloud provider.",
	Action:      command.Setup,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "label",
			Usage: "The label of the created environments - required",
		},
		cli.IntFlag{
			Name:  "count",
			Usage: "The number of environments to create",
			Value: 1,
		},
		cli.StringFlag{
			Name:  "provider",
			Usage: "The cloud provider to use - required",
		},
		cli.StringFlag{
			Name:   "aws-access-key",
			Usage:  "TODO",
			EnvVar: "AWS_ACCESS_KEY_ID",
		},
		cli.StringFlag{
			Name:   "aws-secret-key",
			Usage:  "TODO",
			EnvVar: "AWS_SECRET_ACCESS_KEY",
		},
	},
}

var commandList = cli.Command{
	Name:        "list",
	Usage:       "List created environments",
	Description: "Show a list of ip addresses of the created environments.",
}
