package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
	"github.com/gourd/codec"
	"upper.io/db"
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
	db, err := db.Open(sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./data/sqlite3.db`,
	})
	if err != nil {
		panic(err)
	}

	// create router specific / independent middleware
	ch := &codec.Handler{}
	as := &OAuth2Storage{}

	// provide services to auth provider
	//as.UseServices(ClientService{db}, AuthService{db}, AccessService{db})

	// create router of the specific type
	r := pat.New()

	// add services rest to router
	ps := &PostService{db}
	ps.Rest(r, "/api", "post", "posts")

	cs := &CommentService{db}
	cs.Rest(r, "/api", "comment", "comments")

	// add oauth2 endpoints
	//AddOAuth2(r, "/oauth", ap)

	// create negroni middleware handler
	// with middlewares
	n = negroni.New()
	n.Use(negroni.Wrap(ch))
	n.Use(negroni.Wrap(as.ServeScopes()))

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
