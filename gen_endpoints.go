package main

import (
	"github.com/gourd/gourd/templates"

	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"regexp"
	"strings"
)

func init() {
	// declare sub-command "gourd gen store"
	appendCmds("gen", cli.Command{
		Name:    "endpoints",
		Aliases: []string{"s"},
		Usage:   "generate endpoints from a store type",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Value: "",
				Usage: "type name of the entity, required",
			},
			cli.StringFlag{
				Name:  "store, s",
				Value: "",
				Usage: "type name of the entity store, required",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "output file name; default srcdir/<type>_store.go",
			},
		},
		Action: genEndpoints,
	})

	t, err := templates.Asset("endpoints/endpoints.tpl")
	if err != nil {
		panic(err)
	}
	tpls.Append("gen endpoints:endpoints", string(t))
	tpls.AddDeps("gen endpoints:endpoints", "gen:general")
}

func genEndpointsFn(tn string) string {
	r1 := regexp.MustCompile("[A-Z]+")
	r2 := regexp.MustCompile("^\\_")
	return strings.ToLower(r2.ReplaceAllString(r1.ReplaceAllString(tn, "_$0"), "")) + "_endpoints.go"
}

// generate the endpoints go file
func genEndpoints(c *cli.Context) {

	// files to parse
	var fns []string
	if len(c.Args()) == 0 {
		// TODO: find all .go in the current folder
		fmt.Println("Please provide files to parse")
		os.Exit(1)
	} else {
		fns = c.Args()
	}

	// target entity type
	if c.String("type") == "" {
		fmt.Println("Please provide the target entity type")
		os.Exit(1)
	}
	tn := c.String("type")

	// target store type
	if c.String("store") == "" {
		fmt.Println("Please provide the target store type")
		os.Exit(1)
	}
	sn := c.String("store")

	// read type of type name from given file(s)
	pkg, sts, err := readTypeFile(fns[0], []string{sn})
	if err != nil {
		fmt.Printf("Error parsing \"%s\". Error: %s. Exit.", fns[0], err.Error())
		os.Exit(1)
	}

	// loop through each type found
	for _, st := range sts {

		// output file
		var o string
		if c.String("output") == "" {
			o = genEndpointsFn(tn)
		} else {
			o = c.String("output")
		}

		// create output file (if not exists)
		f, err := os.Create(o)
		defer FormatFile(o)
		defer f.Close()
		if err != nil {
			fmt.Printf("Failed to create output file \"%s\".\n", o)
			fmt.Printf("Error: \"%s\"\nExit.\n", err.Error())
			os.Exit(1)
		}

		// write the generated output to file
		err = tpls.New("gen endpoints:endpoints").Execute(f, map[string]interface{}{
			"Now":   now.Format(TIMEFORMAT),
			"Ver":   VERSION,
			"Pkg":   pkg,
			"Type":  tn,
			"Store": st,
		})
		if err != nil {
			fmt.Printf("Failed to write to file \"%s\".\n", o)
			fmt.Printf("Error: \"%s\"\nExit.\n", err.Error())
			os.Exit(1)
		}
	}

}
