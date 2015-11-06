package compile

import (
	"io"

	"github.com/codegangsta/cli"
)

// DecoderFunc decode a given CLI context into a compile context
type DecoderFunc func(c *cli.Context) (ctx Context, err error)

// EncoderFunc compiles a given comiple context and output result
// to the io.Writer
type EncoderFunc func(w io.Writer, ctx Context) error
