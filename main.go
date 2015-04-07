package main

import (
	"github.com/codegangsta/cli"
	"os"
)

var app *cli.App

func init() {

	// core app
	app = cli.NewApp()
	app.Name = "gourd"
	app.Usage = "CLI tool to generates code for your API server"

}

func main() {
	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "gen",
			Usage: "generate code",
			Action: func(c *cli.Context) {
				println("generate: ", c.Args().First())
			},
			Subcommands: cmds["gen"],
		},
	}

	app.Run(os.Args)
}
