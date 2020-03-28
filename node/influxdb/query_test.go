package influxdb

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestQueryData(*testing.T) {
	c := &Client{
		Address: "192.168.64.16:32532",
	}
	if err := c.New(); err != nil {
		logrus.Error(err)
	}

	// res, err := c.InsertData(
	// 	&protocol.InsertDataParams{
	// 		Org:    "insdim",
	// 		Bucket: "insdim",
	// 		Metrics: []*protocol.RowMetric{
	// 			&protocol.RowMetric{
	// 				Fields: map[string]float64{"memory": 1000, "cpu": 0.93},
	// 			},
	// 		},
	// 	},
	// )

}
