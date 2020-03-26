package influxdb

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewClient(*testing.T) {
	c := &Client{
		Address: "192.168.64.14:30028",
	}

	err := c.New()
	logrus.Info(err, &c.Store)
}
