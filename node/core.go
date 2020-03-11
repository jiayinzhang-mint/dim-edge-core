package node

import (
	"dim-edge-core/protocol"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client dim-edge-node client instance
type Client struct {
	Address string
	Options []grpc.DialOption
	Conn    *grpc.ClientConn
	Store   protocol.StoreServiceClient
}

// New create new node client
func (c *Client) New() (err error) {
	c.Options = append(c.Options, grpc.WithInsecure())

	c.Conn, err = grpc.Dial(c.Address, c.Options...)
	if err != nil {
		err = errors.Wrapf(err,
			"Failed to start grpc connection with address %s",
			c.Address)
		return
	}

	c.Store = protocol.NewStoreServiceClient(c.Conn)

	return
}
