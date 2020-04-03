package prometheus

import (
	"context"
	"dim-edge-core/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	prommodel "github.com/prometheus/common/model"
)

// InitPrometheusAPI init prometheus REST api
func (c *Client) InitPrometheusAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/prometheus").Subrouter()
	privateRouter.HandleFunc("/query", c.handleQuery).Methods("GET")
	privateRouter.HandleFunc("/query/range", c.handleQueryRange).Methods("GET")

	privateRouter.HandleFunc("/alert", c.handleGetAlert).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}

func (c *Client) handleQueryRange(w http.ResponseWriter, r *http.Request) {
	var (
		query    = r.URL.Query().Get("query")
		end      = r.URL.Query().Get("end")
		duration = r.URL.Query().Get("duration")
		step     = r.URL.Query().Get("step")

		result   prommodel.Value
		warnings promv1.Warnings
		err      error
	)

	if result, warnings, err = c.QueryRange(query, end, duration, step); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
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
		query    = r.URL.Query().Get("query")
		result   prommodel.Value
		warnings promv1.Warnings
		err      error
	)

	if result, warnings, err = c.API.Query(context.TODO(), query, time.Now()); err != nil {
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

func (c *Client) handleGetAlert(w http.ResponseWriter, r *http.Request) {
	var (
		a   []promv1.Alert
		err error
	)

	if a, err = c.GetAlert(); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":   "success",
		"alert": a,
	})
	return
}
