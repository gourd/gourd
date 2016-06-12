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

	"github.com/go-restit/lzjson"
	restit "github.com/go-restit/restit/v2"
	"github.com/gorilla/pat"
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
	s := rand.NewSource(time.Now().UnixNano())
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
func testOAuth2(t *testing.T, ctx context.Context, ts *httptest.Server) (user *oauth2.User, token string) {

	// create test client server
	tcsbase := "/example_app/"
	tcspath := tcsbase + "code"
	tcs := httptest.NewServer(testOAuth2ClientApp(tcspath))
	defer tcs.Close()

	// a dummy password for dummy user
	password := "password"

	// create dummy oauth client and user
	c, user := func(tcs *httptest.Server, password, redirect string) (*oauth2.Client, *oauth2.User) {

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
	}(c, user, password, tcs.URL+tcspath)

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

	return user, token

}

// common rest test
func testRest(t *testing.T, ts *httptest.Server, token string, user *oauth2.User, proto *ProtoPosts, name, path string) {

	// some dummy posts
	p1a := dummyNewPost()
	p1b := dummyNewPost()

	// define the path for your handler
	service := restit.NewHTTPService(ts.URL + "/api")

	equals := func(a interface{}) func(interface{}) error {
		return func(b interface{}) (err error) {

			// cast a and b back to Post
			original := a.(Post)
			val := b.(Post)

			if original.ID != val.ID && original.ID != 0 && val.ID != 0 {
				// if either ID is 0, do not check id match
				err = fmt.Errorf("ID not match")
			} else if want, have := original.UID, val.UID; want != have {
				err = fmt.Errorf("UID not match: wanted %#v but got %#v", want, have)
			} else if want, have := original.Title, val.Title; want != have {
				err = fmt.Errorf("Title not match: wanted %#v but got %#v", want, have)
			} else if want, have := original.Body, val.Body; want != have {
				err = fmt.Errorf("Body not match: wanted %#v but got %#v", want, have)
			} else if want, have := original.Size, val.Size; want != have {
				err = fmt.Errorf("Size not match: wanted %#v but got %#v", want, have)
			} else if want, have := original.Date.Unix(), val.Date.Unix(); want != have {
				err = fmt.Errorf("Date not match: wanted %#v but got %#v", want, have)
			}
			return
		}
	}

	nthEquals := func(n int, name, desc string, equals func(interface{}) error) restit.Expectation {
		return restit.Nth(n).Of(name).Is(restit.DescribeJSON(desc, func(node lzjson.Node) (err error) {
			var returned Post
			err = node.Unmarshal(&returned)
			if err != nil {
				return
			}
			return equals(returned)
		}))
	}

	pagingAttrIs := func(name string, number int) restit.Expectation {
		desc := fmt.Sprintf("%#v is %#v", name, number)
		return restit.Describe(desc, func(ctx context.Context, resp restit.Response) (err error) {
			root, err := resp.JSON()
			if err != nil {
				return
			}

			paging := root.Get("paging")
			if want, have := lzjson.TypeObject, paging.Type(); want != have {
				err = fmt.Errorf("\"paging\" is not a %s but is %s", want, have)
			}

			numGot := paging.Get(name)
			if want, have := lzjson.TypeNumber, numGot.Type(); want != have {
				err = fmt.Errorf("%#v is not a %s but is %s", name, want, have)
			}
			if want, have := number, numGot.Int(); want != have {
				err = fmt.Errorf("%#v is not %#v, but %#v", name, want, have)
			}
			return
		})
	}

	attrIsEmpty := func(noun string) restit.Expectation {
		return restit.Describe(noun, func(ctx context.Context, resp restit.Response) (err error) {
			root, err := resp.JSON()
			if err != nil {
				return
			}

			list := root.Get(noun)
			if list.Type() != lzjson.TypeArray {
				err = fmt.Errorf("%#v is not an array", noun)
			}
			if want, have := 0, list.Len(); want != have {
				err = fmt.Errorf("%#v is not %#v, but %#v", noun, want, have)
			}
			return
		})
	}

	var err error
	var resp restit.Response

	resp, err = service.Retrieve("/posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(attrIsEmpty("posts")).
		Expect(pagingAttrIs("limit", 0)).
		Expect(pagingAttrIs("offset", 0)).
		Expect(pagingAttrIs("total", 0)).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Create(p1a, "posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "created from p1a", equals(p1a))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	} else {
		root, _ := resp.JSON()
		root.Get("posts").GetN(0).Unmarshal(&p1a)
	}

	resp, err = service.Create(p1b, "posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "created from p1a", equals(p1b))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	} else {
		root, _ := resp.JSON()
		root.Get("posts").GetN(0).Unmarshal(&p1b)
	}

	resp, err = service.Retrieve("posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(pagingAttrIs("limit", 0)).
		Expect(pagingAttrIs("offset", 0)).
		Expect(pagingAttrIs("total", 2)).
		Expect(nthEquals(0, "posts", "equals p1a", equals(p1a))).
		Expect(nthEquals(1, "posts", "equals p1b", equals(p1b))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Retrieve(fmt.Sprintf("post/%d", p1a.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p1a", equals(p1a))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Retrieve(fmt.Sprintf("post/%d", p1b.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p1b", equals(p1b))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	// TODO: test retrieve list with dummy title

	p3 := dummyNewPost()
	p3.ID = p1a.ID
	resp, err = service.Update(p3, fmt.Sprintf("post/%d", p1a.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p3", equals(p3))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Retrieve(fmt.Sprintf("post/%d", p1a.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p3", equals(p3))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	// test partial update (POST to item endpoint)
	p4 := p3
	p4.UID = 1234
	p4Partial := struct {
		UID int32 `json:"uid"`
	}{UID: p4.UID}
	resp, err = service.Create(p4Partial, fmt.Sprintf("post/%d", p1a.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p4", equals(p4))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Delete(fmt.Sprintf("post/%d", p1a.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p4", equals(p4))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Retrieve("posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(pagingAttrIs("limit", 0)).
		Expect(pagingAttrIs("offset", 0)).
		Expect(pagingAttrIs("total", 1)).
		Expect(nthEquals(0, "posts", "equals p1b", equals(p1b))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Delete(fmt.Sprintf("post/%d", p1b.ID)).
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(nthEquals(0, "posts", "equals p1b", equals(p1b))).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

	resp, err = service.Retrieve("posts").
		AddHeader("Authority", token).
		Expect(restit.StatusCodeIs(http.StatusOK)).
		Expect(pagingAttrIs("limit", 0)).
		Expect(pagingAttrIs("offset", 0)).
		Expect(pagingAttrIs("total", 0)).
		Do()
	if err != nil {
		t.Logf("raw body: %s", resp)
		t.Error(err.Error())
	}

}

// testing the gourd generated server
func TestMainPosts(t *testing.T) {
	factory := NewFactory(`./data/sqlite3.db`)
	ts := httptest.NewServer(NewHandler(factory))
	defer ts.Close()

	ctx := store.WithFactory(context.Background(), factory)

	user, token := testOAuth2(t, ctx, ts)
	if token == "" {
		t.Error("Failed to obtain security token")
		t.FailNow()
		return
	}
	log.Printf("token: %s", token)
	testRest(t, ts, token, user, &ProtoPosts{}, "Post", "/api/posts")
}
