//go:generate gourd gen service -type=Comment -coll=comments $GOFILE

package main

import (
	"time"
)

// Comment is for blog comment
type Comment struct {

	// This is Id of cmment
	Id int32 `db:"id"`

	// This is Post ID
	PostId int32 `db:"post_id"`

	// title of the post
	Title string `db:"title" json:"title"`

	// HTML body of the post
	Body string `db:"body" json:"body"`

	// size in byte of the post
	Size int64 `db:"size"`

	// date when the post is published
	Date time.Time
}
