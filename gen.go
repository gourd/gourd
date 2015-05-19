package main

func init() {
	tpls.Append("gen:general",
		`{{ if .Ver }}// Generated by gourd (version {{ .Ver }}){{ end }}
{{ if .Now }}// Generated at {{ .Now }}{{ end }}
// Note: If you want to re-generate this file in the future,
//       do not change it.

{{ template "package" . }}

import (
{{ template "imports" . }}
)
{{ template "code" . }}
	`)
}