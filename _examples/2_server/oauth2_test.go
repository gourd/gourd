package main

import (
	"github.com/gourd/service"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestOAuth2(t *testing.T) {
	ts := gourdTestServer()
	defer ts.Close()

	password := "password"

	c, u := func() (*Client, *User) {
		r := &http.Request{}

		// generate dummy user
		us, err := service.Providers.MustGet("User")(r)
		if err != nil {
			panic(err)
		}
		u := dummyNewUser(password)
		err = us.Create(service.NewConds(), u)
		if err != nil {
			panic(err)
		}

		// get related dummy client
		cs, err := service.Providers.MustGet("Client")(r)
		if err != nil {
			panic(err)
		}
		c := dummyNewClient(ts.URL + "/appauth/code")
		c.UserId = u.Id
		err = cs.Create(service.NewConds(), c)
		if err != nil {
			panic(err)
		}

		return c, u
	}()

	// build the login form
	form := url.Values{}
	form.Add("username", u.Username)
	form.Add("password", password)
	log.Printf("form send: %s", form.Encode())

	// build the query string
	q := &url.Values{}
	q.Add("response_type", "code")
	q.Add("client_id", c.StrId)
	q.Add("redirect_url", c.RedirectUri)

	// build request
	req, err := http.NewRequest("POST",
		ts.URL+"/oauth/authorize"+"?"+q.Encode(),
		strings.NewReader(form.Encode()))
	if err != nil {
		t.Errorf("Failed to form new request: %s", err.Error())
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// new client to do the request
	hc := &http.Client{}
	r, err := hc.Do(req)
	if err != nil {
		t.Errorf("Failed run the request: %s", err.Error())
	}

	log.Printf("Response: %#v", r)
}
