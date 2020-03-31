package prometheus

import (
	"context"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// GetAlert returns all active alert
func (c *Client) GetAlert() (a []promv1.Alert, err error) {
	res, err := c.API.Alerts(context.TODO())
	a = res.Alerts
	return
}
