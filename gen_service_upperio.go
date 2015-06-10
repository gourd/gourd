package main

import (
	"fmt"
	"github.com/gourd/goparser"
)

func init() {
	tpls.Append("gen service:upperio",
		`{{ define "package" }}package {{ .Pkg }}{{ end }}`,
		`{{ define "imports" }}"net/http"
		"github.com/gourd/service"
		"github.com/gourd/service/upperio"
		"log"
		"upper.io/db"{{ end }}`,
		`{{ define "code" }}


func init() {
	service.Providers.Define("{{ .Type.Name }}", Get{{ .Type.Name }}Service)
}

// Get{{ .Type.Name }}Service implements service.ProviderFunc
func Get{{ .Type.Name }}Service(r *http.Request) (s service.Service, err error) {

	// obtain database
	db, err := upperio.Open(r, "default")
	if err != nil {
		return
	}

	// define service and return
	s = &{{ .Type.Name }}Service{db}
	return
}

// {{ .Type.Name }}Service serves generic CURD for type {{ .Type.Name }}
// Generated by gourd CLI tool
type {{ .Type.Name }}Service struct {
	Db db.Database
}

// Create a {{ .Type.Name }} in the database, of the parent
func (s *{{ .Type.Name }}Service) Create(
	cond service.Conds, ep service.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	//TODO: convert cond into parentkey and
	//      enforce to the entity

	// add the entity to collection
	id, err := coll.Append(ep)
	if err != nil {
		log.Printf("Error creating {{ .Type.Name }}: %s", err.Error())
		err = service.ErrorInternal
		return
	}

	//TODO: apply the key to the entity
	e := ep.(*{{.Type.Name}})
	e.{{ .Id.Name }} = {{ .Id.Type }}(id.({{ .Id.DbType }}))

	return
}

// Search a {{ .Type.Name }} by its condition(s)
func (s *{{ .Type.Name }}Service) Search(
	c service.Conds, lp service.EntityListPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get list condition and ignore the error
	cond, _ := c.GetMap()

	// retrieve all users
	var res db.Result
	if len(cond) == 0 {
		res = coll.Find()
	} else {
		res = coll.Find(db.Cond(cond))
	}

	// handle paging
	if c.GetOffset() != 0 {
		res = res.Skip(uint(c.GetOffset()))
	}
	if c.GetLimit() != 0 {
		res = res.Limit(uint(c.GetLimit()))
	}

	// TODO: also work with c.Cond for ListCond (limit and offset)
	err = res.All(lp)
	if err != nil {
		err = service.ErrorInternal
	}

	return nil
}

// One returns the first {{ .Type.Name }} matches condition(s)
func (s *{{ .Type.Name }}Service) One(
	c service.Conds, ep service.EntityPtr) (err error) {

	// retrieve from database
	l := &[]{{ .Type.Name }}{}
	err = s.Search(c, l)
	if err != nil {
		return
	}

	// if not found, report
	if len(*l) == 0 {
		err = service.ErrorNotFound
		return
	}

	// assign the value of given point
	// to the first retrieved value
	(*ep.(*{{ .Type.Name }})) = (*l)[0]
	return nil
}

// Update {{ .Type.Name }} on condition(s)
func (s *{{ .Type.Name }}Service) Update(
	c service.Conds, ep service.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get by condition and ignore the error
	cond, _ := c.GetMap()
	res := coll.Find(db.Cond(cond))

	// update the matched entities
	err = res.Update(ep)
	if err != nil {
		log.Printf("Error updating {{ .Type.Name }}: %s", err.Error())
		err = service.ErrorInternal
	}
	return
}

// Delete {{ .Type.Name }} on condition(s)
func (s *{{ .Type.Name }}Service) Delete(
	c service.Conds) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get by condition and ignore the error
	cond, _ := c.GetMap()
	res := coll.Find(db.Cond(cond))

	// remove the matched entities
	err = res.Remove()
	if err != nil {
		log.Printf("Error deleting {{ .Type.Name }}: %s", err.Error())
		err = service.ErrorInternal
	}
	return nil
}

// AllocEntity allocate memory for an entity
func (s *{{ .Type.Name }}Service) AllocEntity() service.EntityPtr {
	return &{{ .Type.Name }}{}
}

// AllocEntityList allocate memory for an entity list
func (s *{{ .Type.Name }}Service) AllocEntityList() service.EntityListPtr {
	return &[]{{ .Type.Name }}{}
}

// Len inspect the length of an entity list
func (s *{{ .Type.Name }}Service) Len(pl service.EntityListPtr) int64 {
	el := pl.(*[]{{ .Type.Name }})
	return int64(len(*el))
}

// Coll return the raw upper.io collection
func (s *{{ .Type.Name }}Service) Coll() (coll db.Collection, err error) {
	// get raw collection
	coll, err = s.Db.Collection("{{.Coll}}")
	if err != nil {
		log.Printf("Error connecting collection {{.Coll}}: %s",
			err.Error())
		err = service.ErrorInternal
	}
	return 
}

{{ end }}`)

	tpls.AddDeps("gen service:upperio", "gen:general")

	tpls.AddPrep("gen service:upperio", func(in interface{}) (interface{}, error) {
		var data map[string]interface{}
		var f interface{}
		var id *goparser.FieldSpec
		var ok bool

		if data, ok = in.(map[string]interface{}); !ok {
			return in, fmt.Errorf("Unable to prepare. Incorrect data provided: %#v", in)
		}

		if f, ok = data["Id"]; !ok {
			return in, fmt.Errorf("Unable to prepare. No Id found in given data: %#v",
				data)
		}

		if id, ok = f.(*goparser.FieldSpec); !ok {
			return in, fmt.Errorf("Unable to prepare. Wrong Id type: %#v", f)
		}

		// determine database return type
		// case by case
		dbtype := "int64"
		if id.Type == "string" {
			dbtype = "string"
		}

		// override the field spec
		data["Id"] = &UpperFieldSpec{
			id,
			dbtype,
		}
		return data, nil

	})

}

// UpperFieldSpec represents information of an upper field.
// Includes its original field spec and database type suggestion
type UpperFieldSpec struct {
	*goparser.FieldSpec
	DbType string // type of id returned by upperio
}
