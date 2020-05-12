package influxdb

import (
	"context"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func TestCheckSetup(*testing.T) {
	var err error

	c := &Client{
		Address: "127.0.0.1:9090",
	}
	if err = c.New(); err != nil {
		logrus.Error(err)
	}

	span := c.Tracer.StartSpan("/Store/CheckSetup")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	res, err := c.Store.CheckSetup(ctx, &empty.Empty{})
	c.Conn.Close()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(res.Setup)
}
