package example

import (
	"github.com/gourd/kit/oauth2"
	"github.com/gourd/kit/store"
	"github.com/gourd/kit/store/upperio"
	"golang.org/x/net/context"

	"math/rand"
	"testing"

	"upper.io/db/sqlite"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var ctx context.Context

func init() {

	// define db
	factory := store.NewFactory()
	factory.SetSource(store.DefaultSrc, upperio.Source(sqlite.Adapter, sqlite.ConnectionURL{
		Database: `./data/sqlite3.db`,
	}))
	factory.Set("UserA", store.DefaultSrc, UserAStoreProvider)
	factory.Set("UserB", store.DefaultSrc, UserBStoreProvider)
	factory.Set(oauth2.KeyAccess, store.DefaultSrc, oauth2.AccessDataStoreProvider)
	factory.Set(oauth2.KeyAuth, store.DefaultSrc, oauth2.AuthorizeDataStoreProvider)
	factory.Set(oauth2.KeyClient, store.DefaultSrc, oauth2.ClientStoreProvider)
	factory.Set(oauth2.KeyUser, store.DefaultSrc, oauth2.UserStoreProvider)

	ctx = store.WithFactory(context.Background(), factory)
}

func randStr(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)

}

func TestUserAStore(t *testing.T) {

	us, err := store.Get(ctx, "UserA")
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

	us, err := store.Get(ctx, "UserB")
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
