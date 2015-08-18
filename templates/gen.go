// This file only provides command to generate
// the templates subpackage from "data" directory

//go:generate go-bindata -o templates.go -ignore=gen.go -ignore=templates.go -pkg=templates -prefix=data/ data/...

package templates
