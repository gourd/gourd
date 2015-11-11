{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/pat"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/perm"
	httpservice "github.com/gourd/kit/service/http"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
{{ end }}

{{ define "code" }}

// {{ .Store }}AppendDesc(paths httpservice.Paths) append endpoint descriptor
// with endpoint-specific middleware and DecodeRequestFunc for RESTful HTTP service
func {{ .Store }}AppendDesc(desc httpservice.Desc) httpservice.Desc {

	// variables to use later
	getStore := Get{{ .Store }}
	storeName := "{{ .Store }}"

	// define RESTful descriptor
	paths := desc.Paths()

	// define default middlewares
	var prepareCreate endpoint.Middleware = func(inner endpoint.Endpoint) endpoint.Endpoint {
		return func (ctx context.Context, request interface{}) (respond interface{}, err error) {
			// placeholder: anything you want to do with the entity
			//              before append to database
			return inner(ctx, request)
		}
	}

	var prepareUpdate endpoint.Middleware = func(inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			rmap := request.(map[string]interface{})

			// get context information
			r, ok := gourdctx.HTTPRequest(ctx)
			if !ok {
				serr := store.ErrorInternal
				serr.ServerMsg = "missing request in context"
				err = serr
				return
			}

			el := &[]{{ .Type }}{}
			q := rmap["query"].(store.Query)

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
					paths.Noun().Singular(), err)
				err = serr
				return
			}

			// tell the inner
			if len(*el) > 0 {
				rmap["prev"] = (*el)[0]
			}

			// placeholder: anything you want to do with the entity
			//              before update to database
			return inner(ctx, rmap)
		}
	}

	var prepareList endpoint.Middleware = func (inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = inner(ctx, request)
			if err != nil {
				return
			}

			vmap := response.(map[string]interface{})
			list := vmap[paths.Noun().Plural()].(*[]{{ .Type }})
			if list == nil || *list == nil {
				*list = make([]{{ .Type }}, 0)
			}
			vmap[paths.Noun().Plural()] = list

			// placeholder: anything you want to do with the entity
			//              list response
			return vmap, nil
		}
	}

	// wrap inner response with default protocol
	var prepareProtocol endpoint.Middleware = func (inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			v, err := inner(ctx, request)
			if err != nil {
				return
			}

			switch v.(type) {
			case map[string]interface{}:
				response = store.ExpandResponse(v.(map[string]interface{}))
			default:
				response = store.NewResponse(paths.Noun().Plural(), v)
			}

			return
		}
	}

	// generates response permission checker middleware
	genResponsePermChecker := func (permission string) endpoint.Middleware {
		return func (inner endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (response interface{}, err error) {

				v, err := inner(ctx, request)
				if err != nil {
					return
				}

				m := perm.GetMux(ctx)
				err = m.Allow(ctx, permission, request, v)
				if err != nil {
					return
				}

				response = v
				return

			}
		}
	}

	// generates request permission checker middleware
	genRequestPermChecker := func (permission string) endpoint.Middleware {
		return func (inner endpoint.Endpoint) endpoint.Endpoint {
			return func(ctx context.Context, request interface{}) (response interface{}, err error) {
				m := perm.GetMux(ctx)
				err = m.Allow(ctx, permission, request)
				if err != nil {
					return
				}
				return inner(ctx, request)
			}
		}
	}

	// define middleware chains of all RESTful endpoints
	desc.SetMiddleware("create", endpoint.Chain(
		prepareProtocol,
		prepareCreate,
		genRequestPermChecker("create "+paths.Noun().Singular())))
	desc.SetMiddleware("retrieve", endpoint.Chain(
		prepareProtocol,
		prepareList,
		genResponsePermChecker("retrieve "+paths.Noun().Singular())))
	desc.SetMiddleware("update", endpoint.Chain(
		prepareProtocol,
		prepareUpdate,
		genRequestPermChecker("update "+paths.Noun().Singular())))
	desc.SetMiddleware("list", endpoint.Chain(
		prepareProtocol,
		prepareList,
		genResponsePermChecker("list "+paths.Noun().Singular())))
	desc.SetMiddleware("delete", endpoint.Chain(
		prepareProtocol,
		genRequestPermChecker("delete "+paths.Noun().Singular())))

	//
	// ====
	//

	// decodeIDReq generically decoded :id field
	// (works with pat based URL routing)
	var decodeIDReq httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {
		id := r.URL.Query().Get(":id") // will change
		cond := store.NewConds().Add("id", id)
		request = store.NewQuery().SetConds(cond)
		return
	}

	// decodeListReq decode query for list endpoint
	var decodeListReq httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {

		q := store.NewQuery()

		// parse sort parameter
		sortStr := r.FormValue("sorts")
		if sortStr != "" {
			sorts := strings.Split(sortStr, ",")
			for _, sort := range sorts {
				q.Sort(sort)
			}
		}

		// parse paging request parameter
		offset, limit := func(r *http.Request) (o, l uint64) {
			ostr := r.FormValue("offset")
			lstr := r.FormValue("limit")
			if ostr != "" {
				if ot, err := strconv.ParseUint(ostr, 10, 64); err == nil {
					o = ot
				}
			}
			if lstr != "" {
				if lt, err := strconv.ParseUint(lstr, 10, 64); err == nil {
					l = lt
				}
			}
			return
		}(r)

		// retrieve
		q.SetOffset(offset)
		q.SetLimit(limit)

		request = q
		return
	}

	// decodeJSONReq returns a DecodeRequestFunc that decode request
	// into allocated memory structure
	var decodeJSONReq httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {
		// allocate entity
		request = &{{ .Type }}{}

		// decode request
		dec := json.NewDecoder(r.Body)
		err = dec.Decode(request)
		return
	}

	// decodeUpdate returns a DecodeRequestFunc that decode request
	var decodeUpdate httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {
		rmap := make(map[string]interface{})

		rmap["entity"], err = decodeJSONReq(r)
		if err != nil {
			return
		}

		rmap["query"], err = decodeIDReq(r)
		if err != nil {
			return
		}

		rmap["prev"] = nil

		request = rmap
		return
	}

	desc.SetDecodeFunc("create", decodeJSONReq)
	desc.SetDecodeFunc("retrieve", decodeIDReq)
	desc.SetDecodeFunc("list", decodeListReq)
	desc.SetDecodeFunc("update", decodeUpdate)
	desc.SetDecodeFunc("delete", decodeIDReq)

	return desc
}

