package prometheus

import (
	"context"
	"dim-edge-core/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// InitPrometheusAPI init prometheus REST api
func (c *Client) InitPrometheusAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/prometheus").Subrouter()
	privateRouter.HandleFunc("/metrics", c.handleQueryMetrics).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}

func (c *Client) handleQueryMetrics(w http.ResponseWriter, r *http.Request) {

	var (
		query = r.URL.Query().Get("query")

		end     = r.URL.Query().Get("end")
		endTime = time.Now()

		duration     = r.URL.Query().Get("duration")
		durationTime = time.Minute

		step     = r.URL.Query().Get("step")
		stepTime = time.Second

		err error
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
		utils.RespondWithError(w, r, 500, "params err")
		return
	}

	// get time range
	timeRange := promv1.Range{
		Start: time.Now().Add(-durationTime),
		End:   endTime,
		Step:  stepTime,
	}

	// do query
	result, warnings, err := c.API.QueryRange(context.TODO(),
		query,
		timeRange)

	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":      "success",
		"result":   result,
		"warnings": warnings,
	})

	return
}
