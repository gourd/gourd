package main

import (
	"github.com/yookoala/restit"
	"testing"
)

func TestProtoPosts(t *testing.T) {
	var p restit.Response = &ProtoPosts{}
	p.Count()
	t.Log("ProtoPosts implements restit.Response")
}
