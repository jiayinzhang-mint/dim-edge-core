package node

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewClient(*testing.T) {
	c := &Client{
		Address: "192.168.64.9:31037",
	}

	err := c.New()
	logrus.Info(err, &c.Store)
}
