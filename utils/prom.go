package utils

import (
	"time"

	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// GetTimeRange convert params to time range
func GetTimeRange(end string, duration string, step string) (r promv1.Range, err error) {
	var (
		endTime      = time.Now()
		durationTime = 10 * time.Minute
		stepTime     = 15 * time.Second
	)

	// parse time
	if end != "" {
		endTime, err = time.Parse("2006-01-02T15:04:05.000Z", end)
	}

	if duration != "" {
		durationTime, err = time.ParseDuration(duration)
	}

	if step != "" {
		stepTime, err = time.ParseDuration(step)
	}

	if err != nil {
		return
	}

	r.End = endTime
	r.Start = endTime.Add(-durationTime)
	r.Step = stepTime

	return
}
