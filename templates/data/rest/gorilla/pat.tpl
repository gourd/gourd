{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/pat"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/perm"
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

// {{ .Store }}Rest binds store to pat router
func {{ .Store }}Rest(r *pat.Router, p perm.Mux, base, noun, nounp string) {

	// variables to use later
	allocEntity := func() *{{ .Type }} { return &{{ .Type }}{} }
	allocEntityList := func() *[]{{ .Type }} { return &[]{{ .Type }}{} }
	getStore := Get{{ .Store }}
	storeName := "{{ .Store }}"

	// enforce entity property before create
	var prepareCreate endpoint.Middleware = func(inner endpoint.Endpoint) endpoint.Endpoint {
		return func (ctx context.Context, request interface{}) (respond interface{}, err error) {
			// placeholder: anything you want to do with the entity
			//              before append to database
			return inner(ctx, request)
		}
	}

	// enforce entity property before update
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

			el := allocEntityList()
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
					noun, err)
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

	// prepare list
	var prepareList endpoint.Middleware = func (inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			response, err = inner(ctx, request)
			if err != nil {
				return
			}

			list := response.(*[]{{ .Type }})
			if list == nil || *list == nil {
				*list = make([]{{ .Type }}, 0)
			}

			// placeholder: anything you want to do with the entity
			//              list response
			return list, nil
		}
	}

	// wrap inner response with default protocol
	var prepareProtocol endpoint.Middleware = func (inner endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			v, err := inner(ctx, request)
			if err != nil {
				return
			}

			response = store.NewResponse(nounp, v)
			return

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
		request = allocEntity()

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

	//
	// -----
	//

	// define paths
	nounPath := base + "/" + noun + "/{id}"
	pluralNounPath := base + "/" + nounp
	log.Printf("REST path: %s", pluralNounPath)

	endpoints := {{ .Store }}Endpoints(noun, nounp)

	// Create
	postHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p),
			prepareProtocol,
			prepareCreate,
			genRequestPermChecker("create "+noun))(endpoints["create"]),
		decodeJSONReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("POST", pluralNounPath, postHandler)

	// Retrieve single
	getHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p),
			prepareProtocol,
			prepareList,
			genResponsePermChecker("retrieve "+noun))(endpoints["retrieve"]),
		decodeIDReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("GET", nounPath, getHandler)

	// Retrieve list
	listHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p),
			prepareProtocol,
			prepareList,
			genResponsePermChecker("list "+noun))(endpoints["list"]),
		decodeListReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("GET", pluralNounPath, listHandler)

	// Update single
	putHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p),
			prepareProtocol,
			prepareUpdate,
			genRequestPermChecker("update "+noun))(endpoints["update"]),
		decodeUpdate,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("PUT", nounPath, putHandler)

	// Delete single
	deleteHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p),
			prepareProtocol,
			genRequestPermChecker("delete "+noun))(endpoints["delete"]),
		decodeIDReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("DELETE", nounPath, deleteHandler)

}

{{ end }}
