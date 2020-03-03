package k8s

import (
	"dim-edge-core/utils"
	"net/http"

	"github.com/gorilla/mux"
	v1 "k8s.io/api/core/v1"
)

// InitK8SAPI init k8s REST api
func (c *Client) InitK8SAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/k8s").Subrouter()

	// get service list
	privateRouter.HandleFunc("/service/list", func(w http.ResponseWriter, r *http.Request) {
		var (
			s   *v1.ServiceList
			err error
		)
		if s, err = c.GetServiceList(); err != nil {
			utils.RespondWithError(w, r, 500, err.Error())
			return
		}
		utils.RespondWithJSON(w, r, 200, s)
		return
	}).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}
