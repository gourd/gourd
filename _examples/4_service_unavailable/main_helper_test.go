package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"github.com/gourd/oauth2"
	"github.com/gourd/service"
	"github.com/yookoala/restit"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// could be used repeatedly for different unit test
func gourdTestServer() (ts *httptest.Server) {
	s := gourdServer()
	ts = httptest.NewServer(s)
	return
}

func dummyNewUser(password string) *oauth2.User {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randSeq := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	u := &oauth2.User{
		Username: randSeq(10),
	}
	u.Password = u.Hash(password)
	return u
}

func dummyNewClient(redirectUri string) *oauth2.Client {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randSeq := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	return &oauth2.Client{
		Id:          randSeq(10),
		Secret:      randSeq(10),
		RedirectUri: redirectUri,
		UserId:      randSeq(10),
	}
}

func dummyNewPost() (p Post) {
	s := rand.NewSource(99)
	return Post{
		UID:   rand.New(s).Int31(),
		Title: fmt.Sprintf("Dummy Post %d", rand.New(s).Int31()),
		Body:  fmt.Sprintf("Dummy Body %d", rand.New(s).Int31()),
		Size:  rand.New(s).Int63(),
		Date:  time.Now().AddDate(0, 0, -rand.New(s).Intn(700)),
	}
}

// example client web app in the login
func testOAuth2ClientApp(path string) http.Handler {
	rtr := pat.New()

	// add dummy client reception of redirection
	rtr.Get(path, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		enc := json.NewEncoder(w)
		enc.Encode(map[string]string{
			"code":  r.Form.Get("code"),
			"token": r.Form.Get("token"),
		})
	})

	return rtr
}

// test oauth2
func testOAuth2(t *testing.T, ts *httptest.Server) (token string) {

	// create test client server
	tcsbase := "/example_app/"
	tcspath := tcsbase + "code"
	tcs := httptest.NewServer(testOAuth2ClientApp(tcspath))
	defer tcs.Close()

	// a dummy password for dummy user
	password := "password"

	// create dummy oauth client and user
	c, u := func(tcs *httptest.Server, password, redirect string) (*oauth2.Client, *oauth2.User) {
		r := &http.Request{}

		// generate dummy user
		us, err := service.Providers.Service(r, "User")
		if err != nil {
			panic(err)
		}
		u := dummyNewUser(password)
		err = us.Create(service.NewConds(), u)
		if err != nil {
			panic(err)
		}

		// get related dummy client
		cs, err := service.Providers.Service(r, "Client")
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
	}(tcs, password, tcs.URL+tcsbase)

	// build user request to authorization endpoint
	// get response from client web app redirect uri
	code, err := func(c *oauth2.Client, u *oauth2.User, password, redirect string) (code string, err error) {

		log.Printf("Test retrieving code ====")

		// login form
		form := url.Values{}
		form.Add("username", u.Username)
		form.Add("password", password)
		log.Printf("form send: %s", form.Encode())

		// build the query string
		q := &url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", c.Id)
		q.Add("redirect_uri", redirect)

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

		log.Printf("Response.Request: %#v", resp.Request.URL)

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
	}(c, u, password, tcs.URL+tcspath)

	// quite if error
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// retrieve token from token endpoint
	// get response from client web app redirect uri
	token, err = func(c *oauth2.Client, code, redirect string) (token string, err error) {

		log.Printf("Test retrieving token ====")

		// build user request to token endpoint
		form := &url.Values{}
		form.Add("code", code)
		form.Add("client_id", c.Id)
		form.Add("client_secret", c.Secret)
		form.Add("grant_type", "authorization_code")
		form.Add("redirect_uri", redirect)
		req, err := http.NewRequest("POST",
			ts.URL+"/oauth/token",
			strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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
		var ok bool
		if token, ok = bodyDecoded["access_token"]; !ok {
			err = fmt.Errorf(
				"Unable to parse access_token: %s", err.Error())
		}
		return

	}(c, code, tcs.URL+tcspath)

	// quit if error
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	return token

}

// common rest test
func testRest(t *testing.T, ts *httptest.Server, token string, proto *ProtoPosts, name, path string) {

	// some dummy posts
	p1 := dummyNewPost()

	// REST API
	posts := restit.Rest("Post", ts.URL+"/api/posts")
	post := restit.Rest("Post", ts.URL+"/api/post")

	var err error

	// test list
	t0 := posts.Retrieve("").
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(0)
	_, err = t0.Run()

	if err != nil {
		t.Error(err.Error())
	} else if proto.Status == http.StatusOK {
		t.Error("Status is not \"error\"")
	} else if proto.Code != http.StatusInternalServerError {
		t.Errorf("Status code is not %d", http.StatusInternalServerError)
	}

	// test create
	t1 := posts.Create(&p1).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(0)
	_, err = t1.Run()
	if err != nil {
		t.Error(err.Error())
	} else if proto.Status == http.StatusOK {
		t.Error("Status is not \"error\"")
	} else if proto.Code != http.StatusInternalServerError {
		t.Errorf("Status code is not %d", http.StatusInternalServerError)
	}

	// test retrieve single
	t2 := post.Retrieve("someid").
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(0)
	_, err = t2.Run()
	if err != nil {
		t.Error(err.Error())
	} else if proto.Status == http.StatusOK {
		t.Error("Status is not \"error\"")
	} else if proto.Code != http.StatusInternalServerError {
		t.Errorf("Status code is not %d", http.StatusInternalServerError)
	}

	// test update
	p3 := dummyNewPost()
	t3 := post.Update("someid", p3).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(0)
	_, err = t3.Run()
	if err != nil {
		t.Error(err.Error())
	} else if proto.Status == http.StatusOK {
		t.Error("Status is not \"error\"")
	} else if proto.Code != http.StatusInternalServerError {
		t.Errorf("Status code is not %d", http.StatusInternalServerError)
	}

	// test delete
	t4 := post.Delete(fmt.Sprintf("%d", p3.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(0)
	_, err = t4.Run()
	if err != nil {
		t.Error(err.Error())
	} else if proto.Status == http.StatusOK {
		t.Error("Status is not \"error\"")
	} else if proto.Code != http.StatusInternalServerError {
		t.Errorf("Status code is not %d", http.StatusInternalServerError)
	}

}

// testing the gourd generated server
func TestMainPosts(t *testing.T) {
	ts := gourdTestServer()
	defer ts.Close()

	testRest(t, ts, "", &ProtoPosts{}, "Post", "/api/posts")
}
