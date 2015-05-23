package main

import (
	"encoding/json"
	"github.com/gorilla/pat"
	"github.com/gourd/service"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"fmt"
)

// example client web app in the login
func testOAuth2ClientApp(path1, path2 string) http.Handler {
	rtr := pat.New()

	// add dummy client reception of redirection
	rtr.Get(path1, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		enc := json.NewEncoder(w)
		enc.Encode(map[string]string{
			"code": r.Form.Get("code"),
		})
	})

	rtr.Get(path2, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "success")
	})

	return rtr
}

func TestOAuth2(t *testing.T) {

	// create test oauth2 server
	ts := gourdTestServer()
	defer ts.Close()

	// create test client server
	tcspath1 := "/example_app/code"
	tcspath2 := "/example_app/success"
	tcs := httptest.NewServer(testOAuth2ClientApp(tcspath1, tcspath2))
	defer tcs.Close()

	// a dummy password for dummy user
	password := "password"

	// create dummy oauth client and user
	c, u := func(tcs *httptest.Server, password, redirect string) (*Client, *User) {
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
		c := dummyNewClient(redirect)
		c.UserId = u.Id
		err = cs.Create(service.NewConds(), c)
		if err != nil {
			panic(err)
		}

		return c, u
	}(tcs, password, tcs.URL + tcspath1)

	// build user request to authorization endpoint
	// get response from client web app redirect uri
	code, err := func (c *Client, u *User, password string) (code string, err error) {

		// login form
		form := url.Values{}
		form.Add("username", u.Username)
		form.Add("password", password)
		log.Printf("form send: %s", form.Encode())

		// build the query string
		q := &url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", c.StrId)
		q.Add("redirect_url", c.RedirectUri)

		req, err := http.NewRequest("POST",
			ts.URL+"/oauth/authorize"+"?"+q.Encode(),
			strings.NewReader(form.Encode()))
		if err != nil {
			err = fmt.Errorf("Failed to form new request: %s", err.Error())
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		// new http client to emulate user request
		hc := &http.Client{}
		resp, err := hc.Do(req)
		if err != nil {
			err = fmt.Errorf("Failed run the request: %s", err.Error())
		}

		// request should be redirected to client app with code
		// the testing client app response with a json containing "code"
		// decode the client app json and retrieve the code
		bodyDecoded := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&bodyDecoded)
		var ok bool
		if code, ok = bodyDecoded["code"]; !ok {
			err = fmt.Errorf("Client app failed to retrieve code in the redirection")
		}
		log.Printf("Response Body: %#v", bodyDecoded["code"])

		return
	}(c, u, password)

	// quite if error
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// retrieve token from token endpoint
	// get response from client web app redirect uri
	token, err := func (c *Client, code, redirect string) (token string, err error) {

		log.Printf("Begin token request")

		// build user request to token endpoint
		form := &url.Values{}
		form.Add("code", code)
		form.Add("client_id", c.StrId)
		form.Add("client_secret", c.Secret)
		form.Add("grant_type", "authorization_code")
		form.Add("redirect_uri", redirect)
		log.Printf("token request :%s", form.Encode())
		req, err := http.NewRequest("POST",
			ts.URL+"/oauth/token",
			strings.NewReader(form.Encode()))
		if err != nil {
			t.Errorf("Failed to form new request: %s", err.Error())
		}

		// new http client to emulate user request
		hc := &http.Client{}
		resp, err := hc.Do(req)
		if err != nil {
			err = fmt.Errorf("Failed run the request: %s", err.Error())
		}

		// read token from token endpoint response (json)
		bodyDecoded := make(map[string]string)
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&bodyDecoded)
		log.Printf("Response Body: %#v", bodyDecoded)
		return

	}(c, code, tcs.URL + tcspath2)

	// quite if error
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	log.Printf("token: \"%s\"", token)

}
