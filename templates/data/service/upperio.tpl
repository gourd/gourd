{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"net/http"
	"github.com/gourd/service"
	"github.com/gourd/service/upperio"
	{{ if .Id.IsString }}
	"github.com/satori/go.uuid"
	"encoding/base64"
	"strings"
	{{ end }}
	"log"
	"upper.io/db"
{{ end }}

{{ define "code" }}

func init() {
	// define service provider with proxy
	service.Providers.DefineFunc("{{ .Type.Name }}", func (r *http.Request) (s service.Service, err error) {
		return Get{{ .Type.Name }}Service(r)
	})
}

// Get{{ .Type.Name }}Service provides raw {{ .Type.Name }}Service
func Get{{ .Type.Name }}Service(r *http.Request) (s *{{ .Type.Name }}Service, err error) {

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

	// apply random uuid string to string id
	{{ if .Id.IsString }}
	uid := uuid.NewV4()
	e := ep.(*{{.Type.Name}})
	e.{{ .Id.Name }} = strings.TrimRight(base64.URLEncoding.EncodeToString(uid[:]), "=")
	{{ end }}


	//TODO: convert cond into parentkey and
	//      enforce to the entity

	// add the entity to collection
	{{ if .Id.IsString }}
	_, err = coll.Append(ep)
	{{ else }}
	id, err := coll.Append(ep)
	{{ end }}
	if err != nil {
		log.Printf("Error creating {{ .Type.Name }}: %s", err.Error())
		err = service.ErrorInternal
		return
	}

	{{ if .Id.IsString }}{{ else }}
	// apply the key to the entity
	e := ep.(*{{.Type.Name}})
	e.{{ .Id.Name }} = {{ .Id.Type }}(id.(int64))
	{{ end }}

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

// Close the database session that {{ .Type.Name }} is using
func (s *{{ .Type.Name }}Service) Close() error {
	return s.Db.Close()
}


{{ end }}
