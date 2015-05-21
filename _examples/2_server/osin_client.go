//go:generate gourd gen service -type=Client -coll=oauth2_client -output=osin_client_service.go $GOFILE

package main

// Client implements the osin Client interface
//
type Client struct {
	Id          int64       `db:"id,omitempty"`
	StrId       string      `db:"str_id"`
	Secret      string      `db:"secret"`
	RedirectUri string      `db:"redirect_uri"`
	UserId      int64       `db:"user_id"`
	UserData    interface{} `db:"-"`
}

func (c *Client) GetId() string {
	return c.StrId
}

func (c *Client) GetSecret() string {
	return c.Secret
}

func (c *Client) GetRedirectUri() string {
	return c.RedirectUri
}

func (c *Client) GetUserData() interface{} {
	return c.UserData
}
