package influxdb

import (
	"dim-edge/core/protocol"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"google.golang.org/grpc"
)

// Client dim-edge-node client instance
type Client struct {
	Address string
	Options []grpc.DialOption
	Conn    *grpc.ClientConn
	Store   protocol.StoreServiceClient
	Tracer  ot.Tracer
}

// New create new node client
func (c *Client) New() (err error) {

	c.Options = append(c.Options, grpc.WithInsecure())

	// tracer client middleware
	c.Options = append(c.Options, grpc.WithUnaryInterceptor(
		grpc_opentracing.UnaryClientInterceptor(
			grpc_opentracing.WithTracer(c.Tracer))))

	c.Conn, err = grpc.Dial(c.Address, c.Options...)
	if err != nil {
		err = errors.Wrapf(err,
			"Failed to start grpc connection with address %s",
			c.Address)
		return
	}

	c.Store = protocol.NewStoreServiceClient(c.Conn)
	logrus.Info("🥳 dim-edge/core connected to dim-edge-node grpc service at ", c.Address)

	return
}
