package main

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/pat"
	gourdctx "github.com/gourd/kit/context"
	"github.com/gourd/kit/oauth2"
	"github.com/gourd/kit/perm"
	httpservice "github.com/gourd/kit/service/http"
	"golang.org/x/net/context"
)

// MainHandler
//
// drafted gourd generated server
// will be generated by gourd CLI tool eventually
//
// Should also be a http.Handler that
// cound be tested by wrapping httptest sevrer
func MainHandler() http.Handler {

	// oauth2 manager with default settings
	m := oauth2.NewManager()

	// provide access handlers to
	// NOTE: these are independent to router
	p := perm.NewMux()
	requireAccess := func(ctx context.Context, perm string, info ...interface{}) (err error) {
		// get access from context HTTP Request
		log.Printf("Requesting permission to %s", perm)

		if access := oauth2.GetAccess(ctx); access != nil {
			log.Printf("Access: %#v", access)
		}
		return
	}
	p.HandleFunc("create post", requireAccess)
	p.HandleFunc("retrieve post", requireAccess)
	p.HandleFunc("list post", requireAccess)
	p.HandleFunc("update post", requireAccess)
	p.HandleFunc("delete post", requireAccess)

	// create router of the specific type
	rtr := pat.New()

	// handle individual route with pat (router specific)
	rtfn := func(r *pat.Router) httpservice.RouterFunc {
		return func(path string, methods []string, h http.Handler) error {
			for i := range methods {
				r.Add(methods[i], path, h)
			}
			return nil
		}
	}(rtr)

	// generates noun paths for different API endpoints
	genPath := func(base string) func(noun, nounp string) httpservice.Paths {
		return func(noun, nounp string) httpservice.Paths {
			return httpservice.NewPaths(base,
				httpservice.NewNoun(noun, nounp), "{id}")
		}
	}("/api")

	// common patch to all services
	common := func(services httpservice.Services) httpservice.Services {
		// common middleware
		commonMware := endpoint.Chain(
			gourdctx.ClearGorilla,
			perm.UseMux(p))

		// route all endpoints
		for name := range services {
			// pre-wrap endpoint with common middleware
			services[name].Middlewares.Add(
				httpservice.MWOuter, commonMware)
		}
		return services
	}

	// add store rest to router
	PostStoreRest(rtfn, genPath("post", "posts"), common)
	CommentStoreRest(rtfn, genPath("comment", "comments"), common)
	oauth2.RoutePat(rtr, "/oauth", m.GetEndpoints())

	// add login form to router
	// TODO: need a way to inject templates for login form
	// NOTE: this will be generated if needed to be router specific
	//ah.ServeLogin(rtr, "/login")

	// create negroni middleware handler
	// with middlewares
	n := negroni.New()
	n.Use(negroni.Wrap(m.Middleware()))

	// use router in negroni
	n.UseHandler(rtr)

	return n
}
