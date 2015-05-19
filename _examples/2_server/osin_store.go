package main

import (
	"github.com/RangelReale/osin"
	"github.com/gourd/service"
	"net/http"
)

// OAuth2Storage implements osin.Storage
type OAuth2Storage struct {
	ClientService service.Service
	AuthService   service.Service
	AccessService service.Service
	UserService   service.Service
}

// UseClientIn set the ClientService
func (a *OAuth2Storage) UseClientIn(s service.Service) *OAuth2Storage {
	a.ClientService = s
	return a
}

// UseAuthIn set the AuthService
func (a *OAuth2Storage) UseAuthIn(s service.Service) *OAuth2Storage {
	a.AuthService = s
	return a
}

// UseAccessIn set the AccessService
func (a *OAuth2Storage) UseAccessIn(s service.Service) *OAuth2Storage {
	a.AccessService = s
	return a
}

// UseUserIn set the UserService
func (a *OAuth2Storage) UseUserIn(s service.Service) *OAuth2Storage {
	a.UserService = s
	return a
}

// Clone the storage
func (a *OAuth2Storage) Clone() (c osin.Storage) {
	c = &OAuth2Storage{
		ClientService: a.ClientService,
		AuthService:   a.AuthService,
		AccessService: a.AccessService,
		UserService:   a.UserService,
	}
	return
}

// Close the connection to the storage
func (a *OAuth2Storage) Close() {
	// placeholder now, will revisit when doing mongodb
}

// GetClient loads the client by id (client_id)
func (a *OAuth2Storage) GetClient(id string) (c osin.Client, err error) {

	e := &Client{}
	conds := service.NewConds()
	conds.Add("id", id)
	err = a.ClientService.One(conds, e)
	if err != nil {
		return
	} else if e == nil {
		err = service.Errorf(http.StatusNotFound,
			"Client not found for the id")
		return
	}

	c = e
	return
}

// SaveAuthorize saves authorize data.
func (a *OAuth2Storage) SaveAuthorize(d *osin.AuthorizeData) (err error) {
	e := &AuthorizeData{}
	err = e.ReadOsin(d)
	if err != nil {
		return
	}
	err = a.AuthService.Create(service.NewConds(), e)
	return
}

// LoadAuthorize looks up AuthorizeData by a code.
// Client information MUST be loaded together.
// Optionally can return error if expired.
func (a *OAuth2Storage) LoadAuthorize(code string) (d *osin.AuthorizeData, err error) {

	e := &AuthorizeData{}
	conds := service.NewConds()
	conds.Add("code", code)

	err = a.AuthService.One(conds, e)
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
func (a *OAuth2Storage) RemoveAuthorize(code string) (err error) {
	conds := service.NewConds()
	conds.Add("code", code)
	err = a.AuthService.Delete(conds)
	return
}

// SaveAccess writes AccessData.
// If RefreshToken is not blank, it must save in a way that can be loaded using LoadRefresh.
func (a *OAuth2Storage) SaveAccess(ad *osin.AccessData) (err error) {
	e := &AccessData{}
	err = e.ReadOsin(ad)
	if err != nil {
		return
	}
	err = a.AccessService.Create(service.NewConds(), e)
	return
}

// LoadAccess retrieves access data by token. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (a *OAuth2Storage) LoadAccess(token string) (d *osin.AccessData, err error) {

	e := &AccessData{}
	conds := service.NewConds()
	conds.Add("access_token", token)

	err = a.AccessService.One(conds, e)
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
func (a *OAuth2Storage) RemoveAccess(token string) (err error) {
	conds := service.NewConds()
	conds.Add("access_token", token)
	err = a.AccessService.Delete(conds)
	return
}

// LoadRefresh retrieves refresh AccessData. Client information MUST be loaded together.
// AuthorizeData and AccessData DON'T NEED to be loaded if not easily available.
// Optionally can return error if expired.
func (a *OAuth2Storage) LoadRefresh(token string) (d *osin.AccessData, err error) {

	e := &AccessData{}
	conds := service.NewConds()
	conds.Add("refresh_token", token)

	err = a.AccessService.One(conds, e)
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
func (a *OAuth2Storage) RemoveRefresh(token string) (err error) {
	conds := service.NewConds()
	conds.Add("refresh_token", token)
	err = a.AccessService.Delete(conds)
	return
}
