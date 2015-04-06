package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func init() {
	appendCmds("gen", cli.Command{
		Name:    "service",
		Aliases: []string{"s"},
		Usage:   "generate service from a type",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Value: "",
				Usage: "comma-separated list of type names; must be set",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "output file name; default srcdir/<type>_service.go",
			},
		},
		Action: genService,
	})
}

func genService(c *cli.Context) {

	// files to parse
	fns := make([]string, 0)
	if len(c.Args()) == 0 {
		// TODO: find all .go in the current folder
		fmt.Println("Please provide files to parse")
		os.Exit(1)
	} else {
		fns = c.Args()
	}

	// target type
	if c.String("type") == "" {
		fmt.Println("Please provide the target type(s)")
		os.Exit(1)
	}
	t := c.String("type")

	// output file
	var o string
	if c.String("output") == "" {
		o = strings.ToLower(t) + "_service.go"
	} else {
		o = c.String("output")
	}

	// dummy output
	fmt.Printf("t: %s, o: %s, fns: %s\n", t, o, fns)
}
