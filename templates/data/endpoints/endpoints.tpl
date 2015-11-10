{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

	"fmt"
	"log"
{{ end }}

{{ define "code" }}

// {{ .Store.Name }}Endpoints return CURD endpoints for {{ .Store.Name }}
func {{ .Store.Name }}Endpoints(noun, nounp string) (endpoints map[string]endpoint.Endpoint) {

	// variables to use later
	allocEntity := func() *{{ .Type }} { return &{{ .Type }}{} }
	allocEntityList := func() *[]{{ .Type }} { return &[]{{ .Type }}{} }
	getStore := Get{{ .Store.Name }}
	storeName := "{{ .Store.Name }}"

	// store endpoints here
	// TODO: may have new struct to store
	endpoints = make(map[string]endpoint.Endpoint)

	endpoints["create"] = func(ctx context.Context, e interface{}) (res interface{}, err error) {

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(r)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeName, err)
			err = serr
			return
		}
		defer s.Close()

		// create entity
		log.Printf("create: %#v", e)
		err = s.Create(nil, e)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error creating %s: %#v, entity: %#v", noun, err.Error(), e)
			err = serr
			return
		}

		// encode response
		res = map[string]interface{}{
			nounp: []interface{}{e},
		}
		return

	}

	endpoints["retrieve"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		q := request.(store.Query)

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(r)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeName, err)
			err = serr
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := allocEntityList()

		// retrieve
		err = s.Search(q).All(el)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// encode response
		if s.Len(el) == 0 {
			err = store.ErrorNotFound
			return
		}

		res = map[string]interface{}{
			nounp: el,
		}
		return

	}

	endpoints["list"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		q := request.(store.Query)

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(r)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeName, err)
			err = serr
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := allocEntityList()

		err = s.Search(q).All(el)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		res = map[string]interface{}{
			nounp: el,
		}
		return
	}

	endpoints["update"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		rmap := request.(map[string]interface{})
		q := rmap["query"].(store.Query)
		e := rmap["entity"]
		cond := q.GetConds()

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(r)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeName, err)
			err = serr
			return
		}
		defer s.Close()

		// update entity
		if err = s.Update(cond, e); err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		}

		res = map[string]interface{}{
			nounp: []interface{}{e},
		}
		return
	}

	endpoints["delete"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		// allocate memory for variables
		e := allocEntity()
		el := allocEntityList()

		// store query of id
		q := request.(store.Query)
		cond := q.GetConds()

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(r)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeName, err)
			err = serr
			return
		}
		defer s.Close()

		// find the content of the id
		err = s.Search(q).All(el)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// delete entity
		if err = s.Delete(cond); err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		}

		res = map[string]interface{}{
			nounp: []interface{}{e},
		}
		return
	}

	return
}

{{ end }}
