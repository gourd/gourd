{{ define "package" }}
	package {{ .Pkg }}
{{ end }}

{{ define "imports" }}
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/pat"
	gourdctx "github.com/gourd/context"
	"github.com/gourd/perm"
	"github.com/gourd/service"
	"golang.org/x/net/context"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
{{ end }}

{{ define "code" }}

// {{ .Service.Name }}Endpoints generate endpoints for CURD
func {{ .Service.Name }}Endpoints(noun, nounp string) (endpoints map[string]endpoint.Endpoint) {

	// variables to use later
	blankService := &{{ .Service.Name }}{}
	allocEntity := blankService.AllocEntity
	allocEntityList := blankService.AllocEntityList
	getService := Get{{ .Service.Name }}
	serviceName := "{{ .Service.Name }}"

	// enforce entity property before create
	prepareCreate := func(r *http.Request, e service.EntityPtr) (err error) {
		// placeholder: anything you want to do with the entity
		//              before append to database
		return
	}

	// enforce entity property before update
	prepareUpdate := func(e service.EntityPtr, el service.EntityListPtr) (err error) {
		// placeholder: anything you want to do with the entity
		//              before update to database
		return
	}

	// prepare entity list
	prepareEntityList := func(r *http.Request, el service.EntityListPtr) (err error) {
		// placeholder: anything you want to do with the entity
		//              before update to database
		return
	}

	// handle permission related issue
	permAllow := func(r *http.Request, permission string, info ...interface{}) error {
		m := perm.GetMux(r)
		return m.Allow(r, permission, info...)
	}

	// store endpoints here
	// TODO: may have new struct to store
	endpoints = make(map[string]endpoint.Endpoint)

	endpoints["Create"] = func(ctx context.Context, e interface{}) (res interface{}, err error) {

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := service.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// enforce create integrity
		prepareCreate(r, e)

		// get service
		s, err := getService(r)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s service (%s)", serviceName, err)
			err = serr
			return
		}
		defer s.Close()

		// check permission
		if err = permAllow(r, "create "+noun, e); err != nil {
			err = service.ErrorForbidden
			return
		}

		// create entity
		log.Printf("[!!!] entity to create: %#v", e)
		err = s.Create(nil, e)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error creating %s: %#v, entity: %#v", noun, err.Error(), e)
			err = serr
			return
		}

		// encode response
		res = service.NewResponse(nounp, []interface{}{e})
		return

	}

	endpoints["Retrieve"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		q := request.(service.Query)

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := service.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get service
		s, err := getService(r)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s service (%s)", serviceName, err)
			err = serr
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := allocEntityList()

		// retrieve
		err = s.Search(q).All(el)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// encode response
		if s.Len(el) == 0 {
			err = service.ErrorNotFound
			return
		} else if err = permAllow(r, "load "+noun, el); err != nil {
			serr := service.ErrorForbidden
			serr.ServerMsg = err.Error()
			err = serr
			return
		} else if err = prepareEntityList(r, el); err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		}

		res = service.NewResponse(nounp, el)
		return

	}

	endpoints["List"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		q := request.(service.Query)

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := service.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get service
		s, err := getService(r)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s service (%s)", serviceName, err)
			err = serr
			return
		}
		defer s.Close()

		// allocate memory for variables
		var el interface{} = allocEntityList()

		err = s.Search(q).All(el)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// encode response
		if err = permAllow(r, "list "+noun, el); err != nil {
			serr := service.ErrorForbidden
			serr.ServerMsg = err.Error()
			err = serr
			return
		} else if err = prepareEntityList(r, el); err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		} else if s.Len(el) == 0 {
			el = &[]int{}
		}

		res = service.NewResponse(nounp, el)
		return

	}

	endpoints["Update"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		// allocate memory for variables
		var el interface{} = allocEntityList()

		rmap := request.(map[string]interface{})
		q := rmap["query"].(service.Query)
		e := rmap["entity"]
		cond := q.GetConds()

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := service.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get service
		s, err := getService(r)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s service (%s)", serviceName, err)
			err = serr
			return
		}
		defer s.Close()

		// find the content of the id
		err = s.Search(q).All(el)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// enforce update integrity
		err = prepareUpdate(e, el)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error Preparing with PrepDb %s: %s", noun, err)
			err = serr
			return
		}

		// test permission
		if err = permAllow(r, "update "+noun, el); err != nil {
			serr := service.ErrorForbidden
			serr.ServerMsg = err.Error()
			err = serr
			return
		}

		// update entity
		if err = s.Update(cond, e); err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		}

		res = service.NewResponse(nounp, []interface{}{e})
		return
	}

	endpoints["Delete"] = func(ctx context.Context, request interface{}) (res interface{}, err error) {

		// allocate memory for variables
		e := allocEntity()
		el := allocEntityList()

		// service query of id
		q := request.(service.Query)
		cond := q.GetConds()

		// get context information
		r, ok := gourdctx.HTTPRequest(ctx)
		if !ok {
			serr := service.ErrorInternal
			serr.ServerMsg = "missing request in context"
			err = serr
			return
		}

		// get service
		s, err := getService(r)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error obtaining %s service (%s)", serviceName, err)
			err = serr
			return
		}
		defer s.Close()

		// find the content of the id
		err = s.Search(q).All(el)
		if err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf("error searching %s: %s",
				noun, err)
			err = serr
			return
		}

		// test permission
		if err = permAllow(r, "delete "+noun, el); err != nil {
			serr := service.ErrorForbidden
			serr.ServerMsg = err.Error()
			err = serr
			return
		}

		// delete entity
		if err = s.Delete(cond); err != nil {
			serr := service.ErrorInternal
			serr.ServerMsg = fmt.Sprintf(
				"error encoding %s list: %s",
				noun, err)
			err = serr
			return
		}

		res = service.NewResponse(nounp, []interface{}{e})
		return

	}

	return	
}