// {{ .Store }}Routes defines the router specific routes
// to be use in the {{ .Store }} Paths
func {{ .Store }}Routes(name string, noun httpservice.Noun) (string, string) {
	// define router specific paths for different action type
	// with the method of that action endpoint
	switch name {
		case "update":
			return path.Join(noun.Singular(), "{id}"), "PUT"
		case "retrieve":
			return path.Join(noun.Singular(), "{id}"), "GET"
		case "delete":
			return path.Join(noun.Singular(), "{id}"), "DELETE"
		case "list":
			return noun.Plural(), "GET"
		case "create":
			return noun.Plural(), "POST"
	}
	return "", ""
}

// {{ .Store }}Rest binds store to pat router
func {{ .Store }}Rest(r *pat.Router, p perm.Mux, base, noun, nounp string) {

	log.Printf("REST path: %s/%s", base, noun)

	// handle individual route with pat (router specific)
	patRouterFunc := func(r *pat.Router) (func(string, httpservice.Paths, http.Handler) error) {
		return func(name string, paths httpservice.Paths, h http.Handler) error {
			r.Add(paths.Method(name), paths.Path(name), h)
			return nil
		}
	}

	// generate CRUD endpoints
	endpoints := {{ .Store }}Endpoints(noun, nounp)

	// generate desc (Middleware, DecodeRequestFunc)
	// for the CRUD endpoints
	paths := httpservice.NewPaths(base,
		httpservice.NewNoun(noun, nounp),
		{{ .Store }}Routes)
	desc := httpservice.NewDesc(paths)
	desc = {{ .Store }}AppendDesc(desc)

	mware := endpoint.Chain(
		gourdctx.ClearGorilla,
		perm.UseMux(p))

	func(rf func(string, httpservice.Paths, http.Handler) error, mware endpoint.Middleware, desc httpservice.Desc) {

		// encodeJSONResp encodes given response into JSON
		encodeJSONResp := func(w http.ResponseWriter, response interface{}) (err error) {
			enc := json.NewEncoder(w)
			err = enc.Encode(response)
			return
		}

		// JSONErrorEncoder expands given error to StoreError then encode to JSON
		JSONErrorEncoder := func(w http.ResponseWriter, err error) {

			// quick fix for gokit bad request wrapping problem
			switch err.(type) {
			case httptransport.BadRequestError:
				err = err.(httptransport.BadRequestError).Err
			}

			enc := json.NewEncoder(w)
			err = enc.Encode(store.ExpandError(err))
		}

		// route all endpoints
		for name := range endpoints {

			// pre-wrap endpoint with endpoint specific middleware
			ep := desc.GetMiddleware(name)(endpoints[name])

			// pre-wrap endpoint with general middleware
			ep = endpoint.Chain(
				gourdctx.ClearGorilla,
				perm.UseMux(p))(ep)

			// generate http handler
			h := httptransport.NewServer(
				gourdctx.NewEmpty(),
				ep,
				desc.GetDecodeFunc(name),
				encodeJSONResp,
				httptransport.ServerBefore(gourdctx.UseGorilla),
				httptransport.ServerErrorEncoder(JSONErrorEncoder))

			// route the handler according to path and method
			if err := rf(name, desc.Paths(), h); err != nil {
				panic(err)
			}
		}

	}(patRouterFunc(r), mware, desc)

}

{{ end }}
