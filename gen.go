package main

import (
	"github.com/gourd/gourd/templates"
)

func init() {
	t, err := templates.Asset("general.tpl")
	if err != nil {
		panic(err)
	}
	tpls.Append("gen:general", string(t))
}
