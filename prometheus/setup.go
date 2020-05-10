package prometheus

import (
	"io"
	"os"
	"time"

	ot "github.com/opentracing/opentracing-go"
	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// Client prometheus client
type Client struct {
	Address     string `json:"address"`
	Prometheus  promapi.Client
	API         promv1.API
	Tracer      ot.Tracer
	TraceCloser io.Closer
}

// ConnectToInstance connect to prometheus service
func (c *Client) ConnectToInstance() (err error) {
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

	// create client
	c.Prometheus, err = promapi.NewClient(promapi.Config{
		Address: c.Address,
	})
	if err != nil {
		logrus.Error("💣 error creating prometheus client: ", err)
		os.Exit(1)
	}

	logrus.Info("🥳 dim-edge/core connected to prometheus service at ", c.Address)

	// registor api
	c.API = promv1.NewAPI(c.Prometheus)

	return
}
