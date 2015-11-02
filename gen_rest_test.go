package main

import (
	"testing"
)

func TestGenStoreRestFn(t *testing.T) {
	fn := genStoreRestFn("AlbertRobertCarl")
	ex := "albert_robert_carl_rest.go"
	if fn != ex {
		t.Errorf("genServiceRestFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
