package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/gourd/goparser"
	"github.com/gourd/gourd/compile"
)

func init() {
	// declare sub-command "gourd gen store"
	appendCmds("gen", cli.Command{
		Name:    "store",
		Aliases: []string{"s"},
		Usage:   "generate store from a type",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "type, t",
				Value: "",
				Usage: "type name of object in the store; must be set",
			},
			cli.StringFlag{
				Name:  "storage, s",
				Value: "upperio",
				Usage: "the storage engine behind the store",
			},
			cli.StringFlag{
				Name:  "coll, c",
				Value: "",
				Usage: "name of storage collection (i.e. MySQL table name)",
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
		Action: genStore,
	})
}

func decodeStore(c *cli.Context) (ctx compile.Context, err error) {

	ctx = compile.Context{
		//"Now": now.Format(TIMEFORMAT),
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

	// target type
	if c.String("type") == "" {
		fmt.Println("Please provide the target type(s)")
		os.Exit(1)
	}
	ctx.Set("TypeName", c.String("type"))

	// storage engine
	ctx.Set("Storage", c.String("storage"))

	// collection name
	if coll := c.String("coll"); coll == "" {
		r := regexp.MustCompile("[A-Z]+")
		cn := strings.ToLower(r.ReplaceAllString(ctx.GetStr("TypeName"), "_$0"))
		ctx.Set("Coll", cn)
	} else {
		ctx.Set("Coll", coll)
	}

	// parse type of given type name from given file(s)
	pkg, ts, err := readTypeFile(fns[0], []string{ctx.GetStr("TypeName")})
	if err != nil {
		err = compile.Errorf("Error parsing %#v. Error: %#v. Exit.", fns[0], err.Error())
		return
	}

	if len(ts) != 1 {
		err = compile.Errorf("Type %#v not found", ctx.GetStr("TypeName"))
		return
	}

	ctx.Set("Pkg", pkg)

	t := ts[0]

	// find the first id field of the type
	// (there shouldn't be more than 1 at all)
	var id *goparser.FieldSpec
	for field := range t.FieldsTagged("db", "id") {
		id = field
	}
	if id == nil {
		err = compile.Error("Failed to locate db id of the type %#v.\n"+
			"Cannot generate without id.",
			t.Name)
		return
	}

	ctx.Set("Id", id)
	ctx.Set("Type", t)
	return
}

// encode results to io writer (e.g. file)
func encodeStore(w io.Writer, ctx compile.Context) error {
	// write the generated output to file, according to storage engine
	return tpls.New("gen store:"+ctx.GetStr("Storage")).Execute(w, ctx)
}

// generate the store go file
func genStore(c *cli.Context) (err error) {

	// output file
	var out string
	if c.String("output") == "" {
		out = compile.SubfixFn("store")(c.String("type"))
	} else {
		out = c.String("output")
	}

	// compile the file
	com := compile.NewCompiler(decodeStore, encodeStore)
	if err = compile.CompileToFile(out, c, com); err != nil {
		return
	}

	fmt.Printf("generated %s\n", out)
	return
}
