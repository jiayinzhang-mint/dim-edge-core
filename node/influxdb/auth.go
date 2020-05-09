package influxdb

import (
	"context"
	"dim-edge/core/protocol"

	"github.com/golang/protobuf/ptypes/empty"
)

// ListAuthorization list all authorizations
func (c *Client) ListAuthorization(p *protocol.ListAuthorizationParams) (a []*protocol.Authorization, err error) {
	res, err := c.Store.ListAuthorization(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	a = res.Authorization

	return
}

// CreateAuthorization create an authorization
func (c *Client) CreateAuthorization(p *protocol.CreateAuthorizationParams) (a *protocol.Authorization, err error) {
	a, err = c.Store.CreateAuthorization(context.TODO(), p)

	return
}

// RetrieveAuthorization retrieve one authorization with auth id
func (c *Client) RetrieveAuthorization(p *protocol.RetrieveAuthorizationParams) (a *protocol.Authorization, err error) {
	a, err = c.Store.RetrieveAuthorization(context.TODO(), p)
	return
}

// SignIn sign into influxdb
func (c *Client) SignIn(p *protocol.SignInParams) (err error) {
	_, err = c.Store.SignIn(context.TODO(), p)

	return
}

// SignOut quit influxdb
func (c *Client) SignOut() (err error) {
	_, err = c.Store.SignOut(context.TODO(), &empty.Empty{})

	return
}

// GetMe get my info after signing in
func (c *Client) GetMe() (me *protocol.Me, err error) {
	me, err = c.Store.GetMe(context.TODO(), &empty.Empty{})

	return
}
