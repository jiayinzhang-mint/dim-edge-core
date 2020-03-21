package prometheus

import (
	"context"
	"fmt"
	"os"
	"time"

	promapi "github.com/prometheus/client_golang/api"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// Client prometheus client
type Client struct {
	Address    string `json:"address"`
	Prometheus promapi.Client
}

// ConnectToInstance connect to prometheus service
func (c *Client) ConnectToInstance() (err error) {
	c.Prometheus, err = promapi.NewClient(promapi.Config{
		Address: c.Address,
	})
	if err != nil {
		fmt.Printf("Error creating client: %v\n", err)
		os.Exit(1)
	}

	v1api := promv1.NewAPI(c.Prometheus)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r := promv1.Range{
		Start: time.Now().Add(-time.Hour),
		End:   time.Now(),
		Step:  time.Minute,
	}
	result, warnings, err := v1api.QueryRange(ctx,
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
	return
}
