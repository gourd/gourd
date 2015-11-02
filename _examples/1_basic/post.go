//go:generate gourd gen store -type=Post -coll=posts $GOFILE

package example1

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
	Id int32 `db:"id"`

	// This is User ID
	Uid int32 `db:"uid"`

	// title of the post
	Title string `db:"title" json:"title"`

	// HTML body of the post
	Body string `db:"body" json:"body"`

	// size in byte of the post
	Size int64 `db:"size"`

	// date when the post is published
	Date time.Time
}
