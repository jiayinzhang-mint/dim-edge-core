package node

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
)

func TestCheckSetup(*testing.T) {
	c := &Client{
		Address: "192.168.64.14:30762",
	}
	if err := c.New(); err != nil {
		logrus.Error(err)
	}

	res, err := c.Store.CheckSetup(context.TODO(), &empty.Empty{})
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(res)
}
