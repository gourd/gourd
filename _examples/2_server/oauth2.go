package main

import (
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/gourd/service"
	"log"
	"net/http"
	"net/url"
	"text/template"
)

// OAuth2Endpoints contains http handler func of different endpoints
type OAuth2Endpoints struct {
	Auth  http.HandlerFunc
	Token http.HandlerFunc
}

// OAuth2Handler handles oauth2 related request
// Also provide middleware for other http handler function
// to access scope related information
type OAuth2Handler struct {
	Storage    *OAuth2Storage
	OsinServer *osin.Server
}

// UseOsin set the OsinServer
func (h *OAuth2Handler) InitOsin(cfg *osin.ServerConfig) *OAuth2Handler {
	h.OsinServer = osin.NewServer(cfg, h.Storage)
	return h
}

// Storage provides a osin storage interface
func (h *OAuth2Handler) UseStorage(s *OAuth2Storage) *OAuth2Handler {
	h.Storage = s
	return h
}

// ServeScopes provides a scope handler middleware
func (h *OAuth2Handler) ServeScopes() *ScopesHandler {
	return &ScopesHandler{}
}

// GetEndpoints generate endpoints http handers and return
func (h *OAuth2Handler) GetEndpoints() *OAuth2Endpoints {

	// read login credential
	getLoginCred := func(r *http.Request) (idField, id, password string) {
		idField = "username"
		id = r.Form.Get(idField)
		password = r.Form.Get("password")
		return
	}

	// template for login form
	tmplStr := `
<!DOCTYPE html>
<html>
<body>
	LOGIN {{ .SiteName }}<br/>
	<form action="{{ .FormAction }}" method="POST">
		Login: <input type="text" name="login" /><br/>
		Password: <input type="password" name="password" /><br/>
		<input type="submit"/>
	</form>
</body>
</html>
`

	tmpl, err := template.New("loginForm").Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	// handle login
	handleLogin := func(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) (err error) {

		log.Printf("handleLogin")

		// parse POST input
		r.ParseForm()
		if r.Method == "POST" {

			// get login information from form
			idField, id, password := getLoginCred(r)
			log.Printf("login: %s, %s, %s", idField, id, password)
			if id == "" || password == "" {
				err = fmt.Errorf("Empty Username or Password")
				return
			}

			// obtain user service
			var us service.Service
			us, err = h.Storage.UserService(r)
			if err != nil {
				log.Printf("Error obtaining user service: %s", err.Error())
				err = fmt.Errorf("Internal Server Error")
				return
			}

			// get user from database
			u := us.AllocEntity()
			c := service.NewConds().Add(idField, id)
			err = us.One(c, u)
			if err != nil {
				log.Printf("Error searching user \"%s\": %s", id, err.Error())
				err = fmt.Errorf("Internal Server Error")
				return
			}

			// if user does not exists
			if u == nil {
				log.Printf("Unknown user \"%s\" attempt to login", id)
				err = fmt.Errorf("Username or Password incorrect")
				return
			}

			// cast the user as OAuth2User
			// and do password check
			ou, ok := u.(OAuth2User)
			if !ok {
				log.Printf("User cannot be cast as OAuth2User")
				err = fmt.Errorf("Internal server error")
				return
			}

			// if password does not match
			if !ou.PasswordIs(password) {
				log.Printf("Attempt to login \"%s\" with incorrect password", id)
				err = fmt.Errorf("Username or Password incorrect")
			} else {
				log.Printf("Login \"%s\" success", id)
			}

			// return pointer of user object, allow it to be re-cast
			ar.UserData = u
			return
		}

		// no POST input or incorrect login, show form

		// build action query
		aq := url.Values{}
		aq.Add("response_type", string(ar.Type))
		aq.Add("client_id", ar.Client.GetId())
		aq.Add("state", ar.State)
		aq.Add("scope", ar.Scope)
		aq.Add("redirect_uri", ar.RedirectUri)

		// template variables
		vars := map[string]interface{}{
			"SiteName":   "Gourd: Example 2",
			"FormAction": r.URL.Path + "?" + aq.Encode(),
		}

		// render the form with vars
		err = tmpl.Execute(w, vars)
		if err != nil {
			panic(err)
		}
		return
	}

	ep := OAuth2Endpoints{}

	// authorize endpoint
	ep.Auth = func(w http.ResponseWriter, r *http.Request) {

		log.Printf("auth endpoint")

		srvr := h.OsinServer
		resp := srvr.NewResponse()
		resp.Storage.(*OAuth2Storage).SetRequest(r)

		// handle authorize request with osin
		if ar := srvr.HandleAuthorizeRequest(resp, r); ar != nil {
			log.Printf("handle authorize request")
			if err := handleLogin(ar, w, r); err != nil {
				return
			}
			log.Printf("OAuth2 Authorize Request: User obtained: %#v", ar.UserData)
			ar.Authorized = true
			srvr.FinishAuthorizeRequest(resp, r, ar)
		}
		if resp.InternalError != nil {
			log.Printf("Internal Error: %s", resp.InternalError.Error())
		}
		log.Printf("OAuth2 Authorize Response: %#v", resp)
		osin.OutputJSON(resp, w, r)

	}

	// token endpoint
	ep.Token = func(w http.ResponseWriter, r *http.Request) {

		srvr := h.OsinServer
		resp := srvr.NewResponse()
		resp.Storage.(*OAuth2Storage).SetRequest(r)

		if ar := srvr.HandleAccessRequest(resp, r); ar != nil {
			// TODO: handle authorization
			// check if the user has the permission to grant the scope
			log.Printf("Access successful")
			ar.Authorized = true
			srvr.FinishAccessRequest(resp, r, ar)
		} else if resp.InternalError != nil {
			log.Printf("Internal Error: %s", resp.InternalError.Error())
		}
		log.Printf("OAuth2 Token Response: %#v", resp)
		osin.OutputJSON(resp, w, r)

	}

	return &ep

}
