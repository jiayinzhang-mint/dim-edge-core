package influxdb

import (
	"dim-edge-core/protocol"
	"dim-edge-core/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// InitEdgeNodeAPI init k8s REST api
func (c *Client) InitEdgeNodeAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/edgenode").Subrouter()
	privateRouter.HandleFunc("/influxdb/setup", c.handleCheckSetup).Methods("GET")
	privateRouter.HandleFunc("/influxdb/setup", c.handleSetup).Methods("POST")

	privateRouter.HandleFunc("/influxdb/authorization", c.handleCheckSetup).Methods("GET")
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

func (c *Client) handleSetup(w http.ResponseWriter, r *http.Request) {
	var (
		p   *protocol.SetupParams
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	if err = c.Setup(p); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]string{
		"msg": "success",
	})
	return
}

func (c *Client) handleGetAuthorization(w http.ResponseWriter, r *http.Request) {
	var (
		p = protocol.ListAuthorizationParams{
			User:   r.URL.Query().Get("user"),
			UserID: r.URL.Query().Get("userID"),
			Org:    r.URL.Query().Get("org"),
			OrgID:  r.URL.Query().Get("orgID"),
		}
		a   []*protocol.Authorization
		err error
	)
	a, err = c.ListAuthorization(&p)
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, a)
	return
}
