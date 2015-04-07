package main

import (
	"go/ast"
)

// representation of a general type spec
type genTypeSpec struct {
	Name   string
	Type   string
	Fields []fieldSpec
}

// representation of a general struct field
type fieldSpec struct {
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

// parse a type spec into genTypeSpec
func parseTypeSpec(typeSpec *ast.TypeSpec) (spec genTypeSpec) {

	// name of the type
	if typeSpec.Name != nil {
		spec.Name = typeSpec.Name.Name
	}

	// if this is a struct
	if structType, ok := typeSpec.Type.(*ast.StructType); ok {
		spec.Type = "struct"

		// see FieldList, Field, Tag in package "os/ast"
		for _, f := range structType.Fields.List {
			fspec := fieldSpec{}
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
