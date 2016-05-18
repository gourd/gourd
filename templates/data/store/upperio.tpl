{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/gourd/kit/store"
	"github.com/gourd/kit/store/upperio"
	{{ if .Id.IsString }}"github.com/satori/go.uuid"{{ end }}
	"github.com/go-kit/kit/log"

	{{ if .Id.IsString }}
	"encoding/base64"
	{{ end }}
	"fmt"
	"io/ioutil"
	"upper.io/db"
{{ end }}

{{ define "code" }}

// {{ .Type.Name }}StoreProvider implements store.Provider interface
// provides raw {{ .Type.Name }}Store
func {{ .Type.Name }}StoreProvider(sess interface{}) (s store.Store, err error) {

	var dbSess db.Database
	var ok bool

	logger := log.NewLogfmtLogger(ioutil.Discard)

	if dbSess, ok = sess.(db.Database); !ok {
		err = fmt.Errorf("expected db.Database in sess, got %#v", sess)
		return
	}

	// define store and return
	s = &{{ .Type.Name }}Store{dbSess, logger}
	return
}

// {{ .Type.Name }}Store serves generic CURD for type {{ .Type.Name }}
// Generated by gourd CLI tool
type {{ .Type.Name }}Store struct {
	Db     db.Database
	logger log.Logger
}

// Create a {{ .Type.Name }} in the database, of the parent
func (s *{{ .Type.Name }}Store) Create(
	cond store.Conds, ep store.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// apply random uuid string to string id
	{{ if .Id.IsString }}
	uid := uuid.NewV4()
	e := ep.(*{{.Type.Name}})
	e.{{ .Id.Name }} = base64.RawURLEncoding.EncodeToString(uid[:])
	{{ end }}

	// Marshal the item, if possible
	// (quick fix for upperio problem with db.Marshaler)
	if me, ok := ep.(db.Marshaler); ok {
		ep, err = me.MarshalDB()
		if err != nil {
			return
		}
	}

	// add the entity to collection
	{{ if .Id.IsString }}
	_, err = coll.Append(ep)
	{{ else }}
	id, err := coll.Append(ep)
	{{ end }}
	if err != nil {
		err = s.errorf("Error creating {{ .Type.Name }}: %s", err.Error())
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
func (s *{{ .Type.Name }}Store) Search(
	q store.Query) store.Result {

	return upperio.NewResult(func() (res db.Result, err error) {
		// get collection
		coll, err := s.Coll()
		if err != nil {
			return
		}

		// retrieve entities by given query conditions
		conds := upperio.Conds(q.GetConds())
		if conds == nil {
			res = coll.Find()
		} else {
			res = coll.Find(conds)
		}

		// add sorting information, if any
		res = res.Sort(upperio.Sort(q)...)

		// handle paging
		if q.GetOffset() != 0 {
			res = res.Skip(uint(q.GetOffset()))
		}
		if q.GetLimit() != 0 {
			res = res.Limit(uint(q.GetLimit()))
		}

		return
	})

}

// One returns the first {{ .Type.Name }} matches condition(s)
func (s *{{ .Type.Name }}Store) One(
	c store.Conds, ep store.EntityPtr) (err error) {

	// retrieve results from database
	l := &[]{{ .Type.Name }}{}
	q := store.NewQuery().SetConds(c)

	// dump results into pointer of map / struct
	err = s.Search(q).All(l)
	if err != nil {
		return
	}

	// if not found, report
	if len(*l) == 0 {
		err = store.ErrorNotFound
		return
	}

	// assign the value of given point
	// to the first retrieved value
	(*ep.(*{{ .Type.Name }})) = (*l)[0]
	return nil
}

// Update {{ .Type.Name }} on condition(s)
func (s *{{ .Type.Name }}Store) Update(
	c store.Conds, ep store.EntityPtr) (err error) {

	// get collection
	coll, err := s.Coll()
	if err != nil {
		return
	}

	// get by condition and ignore the error
	cond, _ := c.GetMap()
	res := coll.Find(db.Cond(cond))

	// Marshal the item, if possible
	// (quick fix for upperio problem with db.Marshaler)
	if me, ok := ep.(db.Marshaler); ok {
		ep, err = me.MarshalDB()
		if err != nil {
			return
		}
	}

	// update the matched entities
	err = res.Update(ep)
	if err != nil {
		err = s.errorf("Error updating {{ .Type.Name }}: %s", err.Error())
	}
	return
}

// Delete {{ .Type.Name }} on condition(s)
func (s *{{ .Type.Name }}Store) Delete(
	c store.Conds) (err error) {

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
		err = s.errorf("Error deleting {{ .Type.Name }}: %s", err.Error())
	}
	return nil
}

// AllocEntity allocate memory for an entity
func (s *{{ .Type.Name }}Store) AllocEntity() store.EntityPtr {
	return &{{ .Type.Name }}{}
}

// AllocEntityList allocate memory for an entity list
func (s *{{ .Type.Name }}Store) AllocEntityList() store.EntityListPtr {
	return &[]{{ .Type.Name }}{}
}

// Len inspect the length of an entity list
func (s *{{ .Type.Name }}Store) Len(pl store.EntityListPtr) int64 {
	el := pl.(*[]{{ .Type.Name }})
	return int64(len(*el))
}

// Coll return the raw upper.io collection
func (s *{{ .Type.Name }}Store) Coll() (coll db.Collection, err error) {
	// get raw collection
	coll, err = s.Db.Collection("{{.Coll}}")
	if err != nil {
		err = s.errorf("Error connecting collection {{.Coll}}: %s",
			err.Error())
	}
	return
}

// SetLogger set the logger fotr the {{ .Type.Name }}Store
func (s *{{ .Type.Name }}Store) SetLogger(logger log.Logger) {
	s.logger = logger
}

// Log logs the message with session id
func (s *{{ .Type.Name }}Store) error(msg string) error {
	serr := store.ErrorInternal
	serr.ServerMsg = msg
	s.logger.Log("store", "{{ .Type.Name }}Store", "message", msg)
	return serr
}

// Logf logs the message with session id
func (s *{{ .Type.Name }}Store) errorf(msg string, v ...interface{}) error {
	return s.error(fmt.Sprintf(msg, v...))
}

// Close would not close database connection at all.
// Please use store.CloseAllIn(ctx) to wrap up connections
// in a context
func (s *{{ .Type.Name }}Store) Close() error {
	return nil
}


{{ end }}
