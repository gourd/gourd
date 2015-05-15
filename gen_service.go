package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"regexp"
	"strings"
)

func init() {
	// declare sub-command "gourd gen service"
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
				Name:  "storage, s",
				Value: "upperio",
				Usage: "the storage type behind the service",
			},
			cli.StringFlag{
				Name:  "coll, c",
				Value: "",
				Usage: "name of storage collection (i.e. MySQL table name)",
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

func genServiceFn(tn string) string {
	r1 := regexp.MustCompile("[A-Z]+")
	r2 := regexp.MustCompile("^\\_")
	return strings.ToLower(r2.ReplaceAllString(r1.ReplaceAllString(tn, "_$0"), "")) + "_service.go"
}

// generate the service go file
func genService(c *cli.Context) {

	// files to parse
	var fns []string
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
	tn := c.String("type")

	// separated the type names with comma
	tns := strings.Split(tn, ",")

	// storage
	var s string
	s = c.String("storage")

	// collection name
	var cn string
	cn = c.String("coll")
	if cn == "" {
		r := regexp.MustCompile("[A-Z]+")
		cn = strings.ToLower(r.ReplaceAllString(tn, "_$0"))
	}

	// read type of type name from given file(s)
	pkg, ts, err := readTypeFile(fns[0], tns)
	if err != nil {
		fmt.Printf("Error parsing \"%s\". Error: %s. Exit.", fns[0], err.Error())
		os.Exit(1)
	}

	// loop through each type found
	for _, t := range ts {

		// output file
		var o string
		if c.String("output") == "" {
			o = genServiceFn(t.Name)
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
		err = tpls.New("gen service:"+s).Execute(f, map[string]interface{}{
			"Pkg":  pkg,
			"Type": t,
			"Coll": cn,
		})
		if err != nil {
			fmt.Printf("Failed to write to file \"%s\".\n", o)
			fmt.Printf("Error: \"%s\"\nExit.\n", err.Error())
			os.Exit(1)
		}
	}

}