// {{ .Service.Name }}Rest binds service to pat router
func {{ .Service.Name }}Rest(r *pat.Router, base, noun, nounp string) {

	// variables to use later
	blankService := &{{ .Service.Name }}{}
	allocEntity := blankService.AllocEntity

	// decodeIDReq generically decoded :id field
	// (works with pat based URL routing)
	var decodeIDReq httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {
		id := r.URL.Query().Get(":id") // will change
		cond := service.NewConds().Add("id", id)
		request = service.NewQuery().SetConds(cond)
		return
	}

	// decodeListReq decode query for list endpoint
	var decodeListReq httptransport.DecodeRequestFunc = func(r *http.Request) (request interface{}, err error) {

		q := service.NewQuery()

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

		request = rmap
		return
	}

	// encodeJSONResp encodes given response into JSON
	encodeJSONResp := func(w http.ResponseWriter, response interface{}) (err error) {
		enc := json.NewEncoder(w)
		err = enc.Encode(response)
		return
	}

	// JSONErrorEncoder expands given error to ServiceError then encode to JSON
	JSONErrorEncoder := func(w http.ResponseWriter, err error) {

		// quick fix for gokit bad request wrapping problem
		switch err.(type) {
		case httptransport.BadRequestError:
			err = err.(httptransport.BadRequestError).Err
		}

		enc := json.NewEncoder(w)
		err = enc.Encode(service.ExpandError(err))
	}

	//
	// -----
	//

	// define paths
	nounPath := base + "/" + noun + "/{id}"
	pluralNounPath := base + "/" + nounp
	log.Printf("REST path: %s", pluralNounPath)

	endpoints := {{ .Service.Name }}Endpoints(noun, nounp)

	// define middleware chain for all endpoints
	mwares := endpoint.Chain(
		gourdctx.ClearGorilla)

	// Create
	postHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		mwares(endpoints["Create"]),
		decodeJSONReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("POST", pluralNounPath, postHandler)

	// Retrieve single
	getHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		mwares(endpoints["Retrieve"]),
		decodeIDReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("GET", nounPath, getHandler)

	// Retrieve list
	listHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		mwares(endpoints["List"]),
		decodeListReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("GET", pluralNounPath, listHandler)

	// Update single
	putHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		mwares(endpoints["Update"]),
		decodeUpdate,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("PUT", nounPath, putHandler)

	// Delete single
	deleteHandler := httptransport.NewServer(
		gourdctx.NewEmpty(),
		mwares(endpoints["Delete"]),
		decodeIDReq,
		encodeJSONResp,
		httptransport.ServerBefore(gourdctx.UseGorilla),
		httptransport.ServerErrorEncoder(JSONErrorEncoder))
	r.Add("DELETE", nounPath, deleteHandler)

}

{{ end }}
