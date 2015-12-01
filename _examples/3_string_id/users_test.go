package example

import (
	"github.com/gourd/kit/store"
	"github.com/gourd/kit/store/upperio"

	"math/rand"
	"net/http"
	"testing"
	"upper.io/db/sqlite"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	// define db
	upperio.Define("default", sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./data/sqlite3.db`,
	})
}

func randStr(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}

func TestUserAStore(t *testing.T) {

	r := &http.Request{}
	usp, err := store.Providers.Get("UserA")
	if err != nil {
		t.Logf("Unable to obtain store provider: %s", err.Error())
	}
	us, err := usp.Store(r)
	if err != nil {
		t.Logf("Unable to obtain store: %s", err.Error())
	}

	// create user twice, test key collision
	u1 := UserA{
		Name:     randStr(12),
		Password: randStr(12),
	}
	err = us.Create(nil, &u1)
	if err != nil {
		t.Logf("Error creating user: %s", err.Error())
	}
	t.Logf("%#v", u1)

	u2 := UserA{
		Name:     randStr(12),
		Password: randStr(12),
	}
	err = us.Create(nil, &u2)
	if err != nil {
		t.Logf("Error creating user: %s", err.Error())
	}
	t.Logf("%#v", u2)

}

func TestUserBStore(t *testing.T) {

	r := &http.Request{}
	usp, err := store.Providers.Get("UserB")
	if err != nil {
		t.Logf("Unable to obtain store provider: %s", err.Error())
	}
	us, err := usp.Store(r)
	if err != nil {
		t.Logf("Unable to obtain store: %s", err.Error())
	}

	// create user twice, test key collision
	u1 := UserB{
		Name:     randStr(12),
		Password: randStr(12),
	}
	err = us.Create(nil, &u1)
	if err != nil {
		t.Logf("Error creating user: %s", err.Error())
	}
	t.Logf("%#v", u1)

	u2 := UserB{
		Name:     randStr(12),
		Password: randStr(12),
	}
	err = us.Create(nil, &u2)
	if err != nil {
		t.Logf("Error creating user: %s", err.Error())
	}
	t.Logf("%#v", u2)

}
