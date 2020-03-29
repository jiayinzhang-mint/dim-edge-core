package influxdb

import (
	"dim-edge-core/protocol"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSignIn(*testing.T) {
	c := &Client{
		Address: "127.0.0.1:9090",
	}
	if err := c.New(); err != nil {
		logrus.Error(err)
	}

	err := c.SignIn(&protocol.SignInParams{
		Username: "mint",
		Password: "131001250115zHzH",
	})
	if err != nil {
		logrus.Error(err)
	}
}
