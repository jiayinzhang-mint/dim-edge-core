package prometheus

import (
	"context"
	"dim-edge-core/utils"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prommodel "github.com/prometheus/common/model"
)

// QueryRange get metrics
func (c *Client) QueryRange(query string, end string, duration string, step string) (result prommodel.Value, warnings promv1.Warnings, err error) {
	// get time range
	timeRange, err := utils.GetTimeRange(end, duration, step)
	if err != nil {
		return
	}

	// do query
	if result, warnings, err = c.API.QueryRange(context.TODO(), query, timeRange); err != nil {
		return
	}

	return
}
