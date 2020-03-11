package node

import (
	"dim-edge-core/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// InitEdgeNodeAPI init k8s REST api
func (c *Client) InitEdgeNodeAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/edgenode").Subrouter()
	privateRouter.HandleFunc("/setup", c.handleCheckSetup).Methods("GET")
}

func (c *Client) handleCheckSetup(w http.ResponseWriter, r *http.Request) {
	res, err := c.CheckSetup()
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, res)
	return
}
