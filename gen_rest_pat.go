package main

func init() {
	tpls.Append("gen rest:gorilla/pat",
		`{{ define "package" }}package {{ .Pkg }}{{ end }}`,
		`{{ define "imports" }}
	"github.com/gorilla/pat"
	"github.com/gourd/service"
	"log"
	"net/http"
{{ end }}`,
		`{{ define "code" }}


// dummy rest binder
func (s *{{ .Type.Name }}) Rest(r *pat.Router, base, noun, nounp string) {

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

	log.Printf("REST path: %s", p)

	// Create
	r.Post(p, func(w http.ResponseWriter, r *http.Request) {
		var err error

		// allocate entity
		e := s.AllocEntity()

		// assign encoder and decoder
		reqtDec := GetReqtDec(r)
		respEnc := GetRespEnc(w, r)

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

		// allocate memory for variables
		el := s.AllocEntityList()

		// assign encoder and decoder
		respEnc := GetRespEnc(w, r)

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

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// assign encoder and decoder
		reqtDec := GetReqtDec(r)
		respEnc := GetRespEnc(w, r)

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

		// allocate memory for variables
		e := s.AllocEntity()
		el := s.AllocEntityList()

		// assign encoder and decoder
		respEnc := GetRespEnc(w, r)

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
