package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"regexp"
	"strings"
)

func init() {
	// declare sub-command "gourd gen rest"
	appendCmds("gen", cli.Command{
		Name:    "rest",
		Aliases: []string{"s"},
		Usage:   "generate rest from a store type",
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
				Name:  "router, r",
				Value: "gorilla/pat",
				Usage: "name of router to use",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "output file name; default srcdir/<type>_store.go",
			},
		},
		Action: genStoreRest,
	})
}

func genStoreRestFn(tn string) string {
	r1 := regexp.MustCompile("[A-Z]+")
	r2 := regexp.MustCompile("^\\_")
	return strings.ToLower(r2.ReplaceAllString(r1.ReplaceAllString(tn, "_$0"), "")) + "_rest.go"
}

// generate the store rest go file
func genStoreRest(c *cli.Context) {

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

	// router
	var s string
	s = c.String("router")

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
			o = genStoreRestFn(st.Name)
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
		err = tpls.New("gen rest:"+s).Execute(f, map[string]interface{}{
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
