package example

import (
	"github.com/gourd/service"
	"github.com/gourd/service/upperio"

	"math/rand"
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

func TestUserAService(t *testing.T) {

	usp, err := service.Providers.Get("UserA")
	if err != nil {
		t.Logf("Unable to obtain service provider: %s", err.Error())
	}
	us, err := usp.Service(nil)
	if err != nil {
		t.Logf("Unable to obtain service: %s", err.Error())
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

func TestUserBService(t *testing.T) {

	usp, err := service.Providers.Get("UserB")
	if err != nil {
		t.Logf("Unable to obtain service provider: %s", err.Error())
	}
	us, err := usp.Service(nil)
	if err != nil {
		t.Logf("Unable to obtain service: %s", err.Error())
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
