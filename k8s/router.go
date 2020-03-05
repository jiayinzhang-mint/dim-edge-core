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
	privateRouter.HandleFunc("/service/list", c.handleGetServiceList).Methods("GET")
	privateRouter.HandleFunc("/service", c.handleGetSingleService).Methods("GET")

	privateRouter.HandleFunc("/namespace/list", c.handleGetNamespaceList).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}

func (c *Client) handleGetServiceList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.ServiceList
		err error
	)
	namespace := r.URL.Query().Get("namespace")

	if s, err = c.GetServiceList(namespace); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetSingleService(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.Service
		err error
	)
	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")

	if s, err = c.GetSingleService(namespace, name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetNamespaceList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.NamespaceList
		err error
	)

	if s, err = c.GetNamespaceList(); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}
