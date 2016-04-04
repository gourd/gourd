package main

import (
	"io"

	"github.com/gourd/gourd/compile"
	"github.com/gourd/gourd/templates"

	"fmt"
	"os"

	"github.com/codegangsta/cli"
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
				Name:  "storekey, k",
				Value: "",
				Usage: "the entity store key variable",
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
	tpls.Append("gen endpoints", string(t))
	tpls.AddDeps("gen endpoints", "gen:general")
}

func decodeEndpoints(c *cli.Context) (ctx compile.Context, err error) {

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
	ctx.Set("StoreName", c.String("store"))

	if c.String("storekey") != "" {
		ctx.Set("StoreKey", c.String("storekey"))
	} else {
		ctx.Set("StoreKey", `"`+c.String("type")+"Store"+`"`)
	}

	// parse type of type name from given file(s)
	pkg, sts, err := readTypeFile(fns[0], []string{ctx.GetStr("StoreName")})
	if err != nil {
		err = compile.Errorf("Error parsing %#v. Error: %#v. Exit.", fns[0], err.Error())
		return
	}

	if len(sts) != 1 {
		err = compile.Errorf("Type %#v not found", ctx.GetStr("TypeName"))
		return
	}

	ctx.Set("Pkg", pkg)
	ctx.Set("Store", sts[0])
	return

}

// encode results to io writer (e.g. file)
func encodeEndpoints(w io.Writer, ctx compile.Context) error {
	// write the generated output to file, according to storage engine
	return tpls.New("gen endpoints").Execute(w, ctx)
}

// generate the endpoints go file
func genEndpoints(c *cli.Context) {

	// output file
	var out string
	if c.String("output") == "" {
		out = compile.SubfixFn("endpoints")(c.String("type"))
	} else {
		out = c.String("output")
	}

	// compile the file
	com := compile.NewCompiler(decodeEndpoints, encodeEndpoints)
	if err := compile.CompileToFile(out, c, com); err != nil {
		fmt.Println(err.Error())
		if gerr, ok := err.(compile.GourdError); ok {
			os.Exit(gerr.Code())
		}
		os.Exit(1)
	}

}
