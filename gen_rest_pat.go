package main

func init() {
	tpls.Append("gen rest:gorilla/pat",
		`{{ define "package" }}package {{ .Pkg }}{{ end }}`,
		`{{ define "imports" }}
	"github.com/gorilla/pat"
	"github.com/gourd/codec"
	"github.com/gourd/service"
	"log"
	"net/http"
{{ end }}`,
		`{{ define "code" }}


// dummy rest binder
func {{ .Type.Name }}Rest(r *pat.Router, base, noun, nounp string) {

	// define paths
	p := base + "/" + nounp
	ep := base + "/" + noun + "/{id}"

	// way to identify the key condition (for update, retrieve and delete)
	getKeyCond := func(r *http.Request) (cond service.Conds) {
		// TODO: generate the content of this function
		//       dynamically with gourd
		id := r.URL.Query().Get(":id") // will change
		return service.NewConds().Add("id", id)
	}

	// handle permission related issue
	scopeAllow := func(r *http.Request, respEnc codec.Encoder, act string, info ...interface{}) bool {
		s, ok := GetScopesOk(r)
		if !ok {
			log.Printf("Error loading scopes of current session")
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusInternalServerError,
				"message": "Server error. Please try again later",
			})
			return false
		} else if err := s.Allow(act, info...); err != nil {
			log.Printf("Permission denied")
			respEnc.Encode(map[string]interface{}{
				"status":  "error",
				"code":    http.StatusForbidden,
				"message": err.Error(),
			})
			return false
		}
		return true
	}

	log.Printf("REST path: %s", p)

	// Create
	r.Post(p, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// get service
		sr, err := Get{{ .Type.Name }}(r)
		s, _ := sr.(*{{ .Type.Name }})
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			return
		}

		// allocate entity
		e := s.AllocEntity()

		// assign encoder and decoder
		reqtDec, _ := codec.GetDecoderOk(r)
		respEnc, _ := codec.GetEncoderOk(w, r)

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

		if !scopeAllow(r, respEnc, "create "+noun, e) {
			// test access scope fail, do nothing
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

		// get service
		sr, err := Get{{ .Type.Name }}(r)
		s, _ := sr.(*{{ .Type.Name }})
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			return
		}

		// allocate memory for variables
		el := s.AllocEntityList()

		// assign encoder and decoder
		respEnc, _ := codec.GetEncoderOk(w, r)

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
		} else if !scopeAllow(r, respEnc, "load "+noun, el) {
			// test access scope fail, do nothing
			return
		} else {
			respEnc.Encode(map[string]interface{}{
				"status": "success",
				"code":   http.StatusOK,
				nounp:    el,
			})
		}
	})

	// TODO: Retrieve list

	// Update
	r.Put(ep, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// get service
		sr, err := Get{{ .Type.Name }}(r)
		s, _ := sr.(*{{ .Type.Name }})
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			return
		}

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// assign encoder and decoder
		reqtDec, _ := codec.GetDecoderOk(r)
		respEnc, _ := codec.GetEncoderOk(w, r)

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

		if !scopeAllow(r, respEnc, "update "+noun, el) {
			// test access scope fail, do nothing
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

		// get service
		sr, err := Get{{ .Type.Name }}(r)
		s, _ := sr.(*{{ .Type.Name }})
		if err != nil {
			log.Printf("Error obtaining {{ .Type.Name }} service: %s", err.Error())
			return
		}

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// assign encoder and decoder
		respEnc, _ := codec.GetEncoderOk(w, r)

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

		if !scopeAllow(r, respEnc, "delete "+noun, el) {
			// test access scope fail, do nothing
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
