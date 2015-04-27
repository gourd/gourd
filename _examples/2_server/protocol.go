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
	Status string `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"message,omitempty"`
	Posts  []Post `json:"posts"`
}

func (r *ProtoPosts) Count() int {
	return len(r.Posts)
}

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

func (r *ProtoPosts) GetNth(n int) (item interface{}, err error) {
	// return the nth item
	item = r.Posts[n]
	return
}

func (r *ProtoPosts) Match(a interface{}, b interface{}) (err error) {

	// cast a and b back to Post
	real_a := a.(Post)
	real_b := b.(Post)

	if real_a.Id != real_b.Id && real_a.Id != 0 && real_b.Id != 0 {
		// if either ID is 0, do not check id match
		err = fmt.Errorf("Id not match")
	} else if real_a.Uid != real_b.Uid {
		err = fmt.Errorf("Uid not match")
	} else if real_a.Title != real_b.Title {
		err = fmt.Errorf("Title not match")
	} else if real_a.Body != real_b.Body {
		err = fmt.Errorf("Body not match")
	} else if real_a.Size != real_b.Size {
		err = fmt.Errorf("Size not match")
	} else if !real_a.Date.Equal(real_b.Date) {
		err = fmt.Errorf("Date not match")
	}

	return
}

func (r *ProtoPosts) Reset() {
	r.Status = ""
	r.Posts = make([]Post, 0)
}
