package main

import (
	"testing"
)

func TestGenStoreFn(t *testing.T) {
	fn := genStoreFn("AlbertRobertCarl")
	ex := "albert_robert_carl_store.go"
	if fn != ex {
		t.Errorf("genStoreFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
