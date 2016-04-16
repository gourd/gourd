{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

	"fmt"
{{ end }}

{{ define "code" }}

// {{ .Store.Name }}Endpoints return CURD endpoints for {{ .Store.Name }}
func {{ .Store.Name }}Endpoints(noun, nounp string) (endpoints map[string]endpoint.Endpoint) {

	// variables to use later
	allocEntityList := func() *[]{{ .Type }} { return &[]{{ .Type }}{} }
	storeKey := {{ .StoreKey }}
	getStore := func(ctx context.Context) (s *{{ .Store.Name }}, err error) {
		raw, err := store.Get(ctx, storeKey)
		if err != nil {
			return
		}

		s, ok := raw.(*{{ .Store.Name }})
		if !ok {
			err = fmt.Errorf("store.Get(\"{{ .Store.Name }}\") does not return *{{ .Store.Name }}")
			return
		}
		return
	}

	// store endpoints here
	// TODO: may have new struct to store
	endpoints = make(map[string]endpoint.Endpoint)

	endpoints["create"] = func(ctx context.Context, e interface{}) (res interface{}, err error) {

		// get context information
		r := gourdctx.HTTPRequest(ctx)
		if r == nil {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(ctx)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeKey, err)
			err = serr
			return
		}
		defer s.Close()

		// create entity
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
			nounp: &[]{{ .Type }}{*e.(*{{ .Type }})},
		}
		return

	}

	endpoints["retrieve"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		q := request.(store.Query)

		// get context information
		r := gourdctx.HTTPRequest(ctx)
		if r == nil {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(ctx)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeKey, err)
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
		r := gourdctx.HTTPRequest(ctx)
		if r == nil {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(ctx)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeKey, err)
			err = serr
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := allocEntityList()

		results := s.Search(q)
		count, err := results.Count()
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error counting %s: %s",
				noun, err)
			err = serr
			return
		}

		err = results.All(el)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// TODO: need to fix overflow error of pager variables
		res = map[string]interface{}{
			nounp: el,
			"paging": store.NewPager().
				SetTotal(int(count)).
				SetLimit(int(q.GetLimit()), int(q.GetOffset())),
		}
		return
	}

	endpoints["update"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		rmap := request.(map[string]interface{})
		q := rmap["query"].(store.Query)
		e := rmap["entity"]
		cond := q.GetConds()

		// get context information
		r := gourdctx.HTTPRequest(ctx)
		if r == nil {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(ctx)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeKey, err)
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
			nounp: &[]{{ .Type }}{*e.(*{{ .Type }})},
		}
		return
	}

	endpoints["delete"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		// allocate memory for variables
		el := allocEntityList()

		// store query of id
		q := request.(store.Query)
		cond := q.GetConds()

		// get context information
		r := gourdctx.HTTPRequest(ctx)
		if r == nil {
			serr := store.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get store
		s, err := getStore(ctx)
		if err != nil {
			serr := store.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s store (%s)", storeKey, err)
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
			nounp: el,
		}
		return
	}

	return
}

{{ end }}
