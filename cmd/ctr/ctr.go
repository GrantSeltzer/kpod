package main

import "github.com/urfave/cli"

var ctrSubCommands = []cli.Command{
	listCommand,
}

var ctrFlags = []cli.Flag{}

var ctrCommand = cli.Command{
	Name:        "ctr",
	Usage:       "containers",
	Flags:       ctrFlags,
	Subcommands: ctrSubCommands,
}
