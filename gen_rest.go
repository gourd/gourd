package main

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
	"github.com/gourd/gourd/compile"
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
				Name:  "storekey, k",
				Value: "",
				Usage: "the entity store key variable",
			},
			cli.StringFlag{
				Name:  "router, r",
				Value: "gorilla/pat",
				Usage: "name of router to use",
			},
			cli.BoolFlag{
				Name:  "preserve, p",
				Usage: "if this flag is set, will not overwrite existing file",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: "",
				Usage: "output file name; default srcdir/<type>_store.go",
			},
		},
		Action: genRest,
	})
}

func decodeRest(c *cli.Context) (ctx compile.Context, err error) {

	ctx = compile.Context{
		"Now": now.Format(TIMEFORMAT),
		"Ver": VERSION,
	}

	// files to parse
	var fns []string
	if len(c.Args()) == 0 {
		// TODO: find all .go in the current folder
		err = compile.Error("Please provide files to parse")
		return
	}
	fns = c.Args()

	// target entity type
	if c.String("type") == "" {
		err = compile.Error("Please provide the target entity type")
		return
	}
	ctx.Set("Type", c.String("type"))

	// target store type
	if c.String("store") == "" {
		err = compile.Error("Please provide the target store type")
		return
	}
	ctx.Set("Store", c.String("store"))

	if c.String("storekey") != "" {
		ctx.Set("StoreKey", c.String("storekey"))
	} else {
		ctx.Set("StoreKey", `"`+c.String("store")+`"`)
	}

	// router
	ctx.Set("Router", c.String("router"))

	// read type of type name from given file(s)
	pkg, sts, ierr := readTypeFile(fns[0], []string{ctx.GetStr("Store")})
	if ierr != nil {
		err = compile.Error("Error parsing \"%s\". Error: %s. Exit.", fns[0], ierr.Error())
		return
	}
	ctx.Set("Pkg", pkg)

	if len(sts) < 1 {
		err = compile.Exit("type %#v not found\n", ctx.GetStr("Store"))
	}

	return
}

// encode results to io writer (e.g. file)
func encodeRest(w io.Writer, ctx compile.Context) error {
	// write the generated output to file
	return tpls.New("gen rest:"+ctx.GetStr("Router")).Execute(w, ctx)
}

// generate the store rest go file
func genRest(c *cli.Context) {

	// output file name
	var out string
	if c.String("output") == "" {
		out = compile.SubfixFn("rest")(c.String("type"))
	} else {
		out = c.String("output")
	}

	// compile the file
	com := compile.NewCompiler(decodeRest, encodeRest)
	if err := compile.CompileToFile(out, c, com); err != nil {
		fmt.Println(err.Error())
		if gerr, ok := err.(compile.GourdError); ok {
			os.Exit(gerr.Code())
		}
		os.Exit(1)
	}

	fmt.Printf("generated %s\n", out)

}
