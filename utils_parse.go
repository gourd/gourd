package main

import (
	"fmt"
	"github.com/gourd/goparser"
	"go/parser"
	"go/token"
)

// read TypeSpec of given type names from a file
func readTypeFile(inputPath string, tns []string) (pkg string, specs []*goparser.TypeSpec, err error) {
	fset := token.NewFileSet()

	// inputPath can only be filename
	f, err := parser.ParseFile(fset, inputPath, nil, parser.ParseComments)
	if err != nil {
		return
	}

	// read package name
	if f.Name == nil {
		err = fmt.Errorf("Unknown package name")
	} else {
		pkg = f.Name.Name
	}

	// read types name and details
	for pts := range goparser.DeclsToTypes(f.Decls) {
		if stringInSlice(pts.Name, tns) {
			specs = append(specs, pts)
		}
	}

	// see if all types needed are found
	if len(tns) != len(specs) {
		// TODO: improve this error message. Be specific on missing type.
		err = fmt.Errorf("Not all types can be found.")
	}

	return
}

// test if a string exists in a string slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
