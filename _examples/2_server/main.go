//go:generate gourd gen main $GOFILE
package main

import (
	"github.com/gourd/kit/store/upperio"

	"net/http"
	"os"
	"upper.io/db/sqlite"
)

// port to use
var port string

func init() {
	port = "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
}

// main
//
//gourd:db upperio
//gourd:router gorillamux
//gourd:middleware negroni
//gourd:auth /oauth2 osin
//gourd:rest /api/posts Post
//gourd:rest /api/comments Comment
//
func main() {

	// define db
	upperio.Define("default", sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./data/sqlite3.db`,
	})

	// use the MainHandler (should be gourd generated)
	// alternatively, user may copy the content of MainHandler
	// here then modify
	http.Handle("/", MainHandler())
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
