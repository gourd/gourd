//go:generate gourd gen store -type=UserA -coll=user_a $GOFILE

package example

type UserA struct {
	ID       int64  `db:"id,omitempty"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
