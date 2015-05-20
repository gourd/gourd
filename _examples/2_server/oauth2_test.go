package main

import (
	"log"
	"net/http"
	"testing"
)

func TestOAuth2(t *testing.T) {
	ts := gourdTestServer()
	defer ts.Close()

	// form request
	req, err := http.NewRequest("GET", ts.URL+"/oauth/authorize"+
		"?response_type=code&client_id=testing&"+
		"redirect_url=http://localhost:8080/appauth/code", nil)
	if err != nil {
		t.Errorf("Failed to form new request: %s", err.Error())
	}

	// new client to do the request
	c := &http.Client{}
	r, err := c.Do(req)
	if err != nil {
		t.Errorf("Failed run the request: %s", err.Error())
	}

	log.Printf("Response: %#v", r)
}
