package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/pat"
	"github.com/gourd/codec"
	"github.com/gourd/perm"
	"github.com/gourd/service/upperio"
	"log"
	"net/http"
	"time"
	driver "upper.io/db/postgresql"
)

// getServer
//
// drafted gourd generated server
// will be generated by gourd CLI tool eventually
//
// Should also be a http.Handler that
// cound be tested by wrapping httptest sevrer
func gourdServer() (h http.Handler) {

	// define db
	upperio.Define("default", driver.Adapter, driver.ConnectionURL{
		Database: `some-nonsense`,
	})

	// create router specific / independent middleware
	ch := &codec.Handler{}

	// approve all permission test
	p := perm.NewMux()
	requireAccess := func(r *http.Request, perm string, info ...interface{}) error {
		log.Printf("Requesting permission to %s", perm)
		return nil
	}
	p.HandleFunc("create post", requireAccess)
	p.HandleFunc("load post", requireAccess)
	p.HandleFunc("list post", requireAccess)
	p.HandleFunc("update post", requireAccess)
	p.HandleFunc("delete post", requireAccess)

	// create router of the specific type
	rtr := pat.New()

	// add services rest to router
	PostServiceRest(rtr, "/api", "post", "posts")
	CommentServiceRest(rtr, "/api", "comment", "comments")

	// add login form to router
	// TODO: need a way to inject templates for login form
	// NOTE: this will be generated if needed to be router specific
	//ah.ServeLogin(rtr, "/login")

	// create negroni middleware handler
	// with middlewares
	n := negroni.New()
	n.Use(negroni.Wrap(ch))
	n.Use(negroni.Wrap(p))

	// use router in negroni
	n.UseHandler(rtr)

	return n
}

// gourdMain
//
// gourd generated main
// feel free to copy the code and change
func gourdMain() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        gourdServer(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}