package main

import (
	"log"
	"net/http"
	"time"
)

//go:generate gourd gen main $GOFILE

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
	s := &http.Server{
		Addr:           ":8080",
		Handler:        NewHandler(NewFactory()),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
