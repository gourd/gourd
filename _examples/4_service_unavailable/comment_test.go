package main

import (
	"github.com/gourd/kit/store"
	"testing"
)

// Test if the CommentStore is generated correctly
func TestCommentStore(t *testing.T) {
	var s store.Store = new(CommentStore)
	l := s.AllocEntityList()
	if s.Len(l) != 0 {
		t.Errorf("AllocEntityList is not allocating empty list: %#v", l)
	}
}
