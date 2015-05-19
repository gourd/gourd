//go:generate gourd gen service -type=User -coll=user $GOFILE
package main

import (
	"time"
)

// User of the API server
type User struct {
	Id       int64     `db:"id"`
	Username string    `db:"username"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Name     string    `db:"name"`
	Created  time.Time `db:"created"`
	Updated  time.Time `db:"updated"`
}
