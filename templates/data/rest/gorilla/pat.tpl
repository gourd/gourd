{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/perm"
	httpservice "github.com/gourd/kit/service/http"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
{{ end }}

{{ define "code" }}

func {{ .Store }}Services(paths httpservice.Paths, endpoints map[string]endpoint.Endpoint) (handlers httpservice.Services) {

	// variables to use later
	noun := paths.Noun()
	storeKey := {{ .StoreKey }}
	getStore := func(ctx context.Context) (s *{{ .Store }}, err error) {
		raw, err := store.Get(ctx, storeKey)
		if err != nil {
			return
		}

		s, ok := raw.(*{{ .Store }})
		if !ok {
			err = fmt.Errorf(`store.Get({{ .StoreKey }}) does not return *{{ .StoreKey }}`)
			return
		}
		return
	}


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
			r := gourdctx.HTTPRequest(ctx)
			if r == nil {
				serr := store.ErrorInternal
				serr.ServerMsg = "missing request in context"
				err = serr
				return
			}

			el := &[]{{ .Type }}{}
			q := rmap["query"].(store.Query)

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
					noun.Singular(), err)
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
			list := vmap[noun.Plural()].(*[]{{ .Type }})
			if list == nil || *list == nil {
				*list = make([]{{ .Type }}, 0)
			}
			vmap[noun.Plural()] = list

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
				response = store.NewResponse(noun.Plural(), v)
			}

			return
		}
	}

	// generates response permission checker middleware
	checkPermBefore := func (permission string) endpoint.Middleware {
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
	checkPermAfter := func (permission string) endpoint.Middleware {
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

	//
	// ====
	//

	// decodeIDReq generically decoded :id field
	// (works with pat based URL routing, router specific)
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

	// define middleware chains of all RESTful endpoints
	handlers = make(map[string]*httpservice.Service)

	log.Printf("path: %#v, create", paths.Plural())
	handlers["create"] = httpservice.NewJSONService(
		paths.Plural(), endpoints["create"])
	handlers["create"].Methods = []string{"POST"}
	handlers["create"].DecodeFunc = decodeJSONReq
	handlers["create"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["create"].Middlewares.Add(httpservice.MWPrepare, prepareCreate)
	handlers["create"].Middlewares.Add(httpservice.MWInner,
		checkPermBefore("create "+noun.Singular()))

	handlers["retrieve"] = httpservice.NewJSONService(
		paths.Singular(), endpoints["retrieve"])
	handlers["retrieve"].Methods = []string{"GET"}
	handlers["retrieve"].DecodeFunc = decodeIDReq
	handlers["retrieve"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["retrieve"].Middlewares.Add(httpservice.MWPrepare, prepareList)
	handlers["retrieve"].Middlewares.Add(httpservice.MWInner,
		checkPermAfter("retrieve "+noun.Singular()))


	handlers["update"] = httpservice.NewJSONService(
		paths.Singular(), endpoints["update"])
	handlers["update"].Methods = []string{"PUT"}
	handlers["update"].DecodeFunc = decodeUpdate
	handlers["update"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["update"].Middlewares.Add(httpservice.MWPrepare, prepareUpdate)
	handlers["update"].Middlewares.Add(httpservice.MWInner,
		checkPermBefore("update "+noun.Singular()))


	handlers["list"] = httpservice.NewJSONService(
		paths.Plural(), endpoints["list"])
	handlers["list"].Methods = []string{"GET"}
	handlers["list"].DecodeFunc = decodeListReq
	handlers["list"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["list"].Middlewares.Add(httpservice.MWPrepare, prepareList)
	handlers["list"].Middlewares.Add(httpservice.MWInner,
		checkPermAfter("list "+noun.Singular()))


	handlers["delete"] = httpservice.NewJSONService(
		paths.Singular(), endpoints["delete"])
	handlers["delete"].Methods = []string{"DELETE"}
	handlers["delete"].DecodeFunc = decodeIDReq
	handlers["delete"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["delete"].Middlewares.Add(httpservice.MWInner,
		checkPermBefore("delete "+noun.Singular()))

	return
}

// {{ .Store }}Rest binds store to pat router
func {{ .Store }}Rest(rf httpservice.RouterFunc, paths httpservice.Paths, patches ...httpservice.ServicesPatch) {

	log.Printf("REST path: %s", paths.Plural())

	// generate CRUD endpoints
	endpoints := {{ .Store }}Endpoints(
		paths.Noun().Singular(), paths.Noun().Plural())

	// generate service description (Middleware, DecodeRequestFunc)
	services := {{ .Store }}Services(paths, endpoints)
	services.Patch(patches...)

	// route to all services
	if err := services.Route(rf); err != nil {
		panic(err)
	}

}

{{ end }}
