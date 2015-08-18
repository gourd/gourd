package main

func init() {
	tpls.Append("gen rest:gorilla/pat",
		`{{ define "package" }}package {{ .Pkg }}{{ end }}`,
		`{{ define "imports" }}
	"github.com/gorilla/pat"
	"github.com/gourd/codec"
	"github.com/gourd/perm"
	"github.com/gourd/service"
	"log"
	"net/http"
	"strconv"
{{ end }}`,
		`{{ define "code" }}


// {{ .Type.Name }}Rest binds service to pat router
func {{ .Type.Name }}Rest(r *pat.Router, base, noun, nounp string) {

	// define paths
	p := base + "/" + nounp
	ep := base + "/" + noun + "/{id}"
	epp := base + "/" + nounp

	// way to identify the key condition (for update, retrieve and delete)
	getKeyCond := func(r *http.Request) (cond service.Conds) {
		// TODO: generate the content of this function
		//       dynamically with gourd
		id := r.URL.Query().Get(":id") // will change
		return service.NewConds().Add("id", id)
	}

	// handle permission related issue
	permAllow := func(r *http.Request, respEnc codec.Encoder, permission string, info ...interface{}) error {
		m := perm.GetMux(r)
		return m.Allow(r, permission, info...)
	}

	log.Printf("REST path: %s", p)

	// Create
	r.Post(p, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// assign encoder and decoder
		reqtDec, _ := codec.GetDecoderOk(r)
		respEnc, _ := codec.GetEncoderOk(w, r)

		// get service
		s, err := Get{{ .Type.Name }}(r)
		if err != nil {
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			return
		}
		defer s.Close()

		// allocate entity
		e := s.AllocEntity()

		// decode request
		err = reqtDec.Decode(e)
		if err != nil {
			log.Printf("Error JSON Unmarshal: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Cannot decode request",
			})
			return
		}

		// check permission
		if err = permAllow(r, respEnc, "create "+noun, e); err != nil {
			code, msg := service.ParseError(err)
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    code,
				"message": msg,
			})
			return
		}

		// create entity
		err = s.Create(nil, e)
		if err != nil {
			log.Printf("Error Creating %s: %s", noun, err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Failed to create entity",
			})
			return
		}

		// encode response
		respEnc.Encode(map[string]interface{}{
			"status": "success",
			"code":   http.StatusOK,
			nounp:    []interface{}{e},
		})
	})

	// Retrieve single
	r.Get(ep, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// assign encoder and decoder
		respEnc, _ := codec.GetEncoderOk(w, r)

		// get service
		s, err := Get{{ .Type.Name }}(r)
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := s.AllocEntityList()

		// retrieve
		cond := getKeyCond(r)
		err = s.Search(cond, el)
		if err != nil {
			log.Printf("Error searching %s: %s", noun, err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Failed to find entity",
			})
			return
		}

		// encode response
		if s.Len(el) == 0 {
			respEnc.Encode(map[string]interface{}{
				"status": "error",
				"code":   http.StatusNotFound,
			})
		} else if err = permAllow(r, respEnc, "load "+noun, el); err != nil {
			code, msg := service.ParseError(err)
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    code,
				"message": msg,
			})
			return
		} else {
			respEnc.Encode(map[string]interface{}{
				"status": "success",
				"code":   http.StatusOK,
				nounp:    el,
			})
		}
	})

	// Retrieve list
	r.Get(epp, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// assign encoder and decoder
		respEnc, _ := codec.GetEncoderOk(w, r)

		// get service
		s, err := Get{{ .Type.Name }}(r)
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		defer s.Close()

		// allocate memory for variables
		el := s.AllocEntityList()

		// parse paging request parameter
		offset, limit := func (r *http.Request) (o, l uint64) {
			ostr := r.Form.Get("offset")
			lstr := r.Form.Get("limit")
			if ostr == "" {
				if ot, err := strconv.ParseUint(ostr, 10, 64); err == nil {
					o = ot
				}
			}
			if lstr == "" {
				if lt, err := strconv.ParseUint(ostr, 10, 64); err == nil {
					l = lt
				}
			}
			return
		}(r)

		// retrieve
		cond := service.NewConds()
		cond.SetOffset(offset)
		cond.SetLimit(limit)
		log.Printf("Retrieve list!!!! %#v", cond)

		err = s.Search(cond, el)
		if err != nil {
			log.Printf("Error searching %s: %s", noun, err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Failed to find entity",
			})
			return
		}

		// encode response
		if err = permAllow(r, respEnc, "list "+noun, el); err != nil {
			code, msg := service.ParseError(err)
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    code,
				"message": msg,
			})
			return
		} else if s.Len(el) == 0 {
			respEnc.Encode(map[string]interface{}{
				"status": "success",
				"code":   http.StatusOK,
				nounp:    []int{},
			})
		} else {
			respEnc.Encode(map[string]interface{}{
				"status": "success",
				"code":   http.StatusOK,
				nounp:    el,
			})
		}
	})

	// Update
	r.Put(ep, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// assign encoder and decoder
		reqtDec, _ := codec.GetDecoderOk(r)
		respEnc, _ := codec.GetEncoderOk(w, r)

		// get service
		s, err := Get{{ .Type.Name }}(r)
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		defer s.Close()

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// decode request
		err = reqtDec.Decode(e)
		if err != nil {
			log.Printf("Error JSON Unmarshal: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Cannot decode request",
			})
			return
		}

		// find the content of the id
		cond := getKeyCond(r)
		err = s.Search(cond, el)
		if err != nil {
			log.Printf("Error searching %s: %s", noun, err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Failed to find entity",
			})
			return
		}

		// test permission
		if err = permAllow(r, respEnc, "update "+noun, el); err != nil {
			code, msg := service.ParseError(err)
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    code,
				"message": msg,
			})
			return
		}

		// update entity
		s.Update(cond, e)

		// encode response
		respEnc.Encode(map[string]interface{}{
			"status": "success",
			"code":   http.StatusOK,
			noun:     e,
		})
	})

	// Delete
	r.Delete(ep, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// assign encoder and decoder
		respEnc, _ := codec.GetEncoderOk(w, r)

		// get service
		s, err := Get{{ .Type.Name }}(r)
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Internal Server Error",
			})
			return
		}
		defer s.Close()

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// find the content of the id
		cond := getKeyCond(r)
		err = s.Search(cond, el)
		if err != nil {
			log.Printf("Error searching %s: %s", noun, err.Error())
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusBadRequest,
				"message": "Failed to find entity",
			})
			return
		}

		// test permission
		if err = permAllow(r, respEnc, "delete "+noun, el); err != nil {
			code, msg := service.ParseError(err)
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    code,
				"message": msg,
			})
			return
		}

		// delete entity
		s.Delete(cond)

		// encode response
		respEnc.Encode(map[string]interface{}{
			"status": "success",
			"code":   http.StatusOK,
			noun:     e,
		})
	})

}


{{ end }}`)

	tpls.AddDeps("gen rest:gorilla/pat", "gen:general")
}
