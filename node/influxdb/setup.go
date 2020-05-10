package influxdb

import (
	"context"
	"dim-edge/core/protocol"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
)

// CheckSetup check if influx db has been setup
func (c *Client) CheckSetup() (setup *protocol.CheckSetupRes, err error) {
	span := opentracing.StartSpan("/Store/CheckSetup")
	defer c.GRPCTraceCloser.Close()
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	setup, err = c.Store.CheckSetup(ctx, &empty.Empty{})

	return
}

// Setup setup influxdb
func (c *Client) Setup(p *protocol.SetupParams) (err error) {
	span := opentracing.StartSpan("/Store/Setup")
	defer c.GRPCTraceCloser.Close()
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	_, err = c.Store.Setup(ctx, p)

	return
}
