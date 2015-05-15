package main

import (
	"testing"
)

func TestGenServiceRestFn(t *testing.T) {
	fn := genServiceRestFn("AlbertRobertCarl")
	ex := "albert_robert_carl_rest.go"
	if fn != ex {
		t.Errorf("genServiceRestFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
