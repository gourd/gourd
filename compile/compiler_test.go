package compile_test

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/codegangsta/cli"
	"github.com/gourd/gourd/compile"
)

func TestCompiler(t *testing.T) {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	// dummy codec
	dec := func(c *cli.Context) (ctx compile.Context, err error) {
		ctx = compile.Context{
			"string": "World",
		}
		return
	}
	enc := func(w io.Writer, ctx compile.Context) error {
		fmt.Fprintf(w, "Hello %s", ctx.GetStr("string"))
		return nil
	}

	com := compile.NewCompiler(dec, enc)
	if err := com.Compile(w, nil); err != nil {
		t.Errorf("compile error: %#v", err.Error())
	}

	// flush buffer before read
	w.Flush()

	if want, have := "Hello World", b.String(); want != have {
		t.Logf("len(b) = %#v", b.Len())
		t.Errorf("want: %#v, got: %#v", want, have)
	}
}

func TestCompilerToFile_NoPreserve(t *testing.T) {

	fn := "test.tmp"

	// create *cli.Context with flag "preserve" set
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.Bool("preserve", false, "prevent overwriting")
	c := cli.NewContext(cli.NewApp(), flags, nil)

	// dummy codec
	dec := func(c *cli.Context) (ctx compile.Context, err error) {
		return
	}
	enc := func(w io.Writer, ctx compile.Context) error {
		w.Write([]byte("Hello world"))
		return nil
	}

	// generate a file
	com := compile.NewCompiler(dec, enc)
	if err := compile.CompileToFile(fn, c, com); err != nil {
		t.Errorf("error compiling: %#v", err.Error())
	}
	defer os.Remove(fn)

	// generate the file again
	if err := compile.CompileToFile(fn, c, com); err != nil {
		t.Errorf("expect nil but got %#v", err)
	}

}

func TestCompileToFile_Preserve(t *testing.T) {

	fn := "test.tmp"

	// create *cli.Context with flag "preserve" set
	flags := flag.NewFlagSet("", flag.ContinueOnError)
	flags.Bool("preserve", true, "prevent overwriting")
	c := cli.NewContext(cli.NewApp(), flags, nil)

	// dummy codec
	dec := func(c *cli.Context) (ctx compile.Context, err error) {
		return
	}
	enc := func(w io.Writer, ctx compile.Context) error {
		w.Write([]byte("Hello world"))
		return nil
	}

	// generate a file
	com := compile.NewCompiler(dec, enc)
	if err := compile.CompileToFile(fn, c, com); err != nil {
		t.Errorf("error compiling: %#v", err.Error())
	}
	defer os.Remove(fn)

	// generate the file again
	if err := compile.CompileToFile(fn, c, com); err == nil {
		t.Error("expect error but got nil")
	} else if cerr, ok := err.(compile.GourdError); !ok {
		t.Errorf("expect GourdError got %#v", err)
	} else if want, have := 0, cerr.Code(); want != have {
		t.Errorf("expect code %#v, got %#v", want, have)
	}

}
