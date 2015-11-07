package compile

import (
	"io"
	"os"
	"os/exec"

	"github.com/codegangsta/cli"
)

// Compiler takes a CLI context, decode it and compile to the io writer
type Compiler interface {
	Compile(w io.Writer, c *cli.Context) error
}

// compiler implements Compiler interface
type compiler struct {
	decode DecoderFunc
	encode EncoderFunc
}

// Compile implements the Compiler interface
func (com compiler) Compile(w io.Writer, c *cli.Context) (err error) {

	// decode the compile context from arguments
	ctx, err := com.decode(c)
	if err != nil {
		return
	}

	// generate the file content with template and
	// write to the file stream
	return com.encode(w, ctx)
}

// NewCompiler takes decoder and encoder and use them in
// file compilation
func NewCompiler(dec DecoderFunc, enc EncoderFunc) Compiler {
	return &compiler{
		decode: dec,
		encode: enc,
	}
}

// FormatFile formats a generated file with go formatter
func FormatFile(fn string) error {
	return exec.Command("go", "fmt", fn).Run()
}

// CompileToFile compiles to file according to given CLI flags
// output: output file name
func CompileToFile(fn string, c *cli.Context, com Compiler) (err error) {

	// check if the file exists before generating, except with "force" flag
	if _, ierr := os.Stat(fn); !os.IsNotExist(ierr) {
		if c.Bool("preserve") == true {
			err = Exit("File %#v exists. Preserve the file", fn)
			return
		}
	}

	// create output file (if not exists)
	f, err := os.Create(fn)
	defer FormatFile(fn)
	defer f.Close()
	if err != nil {
		err = Error(
			"Failed to create output file %#v. reason: %#v",
			fn, err.Error())
		return
	}

	// compile the file
	err = com.Compile(f, c)
	if err != nil {
		// clean up the file if has error
		defer os.Remove(fn)
	}
	return
}
