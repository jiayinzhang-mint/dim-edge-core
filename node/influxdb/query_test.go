package influxdb

import (
	"dim-edge/core/protocol"
	"testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
)

func TestInsertData(t *testing.T) {
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

	fields := map[string]float64{"memory": 1000.0, "cpu": 0.93}
	logrus.Info(fields)

	count, err := c.InsertData(
		&protocol.InsertDataParams{
			Org:    "INSDIM",
			Bucket: "INSDIM",
			Metrics: []*protocol.RowMetric{
				&protocol.RowMetric{
					Fields: fields,
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

	logrus.Info("Inserted ", count, "metrics")

}

func TestQueryData(*testing.T) {
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

	res, err := c.QueryData(&protocol.QueryParams{
		QueryString: `from(bucket: "INSDIM")
		|> range(start: -10h)
		|> filter(fn: (r)=>
			r._field == "cpu" and
			r._measurement == "system-metrics" and
			r.hostname == "hal9000"
		)`,
		Org: "INSDIM",
	})

	logrus.Info(res)
}
