package main

import (
	"testing"
)

func TestGenServiceFn(t *testing.T) {
	fn := genServiceFn("AlbertRobertCarl")
	ex := "albert_robert_carl_store.go"
	if fn != ex {
		t.Errorf("genStoreFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
