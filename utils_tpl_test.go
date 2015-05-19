package main

import (
	"bytes"
	"testing"
)

func TestTemplatePrep(t *testing.T) {

	// define prep
	var prep TemplatePrep = func(in interface{}) (interface{}, error) {
		if m, ok := in.(*map[string]string); ok {
			(*m)["output"] = "changed"
		}
		return in, nil
	}

	// prep mtest
	mtest := map[string]string{
		"output": "not changed",
	}
	prep(&mtest)

	// test output
	if mtest["output"] != "changed" {
		t.Errorf("Failed to prep a map variable. Expect \"%s\" but get \"%s\"",
			"changed", mtest["output"])
	}
}

func TestTemplates_AddPrep(t *testing.T) {
	tpls := &Templates{}
	tpls.Append("test", "hello {{ .Val }}")
	tpls.AddPrep("test", func(in interface{}) (interface{}, error) {
		if m, ok := in.(*map[string]string); ok {
			(*m)["Val"] = "world"
		}
		return in, nil
	})
	expect := "hello world"

	// test the output
	var b bytes.Buffer // A Buffer needs no initialization.
	tpls.New("test").Execute(&b, &map[string]string{
		"Val": "foo",
	})

	if b.String() != expect {
		t.Logf("Prep function was not run as expected in template execute")
		t.Errorf("Expecting \"%s\" but get \"%s\"", expect, b.String())
	}
}
