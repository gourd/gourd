package main

import (
	restit "github.com/yookoala/restit/v1"
	"testing"
)

func TestProtoPosts(t *testing.T) {
	var p restit.Response = &ProtoPosts{}
	p.Count()
	t.Log("ProtoPosts implements restit.Response")
}
