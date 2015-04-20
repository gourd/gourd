package main

import (
	"net/http/httptest"
	"testing"
)

// could be used repeatedly for different unit test
func gourdTestServer() (ts *httptest.Server) {
	s := gourdServer()
	ts = httptest.NewServer(s)
	return
}

// testing the gourd generated server
func TestMain(t *testing.T) {
	ts := gourdTestServer()
	defer ts.Close()
	//do something with the test server
}
