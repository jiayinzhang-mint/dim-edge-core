package influxdb

import (
	"context"
	"dim-edge-core/protocol"
)

// ListAuthorization list all authorizations
func (c *Client) ListAuthorization(p *protocol.ListAuthorizationParams) (a []*protocol.Authorization, err error) {
	res, err := c.Store.ListAuthorization(context.TODO(), p)
	a = res.Authorization

	return
}

// CreateAuthorization create an authorization
func (c *Client) CreateAuthorization(p *protocol.CreateAuthorizationParams) (a *protocol.Authorization, err error) {
	a, err = c.Store.CreateAuthorization(context.TODO(), p)

	return
}
