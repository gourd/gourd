package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"go/parser"
	"go/token"
	"os"
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
				Name:  "output, o",
				Value: "",
				Usage: "output file name; default srcdir/<type>_service.go",
			},
		},
		Action: genService,
	})
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

	// output file
	var o string
	if c.String("output") == "" {
		o = strings.ToLower(tn) + "_service.go"
	} else {
		o = c.String("output")
	}

	// read type of type name from given file(s)
	t, ok := readTypeFile(fns[0], tn)
	if !ok {
		fmt.Printf("Type \"%s\" not found in the provided file(s). Exit.\n", tn)
		os.Exit(1)
	}

	// dummy output
	fmt.Printf("t: %#v ==> output: %s\n", t, o)
}

// read typeSpec from files
func readTypeFile(inputPath string, tn string) (spec genTypeSpec, ok bool) {
	fset := token.NewFileSet()

	// TODO: get to know what inputPath can be (e.g. dir?)
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// read package name
	var pkg string
	if f.Name == nil {
		fmt.Println("Error. Unknown package name. Exit.")
		os.Exit(1)
	} else {
		pkg = f.Name.Name
	}

	// read types name and details
	for ts := range filterTypeSpec(filterGenDecl(chanDecls(f.Decls))) {
		pts := parseTypeSpec(ts)
		if pts.Name == tn {
			spec = pts
			ok = true
			break
		}
	}
	spec.Pkg = pkg
	return
}
