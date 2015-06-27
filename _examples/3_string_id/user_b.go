//go:generate gourd gen service -type=UserB -coll=user_b $GOFILE

package example

type UserB struct {
	ID       string `db:"id,omitempty"`
	Name     string `db:"name"`
	Password string `db:"password"`
}
