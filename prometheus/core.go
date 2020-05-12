package prometheus

import (
	"os"

	ot "github.com/opentracing/opentracing-go"
	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
)

// Client prometheus client
type Client struct {
	Address    string `json:"address"`
	Prometheus promapi.Client
	API        promv1.API
	Tracer     ot.Tracer
}

// ConnectToInstance connect to prometheus service
func (c *Client) ConnectToInstance() (err error) {

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
