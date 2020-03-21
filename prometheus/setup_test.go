package prometheus

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/sirupsen/logrus"
)

func TestSetup(*testing.T) {
	c := Client{
		Address: "http://192.168.64.16:30090",
	}

	if err := c.ConnectToInstance(); err != nil {
		logrus.Error(err)
	}

	r := promv1.Range{
		Start: time.Now().Add(-time.Minute),
		End:   time.Now(),
		Step:  time.Minute,
	}

	result, warnings, err := c.API.QueryRange(context.TODO(),
		`container_memory_working_set_bytes{container="dim-edge-influxdb",service="dim-edge-mon-prometheus-op-kubelet"}`,
		r)
	if err != nil {
		fmt.Printf("Error querying Prometheus: %v\n", err)
		os.Exit(1)
	}
	if len(warnings) > 0 {
		fmt.Printf("Warnings: %v\n", warnings)
	}
	fmt.Printf("Result:\n%v\n", result)

}
