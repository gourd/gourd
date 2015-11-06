package compile_test

import (
	"testing"

	"github.com/gourd/gourd/compile"
)

func TestSubfixFn(t *testing.T) {
	have := compile.SubfixFn("rest")("AlbertRobertCarl")
	want := "albert_robert_carl_rest.go"
	if want != have {
		t.Errorf("compile.SubfixFn failed. Expect \"%s\" but got \"%s\"",
			want, have)
	}
}
