package influxdb

import (
	"dim-edge-core/protocol"
	"dim-edge-core/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// InitEdgeNodeAPI init k8s REST api
func (c *Client) InitEdgeNodeAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/edgenode").Subrouter()
	privateRouter.HandleFunc("/influxdb/setup", c.handleCheckSetup).Methods("GET")
	privateRouter.HandleFunc("/influxdb/setup", c.handleSetup).Methods("POST")

	privateRouter.HandleFunc("/influxdb/signin", c.handleSignIn).Methods("POST")
	privateRouter.HandleFunc("/influxdb/signout", c.handleSignOut).Methods("POST")

	privateRouter.HandleFunc("/influxdb/authorization", c.handleGetAuthorization).Methods("GET")

	privateRouter.HandleFunc("/influxdb/bucket/list", c.handleListAllBuckets).Methods("GET")
	privateRouter.HandleFunc("/influxdb/bucket", c.handleRetreiveBucket).Methods("GET")
}

func (c *Client) handleCheckSetup(w http.ResponseWriter, r *http.Request) {
	res, err := c.CheckSetup()
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":   "success",
		"setup": res.Setup,
	})
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

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
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
		utils.RespondWithError(w, r, 401, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":           "success",
		"authorization": a,
	})
	return
}

func (c *Client) handleSignIn(w http.ResponseWriter, r *http.Request) {
	var (
		p   *protocol.SignInParams
		err error
	)

	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	if err = c.SignIn(p); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg": "success",
	})
	return
}

func (c *Client) handleSignOut(w http.ResponseWriter, r *http.Request) {
	if err := c.SignOut(); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg": "success",
	})
	return
}

func (c *Client) handleListAllBuckets(w http.ResponseWriter, r *http.Request) {
	pageInt, _ := strconv.Atoi(r.URL.Query().Get("page"))
	sizeInt, _ := strconv.Atoi(r.URL.Query().Get("size"))
	var (
		p = &protocol.ListAllBucketsParams{
			Page:  int32(pageInt),
			Size:  int32(sizeInt),
			Org:   r.URL.Query().Get("org"),
			OrgID: r.URL.Query().Get("orgID"),
			Name:  r.URL.Query().Get("name"),
		}
		b   []*protocol.Bucket
		err error
	)

	b, err = c.ListAllBuckets(p)
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":    "success",
		"bucket": b,
	})
	return
}

func (c *Client) handleRetreiveBucket(w http.ResponseWriter, r *http.Request) {
	var (
		p = &protocol.RetrieveBucketParams{
			BucketID: r.URL.Query().Get("bucketID"),
		}
		b   *protocol.Bucket
		err error
	)

	b, err = c.RetrieveBucket(p)
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, map[string]interface{}{
		"msg":    "success",
		"bucket": b,
	})
	return
}
