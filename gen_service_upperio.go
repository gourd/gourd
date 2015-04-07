package main

func init() {
	tpls.Append("gen service:upperio",
		`{{ define "package" }}package {{ .Type.Pkg }}{{ end }}`,
		`{{ define "imports" }}"github.com/gourd/service"
		"upper.io/db"{{ end }}`,
		`{{ define "code" }}

// {{ .Type.Name }}Service serves generic CURD for type {{ .Type.Name }}
// Generated by gourd CLI tool
type {{ .Type.Name }}Service struct {
	Db db.Database
}

// Create a {{ .Type.Name }} in the database, of the parent
func (t *{{ .Type.Name }}Service) Create(pk service.ParentKey, ep service.EntityPtr) error {
	//TODO: implement this
	return nil
}

// Search a {{ .Type.Name }} by its condition
func (t *{{ .Type.Name }}Service) Search(
	cond service.Conds, elp service.EntityListPtr) error {
	//TODO: implement this
	return nil
}

func (t *{{ .Type.Name }}Service) List(
	pk service.ParentKey, elp service.EntityListPtr) error {
	//TODO: implement this
	return nil
}

func (t *{{ .Type.Name }}Service) Retrieve(
	k service.Key, pk service.ParentKey, elp service.EntityListPtr) error {
	//TODO: implement this
	return nil
}

func (t *{{ .Type.Name }}Service) Update(
	k service.Key, pk service.ParentKey, ep service.EntityPtr) error {
	//TODO: implement this
	return nil
}

func (t *{{ .Type.Name }}Service) Delete(
	k service.Key, pk service.ParentKey) error {
	//TODO: implement this
	return nil
}

// AllocEntity allocate memory for an entity
func (t *{{ .Type.Name }}Service) AllocEntity() service.EntityPtr {
	return &{{ .Type.Name }}{}
}

// AllocEntity allocate memory for an entity list
func (t *{{ .Type.Name }}Service) AllocEntityList() service.EntityListPtr {
	return &[]{{ .Type.Name }}{}
}

// Len inspect the length of an entity list
func (t *{{ .Type.Name }}Service) Len(pl service.EntityListPtr) int64 {
	el := pl.(*[]{{ .Type.Name }})
	return int64(len(*el))
}

{{ end }}`)
	tpls.AddDeps("gen service:upperio", "gen:general")
}
