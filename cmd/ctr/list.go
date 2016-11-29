package main

import "github.com/urfave/cli"

var listFlags = []cli.Flag{}

var listCommand = cli.Command{
	Name:  "list",
	Usage: "list running pods",
	Flags: listFlags,
	Action: func(context *cli.Context) error {
		return nil
	},
}
