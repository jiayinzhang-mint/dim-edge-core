package influxdb

import (
	"context"
	"dim-edge-core/protocol"
)

// QueryData query data
func (c *Client) QueryData(p *protocol.QueryParams) (r *protocol.QueryRes, err error) {
	r, err = c.Store.QueryData(context.TODO(), p)
	return
}

// InsertData insert data
func (c *Client) InsertData(p *protocol.InsertDataParams) (r *protocol.InsertDataRes, err error) {
	r, err = c.Store.InsertData(context.TODO(), p)
	return
}
