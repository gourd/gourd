package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// representation of a general type spec
type pTypeSpec struct {
	Name   string
	Type   string
	Fields []pFieldSpec
}

// representation of a general struct field
type pFieldSpec struct {
	Name string
	Type string
	Tag  string
}

// channel the slice of ast.Decl out for filtering
func chanDecls(decls []ast.Decl) <-chan ast.Decl {
	cout := make(chan ast.Decl)
	go func(decls []ast.Decl, cout chan ast.Decl) {
		defer close(cout)
		for _, decl := range decls {
			cout <- decl
		}
	}(decls, cout)
	return cout
}

// filter only GetDecl
func filterGenDecl(cin <-chan ast.Decl) <-chan *ast.GenDecl {
	cout := make(chan *ast.GenDecl)
	go func(cin <-chan ast.Decl, cout chan *ast.GenDecl) {
		defer close(cout)
		for decl := range cin {
			if genDecl, ok := decl.(*ast.GenDecl); ok {
				cout <- genDecl
			}
		}
	}(cin, cout)
	return cout
}

// filter only TypeSpec
func filterTypeSpec(cin <-chan *ast.GenDecl) <-chan *ast.TypeSpec {
	cout := make(chan *ast.TypeSpec)
	go func(cin <-chan *ast.GenDecl, cout chan *ast.TypeSpec) {
		defer close(cout)
		for decl := range cin {
			// Note: may further break up this into
			//       another pipeline function
			for _, spec := range decl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					cout <- typeSpec
				}
			}
		}
	}(cin, cout)
	return cout
}

// filter only StructType
func filterStructType(cin <-chan *ast.TypeSpec) <-chan *ast.StructType {
	cout := make(chan *ast.StructType)
	go func(cin <-chan *ast.TypeSpec, cout chan *ast.StructType) {
		defer close(cout)
		for spec := range cin {
			if structType, ok := spec.Type.(*ast.StructType); ok {
				cout <- structType
			}
		}
	}(cin, cout)
	return cout
}

// parse a type spec into pTypeSpec
func parseTypeSpec(typeSpec *ast.TypeSpec) (spec pTypeSpec) {

	// name of the type
	if typeSpec.Name != nil {
		spec.Name = typeSpec.Name.Name
	}

	// if this is a struct
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		spec.Type = "struct"

		// see FieldList, Field, Tag in package "os/ast"
		for _, f := range structType.Fields.List {
			fspec := pFieldSpec{}
			if len(f.Names) > 0 {
				fspec.Name = f.Names[0].Name
			}
			if f.Tag != nil {
				// remove the "`" from beginning and end
				fspec.Tag = f.Tag.Value[1 : len(f.Tag.Value)-1]
			}
			if f.Type != nil {
				if id, ok := f.Type.(*ast.Ident); ok {
					fspec.Type = id.Name
				}
			}
			spec.Fields = append(spec.Fields, fspec)
		}
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

// read pTypeSpec of given type names from a file
func readTypeFile(inputPath string, tns []string) (pkg string, specs []pTypeSpec, err error) {
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
	for ts := range filterTypeSpec(filterGenDecl(chanDecls(f.Decls))) {
		pts := parseTypeSpec(ts)
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
