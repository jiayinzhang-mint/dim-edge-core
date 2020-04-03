package prometheus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestGetAlert(*testing.T) {
	c := Client{
		Address: "http://192.168.64.18:30090",
	}

	if err := c.ConnectToInstance(); err != nil {
		logrus.Error(err)
	}

	a, err := c.GetAlert()
	if err != nil {
		logrus.Error(err)
	}

	logrus.Info(a)
}
