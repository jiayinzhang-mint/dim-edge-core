package prometheus

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/prometheus/common/model"
	"github.com/sirupsen/logrus"
)

func TestQuery(*testing.T) {
	c := Client{
		Address: "http://192.168.64.18:30090",
	}

	if err := c.ConnectToInstance(); err != nil {
		logrus.Error(err)
	}

	res, _, err := c.QueryRange(`
		sum(
			delta(
				container_cpu_usage_seconds_total{
					container="dim-edge-node",
					pod="dim-edge-node-6446bb4f59-bbr9s",service = "dim-edge-mon-prometheus-op-kubelet"
				}[1m]
			)
		) 
		/
		(
			sum(
				container_spec_cpu_quota{
					container="dim-edge-node",pod="dim-edge-node-6446bb4f59-bbr9s",service = "dim-edge-mon-prometheus-op-kubelet"
				}
			) /100000*60 
		)`,
		"",
		"1h",
		"",
	)
	if err != nil {
		logrus.Error(err)
	}

	arr := [][]string{
		{"ds", "y"},
	}
	for _, v := range res.(model.Matrix) {
		for _, col := range v.Values {
			value := col.Value.String()
			ts := col.Timestamp.Time().Unix()

			tsString := strconv.Itoa(int(ts))

			arr = append(arr, []string{tsString, value})
		}
	}

	file, err := os.Create("res.csv")
	if err != nil {
		logrus.Error(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range arr {
		if err := writer.Write(value); err != nil {
			logrus.Error(err)
		}
	}

	logrus.Infof(`%#v`, arr)
}
