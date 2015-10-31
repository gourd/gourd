//go:generate gourd gen service -type=Post -coll=posts $GOFILE
//go:generate gourd gen rest -type=Post -service=PostService post_service.go

package main

import (
	"time"
)

// Post is for blog post
// This is a multiple line comment
/**
 * This is another multiple line comment
 * Just another type
 */
type Post struct {

	// This is Id
	ID int32 `db:"id,omitempty" json:"id"`

	// This is User ID
	UID int32 `db:"uid" json:"uid"`

	// title of the post
	Title string `db:"title" json:"title"`

	// HTML body of the post
	Body string `db:"body" json:"body"`

	// size in byte of the post
	Size int64 `db:"size" json:"size"`

	// date when the post is published
	Date time.Time `db:"date" json:"date"`
}
