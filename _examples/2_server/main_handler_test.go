package main

import (
	"github.com/gourd/kit/oauth2"
	"github.com/gourd/kit/store"
	"golang.org/x/net/context"

	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/pat"
	"github.com/yookoala/restit"
)

// could be used repeatedly for different unit test
func mainTestServer() (ts *httptest.Server) {
	ts = httptest.NewServer(NewHandler(NewFactory(`./data/sqlite3.db`)))
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

func dummyNewClient(redirectURI string) *oauth2.Client {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randSeq := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	}

	return &oauth2.Client{
		ID:          randSeq(10),
		Secret:      randSeq(10),
		RedirectURI: redirectURI,
		UserID:      randSeq(10),
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
func testOAuth2(t *testing.T, ctx context.Context, ts *httptest.Server) (token string) {

	// create test client server
	tcsbase := "/example_app/"
	tcspath := tcsbase + "code"
	tcs := httptest.NewServer(testOAuth2ClientApp(tcspath))
	defer tcs.Close()

	// a dummy password for dummy user
	password := "password"

	// create dummy oauth client and user
	c, u := func(tcs *httptest.Server, password, redirect string) (*oauth2.Client, *oauth2.User) {

		// generate dummy user
		us, err := store.Get(ctx, oauth2.KeyUser)
		if err != nil {
			panic(err)
		}
		u := dummyNewUser(password)
		err = us.Create(store.NewConds(), u)
		if err != nil {
			panic(err)
		}

		// get related dummy client
		cs, err := store.Get(ctx, oauth2.KeyClient)
		if err != nil {
			panic(err)
		}
		c := dummyNewClient(redirect)
		c.UserID = u.ID
		err = cs.Create(store.NewConds(), c)
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
		form.Add("user_id", u.Username)
		form.Add("password", password)
		log.Printf("form send: %s", form.Encode())

		// build the query string
		q := &url.Values{}
		q.Add("response_type", "code")
		q.Add("client_id", c.ID)
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
		form.Add("client_id", c.ID)
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
		if err := dec.Decode(&bodyDecoded); err != nil {
			err = fmt.Errorf("Unable to parse access_token: %#v", err.Error())
		}

		log.Printf("Response Body: %#v", bodyDecoded)
		var ok bool
		if token, ok = bodyDecoded["access_token"]; !ok {
			err = fmt.Errorf("Unable to find access_token in response")
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
	p1a := dummyNewPost()
	p1b := dummyNewPost()

	// REST API
	posts := restit.Rest("Post", ts.URL+"/api/posts")
	post := restit.Rest("Post", ts.URL+"/api/post")

	var err error

	// test list
	t0 := posts.Retrieve("").
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectStatus(http.StatusOK).
		ExpectResultCount(0)
	_, err = t0.Run()
	if err != nil {
		t.Error(err.Error())
	} else if want, have := http.StatusOK, proto.Status; want != have {
		t.Errorf("want: %#v, got %#v",
			want, have)
		t.FailNow()
	} else if proto.Posts == nil {
		t.Errorf("Posts field should be empty array but get \"%#v\"",
			proto.Posts)
		t.FailNow()
	}

	// test paging variables
	if v, ok := proto.Paging["total"]; !ok {
		t.Errorf("paging.total not found")
	} else if want, have := 0, v; want != have {
		t.Errorf("paging.total expect: %#v, got: %#v", want, have)
	}
	if v, ok := proto.Paging["limit"]; !ok {
		t.Errorf("paging.limit not found")
	} else if want, have := 0, v; want != have {
		t.Errorf("paging.limit expect: %#v, got: %#v", want, have)
	}
	if v, ok := proto.Paging["offset"]; !ok {
		t.Errorf("paging.offset not found")
	} else if want, have := 0, v; want != have {
		t.Errorf("paging.offset expect: %#v, got: %#v", want, have)
	}

	// test create a
	t1a := posts.Create(&p1a).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(1).
		ExpectResultsValid().
		ExpectResultNth(0, p1a)
	_, err = t1a.Run()
	p2ap, err := proto.GetNth(0)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	p2a := p2ap.(Post) // created post

	// test create a
	t1b := posts.Create(&p1b).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(1).
		ExpectResultsValid().
		ExpectResultNth(0, p1b)
	_, err = t1b.Run()
	p2bp, err := proto.GetNth(0)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	p2b := p2bp.(Post) // created post

	// test retrieve single
	t2a := post.Retrieve(fmt.Sprintf("%d", p2a.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCountNot(0)
	_, err = t2a.Run()
	if err != nil {
		t.Error(err.Error())
	}
	t2b := post.Retrieve(fmt.Sprintf("%d", p2b.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCountNot(0)
	_, err = t2b.Run()
	if err != nil {
		t.Error(err.Error())
	}

	// test retrieve list
	t2l := posts.Retrieve("").
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(2)
	_, err = t2l.Run()
	if err != nil {
		t.Fatal(err.Error())
	}

	// test paging variables
	if v, ok := proto.Paging["total"]; !ok {
		t.Errorf("paging.total not found")
	} else if want, have := 2, v; want != have {
		t.Errorf("paging.total expect: %#v, got: %#v", want, have)
	}
	if v, ok := proto.Paging["limit"]; !ok {
		t.Errorf("paging.limit not found")
	} else if want, have := 0, v; want != have {
		t.Errorf("paging.limit expect: %#v, got: %#v", want, have)
	}
	if v, ok := proto.Paging["offset"]; !ok {
		t.Errorf("paging.offset not found")
	} else if want, have := 0, v; want != have {
		t.Errorf("paging.offset expect: %#v, got: %#v", want, have)
	}

	// TODO: test retrieve list with dummy title

	// test update, then retrieve single to compare
	p3 := dummyNewPost()
	p3.ID = p2a.ID
	t3 := post.Update(fmt.Sprintf("%d", p3.ID), p3).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectStatus(200)
	_, err = t3.Run()
	if err != nil {
		t.Error(err.Error())
	}
	t4 := post.Retrieve(fmt.Sprintf("%d", p3.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectStatus(200). // Success
		ExpectResultCount(1).
		ExpectResultNth(0, p3)
	_, err = t4.Run()
	if err != nil {
		t.Error(err.Error())
	}

	// test delete
	t5 := post.Delete(fmt.Sprintf("%d", p3.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectStatus(200)
	_, err = t5.Run()
	if err != nil {
		t.Error(err.Error())
	}
	t6 := post.Retrieve(fmt.Sprintf("%d", p3.ID)).
		AddHeader("Authority", token).
		WithResponseAs(proto) // Not found anymore
	_, err = t6.Run()
	if err != nil {
		t.Error(err.Error())
	}
	// test retrieve list
	t6l := posts.Retrieve("").
		AddHeader("Authority", token).
		WithResponseAs(proto).
		ExpectResultCount(1)
	_, err = t6l.Run()
	if err != nil {
		t.Fatal(err.Error())
	}

}

// testing the gourd generated server
func TestMainPosts(t *testing.T) {
	factory := NewFactory(`./data/sqlite3.db`)
	ts := httptest.NewServer(NewHandler(factory))
	defer ts.Close()

	ctx := store.WithFactory(context.Background(), factory)

	token := testOAuth2(t, ctx, ts)
	if token == "" {
		t.Error("Failed to obtain security token")
		t.FailNow()
		return
	}
	log.Printf("token: %s", token)
	testRest(t, ts, token, &ProtoPosts{}, "Post", "/api/posts")
}
