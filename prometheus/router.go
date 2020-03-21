package prometheus

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitPrometheusAPI init prometheus REST api
func (c *Client) InitPrometheusAPI(r *mux.Router) {
	http.Handle("/metrics", promhttp.Handler())

}
