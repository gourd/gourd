{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gourd/kit/perm"
	httpservice "github.com/gourd/kit/service/http"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

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
			httpservice.EnforceCreate(request)
			return inner(ctx, request)
		}
	}

	var prepareUpdate endpoint.Middleware = func(inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			sReq := request.(*httpservice.Request)
			el := &[]{{ .Type }}{}
			q := sReq.Query

			// get store
			s, err := getStore(ctx)
			if err != nil {
				serr := store.ErrorInternal
				serr.ServerMsg = fmt.Sprintf("error obtaining %T(%v) store (%s)", storeKey, storeKey, err)
				err = serr
				return
			}
			defer s.Close()

			// find the previous content of the id
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
				sReq.Previous = &(*el)[0]
			}

			// enforce agreement on sReq.Payload with previous sReq.Entity
			httpservice.EnforceUpdate(sReq.Previous, sReq.Payload)

			// placeholder: anything you want to do with the entity
			//              before update to database
			return inner(ctx, sReq)
		}
	}

	var preparePartialUpdate endpoint.Middleware = func(inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			sReq := request.(*httpservice.Request)
			el := &[]{{ .Type }}{}
			q := sReq.Query

			// get store
			s, err := getStore(ctx)
			if err != nil {
				serr := store.ErrorInternal
				serr.ServerMsg = fmt.Sprintf("error obtaining %T(%v) store (%s)", storeKey, storeKey, err)
				err = serr
				return
			}
			defer s.Close()

			// find the previous content of the id
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
				prev := (*el)[0]
				sReq.Previous = &prev
			}

			toUpdate := (*el)[0]

			if dec, ok := httpservice.PartialDecoderFrom(ctx); ok {
				err = dec.Decode(&toUpdate)
				if err != nil {
					// return validataion error directly
					return
				}
				sReq.Payload = &toUpdate
			}

			// enforce agreement on sReq.Payload with previous 	sReq.Entity
			httpservice.EnforceUpdate(sReq.Previous, sReq.Payload)

			// enforce agreement on sReq.Payload with previous sReq.Entity
			httpservice.EnforceUpdate(sReq.Previous, toUpdate)

			// placeholder: anything you want to do with the entity
			//              before update to database
			return inner(ctx, sReq)
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
				m := perm.GetMux(ctx)
				err = m.Allow(ctx, permission, request)
				if err != nil {
					return
				}
				return inner(ctx, request)
			}
		}
	}

	// generates request permission checker middleware
	checkPermAfter := func (permission string) endpoint.Middleware {
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

	//
	// ==== raw decode functions
	//

	decodeServiceIDReq := func(ctx context.Context, r *http.Request) (request *httpservice.Request, err error) {
		id := r.URL.Query().Get(":id") // will change
		cond := store.NewConds().Add("id", id)
		request = &httpservice.Request{
			Request: r,
			Query: store.NewQuery().SetConds(cond),
		}
		return
	}

	contextDecodeBody := func(ctx context.Context, r *http.Request) (entity *{{ .Type }}, err error) {

		// allocate entity
		entity = &{{ .Type }}{}

		if dec, ok := httpservice.DecoderFrom(ctx); ok {
			// allocate entity
			err = dec.Decode(entity)
			if err != nil {
				return
			}
		} else {
			err = fmt.Errorf("decoder not found in context")
		}
		return
	}

	//
	// ==== httptransport.DecodeRequestFunc implementations
	//

	// decodeIDReq generically decoded :id field
	// (works with pat based URL routing, router specific)
	var decodeIDReq httptransport.DecodeRequestFunc = func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		return decodeServiceIDReq(ctx, r)
	}

	// decodeListReq decode query for list endpoint
	var decodeListReq httptransport.DecodeRequestFunc = func(ctx context.Context, r *http.Request) (request interface{}, err error) {

		sReq := &httpservice.Request{
			Request: r,
			Query: store.NewQuery(),
		}

		// parse sort parameter
		sortStr := r.FormValue("_sort")
		if sortStr != "" {
			sorts := strings.Split(sortStr, ",")
			for _, sort := range sorts {
				sReq.Query.Sort(sort)
			}
		}

		// parse paging request parameter
		offset, limit := func(r *http.Request) (o, l uint64) {
			ostr := r.FormValue("_offset")
			lstr := r.FormValue("_limit")
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
		sReq.Query.SetOffset(offset)
		sReq.Query.SetLimit(limit)

		request = sReq
		return
	}

	// decodeJSONReq returns a DecodeRequestFunc that decode request
	// into allocated memory structure
	var decodeJSONReq httptransport.DecodeRequestFunc = func(ctx context.Context, r *http.Request) (request interface{}, err error) {
		return contextDecodeBody(ctx, r)
	}

	// decodeUpdate returns a DecodeRequestFunc that decode request
	var decodeUpdate httptransport.DecodeRequestFunc = func(ctx context.Context, r *http.Request) (request interface{}, err error) {

		sReq, err := decodeServiceIDReq(ctx, r)
		if err != nil {
			return
		}

		sReq.Payload, err = contextDecodeBody(ctx, r)
		if err != nil {
			return
		}

		request = sReq
		return
	}

	// decodePartialUpdate returns a DecodeRequestFunc that decode request
	var decodePartialUpdate httptransport.DecodeRequestFunc = func(ctx context.Context, r *http.Request) (request interface{}, err error) {

		sReq, err := decodeServiceIDReq(ctx, r)
		if err != nil {
			return
		}

		request = sReq
		return
	}

	//
	// ==== httpservce.Services
	//

	// define middleware chains of all RESTful endpoints
	handlers = make(map[string]*httpservice.Service)

	handlers["create"] = httpservice.NewJSONService(
		paths.Plural(), endpoints["create"])
	handlers["create"].Weight = 1
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

	handlers["update_partially"] = httpservice.NewJSONService(
		paths.Singular(), endpoints["update"])
	handlers["update_partially"].Methods = []string{"POST"}
	handlers["update_partially"].DecodeFunc = decodePartialUpdate
	handlers["update_partially"].Middlewares.Add(httpservice.MWProtocol, prepareProtocol)
	handlers["update_partially"].Middlewares.Add(httpservice.MWPrepare, preparePartialUpdate)
	handlers["update_partially"].Middlewares.Add(httpservice.MWInner,
		checkPermBefore("update "+noun.Singular()))

	handlers["list"] = httpservice.NewJSONService(
		paths.Plural(), endpoints["list"])
	handlers["list"].Weight = 1
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

// {{ .Type }}Rest binds store to pat router
func {{ .Type }}Rest(rf httpservice.RouterFunc, paths httpservice.Paths, patches ...httpservice.ServicesPatch) {

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
