package main

import (
	"github.com/gourd/kit/store"
	"testing"
)

// Test if the PostStore is generated correctly
func TestPostStore(t *testing.T) {
	var s store.Store = new(PostStore)
	l := s.AllocEntityList()
	if s.Len(l) != 0 {
		t.Errorf("AllocEntityList is not allocating empty list: %#v", l)
	}
}
