package main

import (
	"github.com/codegangsta/cli"
	"os"
	"time"
)

var app *cli.App
var now time.Time

// VERSION stores the version string of this library
const VERSION = "0.1dev"

// TIMEFORMAT stores the time format to be used in
// timestamp generation
const TIMEFORMAT = "2006/01/02 15:04:05 (-0700)"

func init() {

	// core app
	app = cli.NewApp()
	app.Name = "gourd"
	app.Usage = "CLI tool to generates code for your API server"

	// define now
	now = time.Now()
}

func main() {
	app.Flags = []cli.Flag{}
	app.Version = VERSION

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
