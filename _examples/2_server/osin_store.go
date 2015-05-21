package main

import (
	"github.com/RangelReale/osin"
	"github.com/gourd/service"
	"log"
	"net/http"
)

// OAuth2Storage implements osin.Storage
type OAuth2Storage struct {
	r             *http.Request
	ClientService service.ProviderFunc
	AuthService   service.ProviderFunc
	AccessService service.ProviderFunc
	UserService   service.ProviderFunc
}

// SetRequest set the ClientService
func (store *OAuth2Storage) SetRequest(r *http.Request) *OAuth2Storage {
	store.r = r
	return store
}

// UseClientFrom set the ClientService
func (store *OAuth2Storage) UseClientFrom(p service.ProviderFunc) *OAuth2Storage {
	store.ClientService = p
	return store
}

// UseAuthFrom set the AuthService
func (store *OAuth2Storage) UseAuthFrom(p service.ProviderFunc) *OAuth2Storage {
	store.AuthService = p
	return store
}

// UseAccessFrom set the AccessService
func (store *OAuth2Storage) UseAccessFrom(p service.ProviderFunc) *OAuth2Storage {
	store.AccessService = p
	return store
}

// UseUserFrom set the UserService
func (store *OAuth2Storage) UseUserFrom(p service.ProviderFunc) *OAuth2Storage {
	store.UserService = p
	return store
}

// Clone the storage
func (store *OAuth2Storage) Clone() (c osin.Storage) {
	c = &OAuth2Storage{
		ClientService: store.ClientService,
		AuthService:   store.AuthService,
		AccessService: store.AccessService,
		UserService:   store.UserService,
	}
	return
}

// Close the connection to the storage
func (store *OAuth2Storage) Close() {
	// placeholder now, will revisit when doing mongodb
}

// GetClient loads the client by id (client_id)
func (store *OAuth2Storage) GetClient(strId string) (c osin.Client, err error) {

	log.Printf("GetClient %s", strId)

	srv, err := store.ClientService(store.r)
	if err != nil {
		log.Printf("Unable to get client service")
		return
	}

	e := &Client{}
	conds := service.NewConds()
	conds.Add("str_id", strId)

	err = srv.One(conds, e)
	if err != nil {
		log.Printf("%#v", conds)
		log.Printf("Failed running One()")
		return
	} else if e == nil {
		log.Printf("Client not found for the str_id")
		err = service.Errorf(http.StatusNotFound,
			"Client not found for the str_id")
		return
	}

	c = e
	return
}

// SaveAuthorize saves authorize data.
func (store *OAuth2Storage) SaveAuthorize(d *osin.AuthorizeData) (err error) {

	log.Printf("SaveAuthorize %v", d)

	srv, err := store.AuthService(store.r)
	if err != nil {
		return
	}

	e := &AuthorizeData{}
	err = e.ReadOsin(d)
	if err != nil {
		return
	}
	err = srv.Create(service.NewConds(), e)
	return
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (store *OAuth2Storage) LoadAuthorize(code string) (d *osin.AuthorizeData, err error) {

	log.Printf("LoadAuthorize %s", code)

	srv, err := store.AuthService(store.r)
	if err != nil {
		return
	}

	e := &AuthorizeData{}
	conds := service.NewConds()
	conds.Add("code", code)

	err = srv.One(conds, e)
	if err != nil {
		return
	} else if e == nil {
		err = service.Errorf(http.StatusNotFound,
			"AuthorizeData not found for the code")
		return
	}

	d = e.ToOsin()
	return
}

// RemoveAuthorize revokes or deletes the authorization code.
func (store *OAuth2Storage) RemoveAuthorize(code string) (err error) {

	log.Printf("RemoveAuthorize %s", code)

	srv, err := store.AuthService(store.r)
	if err != nil {
		return
	}

	conds := service.NewConds()
	conds.Add("code", code)
	err = srv.Delete(conds)
	return
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (store *OAuth2Storage) SaveAccess(ad *osin.AccessData) (err error) {

	log.Printf("SaveAccess %v", ad)

	srv, err := store.AccessService(store.r)
	if err != nil {
		return
	}

	e := &AccessData{}
	err = e.ReadOsin(ad)
	if err != nil {
		return
	}
	err = srv.Create(service.NewConds(), e)
	return
}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (store *OAuth2Storage) LoadAccess(token string) (d *osin.AccessData, err error) {

	log.Printf("Loadaccess %v", token)

	srv, err := store.AccessService(store.r)
	if err != nil {
		return
	}

	e := &AccessData{}
	conds := service.NewConds()
	conds.Add("access_token", token)

	err = srv.One(conds, e)
	if err != nil {
		return
	} else if e == nil {
		err = service.Errorf(http.StatusNotFound,
			"AccessData not found for the token")
		return
	}

	d = e.ToOsin()
	return
}

// RemoveAccess revokes or deletes an AccessData.
func (store *OAuth2Storage) RemoveAccess(token string) (err error) {

	log.Printf("RemoveAccess %v", token)

	srv, err := store.AccessService(store.r)
	if err != nil {
		return
	}

	conds := service.NewConds()
	conds.Add("access_token", token)
	err = srv.Delete(conds)
	return
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (store *OAuth2Storage) LoadRefresh(token string) (d *osin.AccessData, err error) {

	log.Printf("LoadRefresh %v", token)

	srv, err := store.AccessService(store.r)
	if err != nil {
		return
	}

	e := &AccessData{}
	conds := service.NewConds()
	conds.Add("refresh_token", token)

	err = srv.One(conds, e)
	if err != nil {
		return
	} else if e == nil {
		err = service.Errorf(http.StatusNotFound,
			"AccessData not found for the refresh token")
		return
	}

	d = e.ToOsin()
	return
}

// RemoveRefresh revokes or deletes refresh AccessData.
func (store *OAuth2Storage) RemoveRefresh(token string) (err error) {

	log.Printf("RemoveRefresh %v", token)

	srv, err := store.AccessService(store.r)
	if err != nil {
		return
	}

	conds := service.NewConds()
	conds.Add("refresh_token", token)
	err = srv.Delete(conds)
	return
}
