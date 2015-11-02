package main

import (
	"testing"
)

func TestGenEndpointsFn(t *testing.T) {
	fn := genEndpointsFn("AlbertRobertCarl")
	ex := "albert_robert_carl_endpoints.go"
	if fn != ex {
		t.Errorf("genEndpointsFn failed. Expect \"%s\" but got \"%s\"",
			ex, fn)
	}
}
