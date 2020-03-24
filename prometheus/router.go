package prometheus

import (
	"context"
	"dim-edge-core/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// InitPrometheusAPI init prometheus REST api
func (c *Client) InitPrometheusAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/prometheus").Subrouter()
	privateRouter.HandleFunc("", c.handleQuery).Methods("GET")
	privateRouter.HandleFunc("/range", c.handleQueryRange).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}

func (c *Client) handleQueryRange(w http.ResponseWriter, r *http.Request) {

	var (
		query    = r.URL.Query().Get("query")
		end      = r.URL.Query().Get("end")
		duration = r.URL.Query().Get("duration")
		step     = r.URL.Query().Get("step")
		err      error
	)

	// get time range
	timeRange, err := utils.GetTimeRange(end, duration, step)
	if err != nil {
		utils.RespondWithError(w, r, 400, err.Error())
		return
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

func (c *Client) handleQuery(w http.ResponseWriter, r *http.Request) {
	var (
		query = r.URL.Query().Get("query")
		// queryTime = r.URL.Query().Get("queryTime")
		err error
	)

	// timeValue, err := time.Parse("2006-01-02T15:04:05.000Z", queryTime)

	// do query
	result, warnings, err := c.API.Query(context.TODO(),
		query,
		time.Now())

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
