package main

import (
	"net/http/httptest"
	"testing"
)

// could be used repeatedly for different unit test
func getTestServer() (ts *httptest.Server) {
	s := getServer()
	ts = httptest.NewServer(s)
	return
}

// testing the gourd generated server
func TestMain(t *testing.T) {
	ts := getTestServer()
	defer ts.Close()
	//do something with the test server
}
