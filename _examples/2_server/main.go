//go:generate gourd gen main $GOFILE
package main

import (
	"net/http"
	"os"
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

	// use the MainHandler (should be gourd generated)
	// alternatively, user may copy the content of MainHandler
	// here then modify
	http.Handle("/", NewHandler(NewFactory(`./data/sqlite3.db`)))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
