package main

import (
	"io"
	"os/exec"
	"text/template"
)

var tpls Templates

// Example use:
// ------------
//
// // Build template map
// tpls.Append("test", `{{ template "mytest" . }}`)
// tpls.Append("test subcmd1", `{{ define "mytest" }} hello test 1 {{ end }}`)
// tpls.Append("test subcmd2", `{{ define "mytest" }} hello test 2 {{ end }}`)
// tpls.AddDeps("test subcmd1", "test")
// tpls.AddDeps("test subcmd2", "test")
//
// // Render with selected template
// // Dependencies handled
// tpls.New("test subcmd1").
//    Execute(os.Stdout, nil)
//

// Templates is a generic text template map
// with dependency handling
type Templates struct {
	Tpls  map[string][]string
	Deps  map[string][]string
	Preps map[string]TemplatePrep
}

// TemplatePrep is function to be run on data
// everytime before a template is executed
type TemplatePrep func(in interface{}) (data interface{}, err error)

// Template is a generic text template
// with dependency handling
type Template struct {
	Tpls     *Templates
	Prep     TemplatePrep
	Used     map[string]bool
	Rendered *template.Template
}

// MustInit make sure the tpl map is ready
func (ts *Templates) MustInit(n string) *Templates {

	// make sure the tpl map is ready
	if ts.Tpls == nil {
		ts.Tpls = make(map[string][]string)
	}

	// make sure the tpl map is ready
	if ts.Deps == nil {
		ts.Deps = make(map[string][]string)
	}

	// make sure the prep map is ready
	if ts.Preps == nil {
		ts.Preps = make(map[string]TemplatePrep)
	}

	return ts
}

// Append add a template to the template map
func (ts *Templates) Append(n string, tpl ...string) *Templates {

	// make sure the tpl map is ready
	ts.MustInit(n)

	// append the sub-command to the command list
	ts.Tpls[n] = append(ts.Tpls[n], tpl...)

	return ts
}

// AddDeps define dependencies
func (ts *Templates) AddDeps(n string, deps ...string) *Templates {

	// make sure the tpl map is ready
	ts.MustInit(n)

	ts.Deps[n] = append(ts.Deps[n], deps...)
	return ts
}

// AddPrep adds a preparation function.
// The function will be run on the input parameters
// each time the template of the name is used
func (ts *Templates) AddPrep(n string, prep TemplatePrep) *Templates {
	// make sure the tpl map is ready
	ts.MustInit(n)

	ts.Preps[n] = prep
	return ts
}

// New create new template from template map
func (ts *Templates) New(name string) *Template {
	// initialize rendered template
	t := &Template{
		Tpls:     ts,
		Used:     make(map[string]bool),
		Rendered: template.New(name),
	}

	// define preparation function
	if prep, ok := ts.Preps[name]; ok {
		t.Prep = prep
	} else {
		t.Prep = func(in interface{}) (interface{}, error) {
			return in, nil
		}
	}

	// render the template
	t.Use(name)
	return t
}

// Use a templates in template map of the name
func (t *Template) Use(n string) *Template {
	// prevent circular dependencies
	if _, ok := t.Used[n]; ok {
		return t
	}
	t.Used[n] = true

	// resolve dependencies
	if deps, ok := t.Tpls.Deps[n]; ok {
		for _, dep := range deps {
			t.Use(dep)
		}
	}

	// use the template
	if _, ok := t.Tpls.Tpls[n]; ok {
		for _, tpl := range t.Tpls.Tpls[n] {
			t.Rendered = template.Must(t.Rendered.Parse(tpl))
		}
	}
	return t
}

// Execute the rendered template
func (t *Template) Execute(out io.Writer, data interface{}) error {
	data, err := t.Prep(data)
	if err != nil {
		return err
	}
	return t.Rendered.Execute(out, data)
}

// FormatFile formats a generated file with go formatter
func FormatFile(fn string) error {
	return exec.Command("go", "fmt", fn).Run()
}
