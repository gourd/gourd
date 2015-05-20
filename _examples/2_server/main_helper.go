package main

import (
	"github.com/RangelReale/osin"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
	"github.com/gourd/codec"
	"github.com/gourd/service"
	"github.com/gourd/service/upperio"
	"upper.io/db/sqlite"
)

// getServer
//
// drafted gourd generated server
// will be generated by gourd CLI tool eventually
//
// Should also be a http.Handler that
// cound be tested by wrapping httptest sevrer
func gourdServer() (n *negroni.Negroni) {

	// define db
	upperio.Define("default", sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./data/sqlite3.db`,
	})

	// create router specific / independent middleware
	ch := &codec.Handler{}
	ah := &OAuth2Handler{} // will be handled by another gourd repo

	// provide services to auth storage
	// NOTE: these are independent to router
	as := &OAuth2Storage{}
	as.UseClientFrom(service.Providers.MustGet("Client"))
	as.UseAuthFrom(service.Providers.MustGet("AuthorizeData"))
	as.UseAccessFrom(service.Providers.MustGet("AccessData"))
	as.UseUserFrom(service.Providers.MustGet("User"))
	ah.UseStorage(as)

	// provide storage to osin server
	// provide osin server to OAuth2Handler
	cfg := osin.NewServerConfig()
	cfg.AllowGetAccessRequest = true
	cfg.AllowClientSecretInParams = true
	cfg.AllowedAccessTypes = osin.AllowedAccessType{
		osin.AUTHORIZATION_CODE,
		osin.REFRESH_TOKEN,
	}
	cfg.AllowedAuthorizeTypes = osin.AllowedAuthorizeType{
		osin.CODE,
		osin.TOKEN,
	}
	ah.InitOsin(cfg)

	// provide access handlers to
	// NOTE: these are independent to router
	//ah.Handle("edit post", someAccessHandler)
	//ah.HandleFunc("edit comment", someAccessHandleFunc)

	// create router of the specific type
	r := pat.New()

	// add services rest to router
	PostServiceRest(r, "/api", "post", "posts")
	CommentServiceRest(r, "/api", "comment", "comments")

	// add oauth2 endpoints to router
	// NOTE: this will be generated if needed to be router specific
	//ah.ServeEndpoints(r, "/oauth")

	// add login form to router
	// TODO: need a way to inject templates for login form
	// NOTE: this will be generated if needed to be router specific
	//ah.ServeLogin(r, "/login")

	// create negroni middleware handler
	// with middlewares
	n = negroni.New()
	n.Use(negroni.Wrap(ch))
	n.Use(negroni.Wrap(ah.ServeScopes()))

	// use router in negroni
	n.UseHandler(r)

	return
}

// gourdMain
//
// gourd generated main
// feel free to copy the code and change
func gourdMain() {
	s := gourdServer()
	if s != nil {
		s.Run(":8080") // negroni specific
	}
}
