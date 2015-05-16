package main

import (
	"testing"
)

// test if BasicScopes implements Scopes
func TestBasicScopes(t *testing.T) {
	var s Scopes = &BasicScopes{}
	_ = s
}
