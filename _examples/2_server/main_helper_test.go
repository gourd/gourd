package main

import (
	"fmt"
	"github.com/yookoala/restit"
	"math/rand"
	"net/http/httptest"
	"testing"
	"time"
)

// could be used repeatedly for different unit test
func gourdTestServer() (ts *httptest.Server) {
	s := gourdServer()
	ts = httptest.NewServer(s)
	return
}

func dummyNewPost() (p Post) {
	s := rand.NewSource(99)
	return Post{
		Uid:   rand.New(s).Int31(),
		Title: fmt.Sprintf("Dummy Post %d", rand.New(s).Int31()),
		Body:  fmt.Sprintf("Dummy Body %d", rand.New(s).Int31()),
		Size:  rand.New(s).Int63(),
		Date:  time.Now().AddDate(0, 0, -rand.New(s).Intn(700)),
	}
}

// common rest test
func testRest(t *testing.T, proto restit.Response, name, path string) {
	ts := gourdTestServer()
	defer ts.Close()

	// some dummy posts
	p1 := dummyNewPost()

	// REST API
	posts := restit.Rest("Post", ts.URL+"/api/posts")
	post := restit.Rest("Post", ts.URL+"/api/post")

	// test create
	t1 := posts.Create(&p1).
		WithResponseAs(proto).
		ExpectResultCount(1).
		ExpectResultsValid().
		ExpectResultNth(0, p1)
	_, err := t1.Run()
	p2p, err := proto.GetNth(0)
	p2 := p2p.(Post) // created post

	// test retrieve single
	t2 := post.Retrieve(fmt.Sprintf("%d", p2.Id)).
		WithResponseAs(proto).
		ExpectResultCountNot(0)
	_, err = t2.Run()

	// TODO: test retrieve list, expecte to have non-zero count
	// TODO: test retrieve list with dummy title

	// TODO: test update
	// TODO: test delete

	if err != nil {
		t.Error(err.Error())
	}
}

// testing the gourd generated server
func TestMainPosts(t *testing.T) {
	testRest(t, &ProtoPosts{}, "Post", "/api/posts")
}
