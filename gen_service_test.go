package main

import (
	"testing"
)

func TestGenServiceFn(t *testing.T) {
	fn := genServiceFn("AlbertRobertCarl")
	ex := "albert_robert_carl_service.go"
	if fn != ex {
		t.Errorf("genServiceFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
