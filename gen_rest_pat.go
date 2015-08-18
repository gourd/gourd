package main

import (
	"github.com/gourd/gourd/templates"
)

func init() {
	t, err := templates.Asset("rest/gorilla/pat.tpl")
	if err != nil {
		panic(err)
	}
	tpls.Append("gen rest:gorilla/pat", string(t))
	tpls.AddDeps("gen rest:gorilla/pat", "gen:general")
}
