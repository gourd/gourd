//go:generate gourd gen service -type=Post $GOFILE

package gourd_example

import (
	"time"
)

type Post struct {
	Id    int32  `db:"id"`
	Title string `db:"title" json:"title"`
	Body  string `db:"body" json:"body"`
	Size  int64  `db:"size"`
	Date  time.Time
}
