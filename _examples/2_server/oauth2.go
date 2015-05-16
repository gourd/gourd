package main

import (
	"github.com/gorilla/context"
	"net/http"
)

// OAuth2Storage handles storage
type OAuth2Storage struct {
}

func (o OAuth2Storage) ServeScopes() *ScopesHandler {
	return &ScopesHandler{}
}

// ScopesHandler provides scopes to handlers through gorilla context
// TODO: some how accept a function to handle ACL test
type ScopesHandler struct {
}

func (sh ScopesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: provide scope according to database and session information
	context.Set(r, "scopes", &BasicScopes{})
}

// Scopes is the abstract interface to access
// current OAuth2 scopes. Allow http handler to
// check permission of current session
type Scopes interface {
	Allow(act string, info ...interface{}) error
}

// BasicScopes is the primary implementation
// of Scope interface
type BasicScopes struct {
	Scopes []string
}

// Allow check
func (s BasicScopes) Allow(act string, info ...interface{}) error {
	return nil
}

// GetScopesOk retrieves scope from gorilla context
func GetScopesOk(r *http.Request) (Scopes, bool) {
	if val, ok := context.GetOk(r, "scopes"); ok {
		if s, ok := val.(Scopes); ok {
			return s, ok
		}
	}
	return nil, false
}
