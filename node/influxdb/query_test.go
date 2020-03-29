package influxdb

import (
	"dim-edge-core/protocol"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
)

func TestQueryData(t *testing.T) {
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

	_, err = c.InsertData(
		&protocol.InsertDataParams{
			Org:    "insdim",
			Bucket: "insdim",
			Metrics: []*protocol.RowMetric{
				&protocol.RowMetric{
					Fields: map[string]float64{"memory": 1000.0, "cpu": 0.93},
					Name:   "system-metrics",
					Tags:   map[string]string{"hostname": "hal9000"},
					Ts:     ptypes.TimestampNow(),
				},
			},
		},
	)
	if err != nil {
		logrus.Error(err)
		t.Error(err)
	}

}
