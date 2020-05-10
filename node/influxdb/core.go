package influxdb

import (
	"dim-edge/core/protocol"
	"io"
	"time"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"

	"google.golang.org/grpc"
)

// Client dim-edge-node client instance
type Client struct {
	Address     string
	Options     []grpc.DialOption
	Conn        *grpc.ClientConn
	Store       protocol.StoreServiceClient
	Tracer      ot.Tracer
	TraceCloser io.Closer
}

// New create new node client
func (c *Client) New() (err error) {

	// init a new tracer
	jcfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	c.Tracer, c.TraceCloser, err = jcfg.New(
		"dim-edge-core",
		jaegercfg.Logger(jaeger.StdLogger),
	)
	if err != nil {
		return
	}

	ot.SetGlobalTracer(c.Tracer)
	defer c.TraceCloser.Close()

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
