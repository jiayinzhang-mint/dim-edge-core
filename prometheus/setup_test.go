package prometheus

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSetup(*testing.T) {
	c := Client{
		Address: "http://192.168.64.16:30090",
	}

	if err := c.ConnectToInstance(); err != nil {
		logrus.Error(err)
	}

}
