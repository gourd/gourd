package example1

import (
	"github.com/gourd/service"
	"testing"
)

// Test if the PostService is generated correctly
func TestService(t *testing.T) {
	var s service.Service = new(PostService)
	l := s.AllocEntityList()
	if s.Len(l) != 0 {
		t.Errorf("AllocEntityList is not allocating empty list: %#v", l)
	}
}
