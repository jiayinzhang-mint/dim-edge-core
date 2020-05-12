package prometheus

import (
	"context"
	"dim-edge/core/utils"

	"github.com/opentracing/opentracing-go"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prommodel "github.com/prometheus/common/model"
)

// QueryRange get metrics
func (c *Client) QueryRange(query string, end string, duration string, step string) (result prommodel.Value, warnings promv1.Warnings, err error) {
	span := opentracing.StartSpan("/Prometheus/QueryRange")
	defer span.Finish()
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// get time range
	timeRange, err := utils.GetTimeRange(end, duration, step)
	if err != nil {
		return
	}

	// do query
	if result, warnings, err = c.API.QueryRange(ctx, query, timeRange); err != nil {
		return
	}

	return
}
