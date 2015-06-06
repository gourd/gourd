package main

import (
	"fmt"
)

// ProtoCommon represents a common protocol
type ProtoCommon struct {
	Status  string        `json:"status"`
	Code    int           `json:"code"`
	Msg     string        `json:"message,omitempty"`
	Results []interface{} `json:"results"`
}

// ProtoPosts represtnts protocol to communicate Posts
type ProtoPosts struct {
	Status string  `json:"status"`
	Code   int     `json:"code"`
	Msg    string  `json:"message,omitempty"`
	Posts  *[]Post `json:"posts,omitempty"`
	Post   *Post   `json:"post,omitempty"`
}

// Count counts the number of item in the protocol entity list
func (r *ProtoPosts) Count() int {
	if r.Posts == nil {
		return 0
	}
	ps := r.Posts
	return len(*ps)
}

// NthValid examines if the nth item in the entity list is valid
func (r *ProtoPosts) NthValid(n int) (err error) {
	// some test to see nth record is valid
	// such as:
	/*
		if r.Posts[n].Id == 0 {
			return fmt.Errorf("The thing has a StuffId = 0")
		}
	*/
	return
}

// GetNth retrieve the nth item from entity list
func (r *ProtoPosts) GetNth(n int) (item interface{}, err error) {
	// return the nth item
	ps := r.Posts
	item = (*ps)[n]
	return
}

// Match test if the given 2 items matches each other
func (r *ProtoPosts) Match(a interface{}, b interface{}) (err error) {

	// cast a and b back to Post
	realA := a.(Post)
	realB := b.(Post)

	if realA.ID != realB.ID && realA.ID != 0 && realB.ID != 0 {
		// if either ID is 0, do not check id match
		err = fmt.Errorf("ID not match")
	} else if realA.UID != realB.UID {
		err = fmt.Errorf("UID not match")
	} else if realA.Title != realB.Title {
		err = fmt.Errorf("Title not match")
	} else if realA.Body != realB.Body {
		err = fmt.Errorf("Body not match")
	} else if realA.Size != realB.Size {
		err = fmt.Errorf("Size not match")
	} else if realA.Date.Unix() != realB.Date.Unix() {
		err = fmt.Errorf("Date not match")
	}

	return
}

// Reset reset the current list
func (r *ProtoPosts) Reset() {
	r.Status = ""
	var ps []Post
	r.Posts = &ps
	r.Post = nil
}
