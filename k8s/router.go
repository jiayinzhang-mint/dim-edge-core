package k8s

import (
	"dim-edge-core/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	appv1 "k8s.io/api/apps/v1"
	scalev1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/api/core/v1"
	metricsv1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// InitK8SAPI init k8s REST api
func (c *Client) InitK8SAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/k8s").Subrouter()

	privateRouter.HandleFunc("/statefulset/list", c.handleGetStatefulSetList).Methods("GET")
	privateRouter.HandleFunc("/statefulset", c.handleGetOneStatefulSet).Methods("GET")
	privateRouter.HandleFunc("/statefulset", c.handleCreateStatefulSet).Methods("POST")
	privateRouter.HandleFunc("/statefulset/scale", c.handleUpdateStatefulSetScale).Methods("PUT")

	privateRouter.HandleFunc("/node/list", c.handleGetNodeList).Methods("GET")
	privateRouter.HandleFunc("/node", c.handleGetOneNode).Methods("GET")
	privateRouter.HandleFunc("/node/metrics/list", c.handleGetNodeMetricsList).Methods("GET")
	privateRouter.HandleFunc("/node/metrics", c.handleGetOneNodeMetrics).Methods("GET")

	privateRouter.HandleFunc("/service/list", c.handleGetServiceList).Methods("GET")
	privateRouter.HandleFunc("/service", c.handleGetSingleService).Methods("GET")

	privateRouter.HandleFunc("/pod/list", c.handleGetPodList).Methods("GET")
	privateRouter.HandleFunc("/pod", c.handleGetSinglePod).Methods("GET")
	privateRouter.HandleFunc("/pod/metrics/list", c.handleGetPodMetricsList).Methods("GET")
	privateRouter.HandleFunc("/pod/metrics", c.handleGetOnePodMetrics).Methods("GET")

	privateRouter.HandleFunc("/volume/claim/list", c.handleGetVolumeClaimList).Methods("GET")
	privateRouter.HandleFunc("/volume/claim", c.handleGetOneVolumeClaim).Methods("GET")
	privateRouter.HandleFunc("/volume/list", c.handleGetVolumeList).Methods("GET")
	privateRouter.HandleFunc("/volume", c.handleGetOneVolume).Methods("GET")

	privateRouter.HandleFunc("/namespace/list", c.handleGetNamespaceList).Methods("GET")

	// privateRouter.Use(auth.CheckAuth)
}

func (c *Client) handleGetStatefulSetList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *appv1.StatefulSetList
		err error
	)

	namespace := r.URL.Query().Get("namespace")
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetStatefulSetList(namespace, matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOneStatefulSet(w http.ResponseWriter, r *http.Request) {
	var (
		s   *appv1.StatefulSet
		err error
	)

	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")
	if s, err = c.GetOneStatefulSet(namespace, name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleCreateStatefulSet(w http.ResponseWriter, r *http.Request) {
	var (
		s              *appv1.StatefulSet
		newStatefulSet *appv1.StatefulSet
		err            error
	)

	if err = json.NewDecoder(r.Body).Decode(&s); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	newStatefulSet, err = c.CreateStatefulSet(s.Namespace, s)
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, newStatefulSet)
	return
}

func (c *Client) handleUpdateStatefulSetScale(w http.ResponseWriter, r *http.Request) {
	var (
		s        *scalev1.Scale
		newScale *scalev1.Scale
		err      error
	)

	if err = json.NewDecoder(r.Body).Decode(&s); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	newScale, err = c.UpdateStatefulSetScale(s.Namespace, s.Name, s.Spec.Replicas)
	if err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}

	utils.RespondWithJSON(w, r, 200, newScale)
	return
}

func (c *Client) handleGetNodeList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.NodeList
		err error
	)

	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetNodeList(matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOneNode(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.Node
		err error
	)

	name := r.URL.Query().Get("name")
	if s, err = c.GetOneNode(name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetNodeMetricsList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *metricsv1.NodeMetricsList
		err error
	)

	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetNodeMetricsList(matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOneNodeMetrics(w http.ResponseWriter, r *http.Request) {
	var (
		s   *metricsv1.NodeMetrics
		err error
	)

	name := r.URL.Query().Get("name")

	if s, err = c.GetOneNodeMetrics(name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetServiceList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.ServiceList
		err error
	)
	namespace := r.URL.Query().Get("namespace")
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetServiceList(namespace, matchLabels); err != nil {
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

func (c *Client) handleGetPodList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.PodList
		err error
	)
	namespace := r.URL.Query().Get("namespace")
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetPodList(namespace, matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetSinglePod(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.Pod
		err error
	)
	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")

	if s, err = c.GetSinglePod(namespace, name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetPodMetricsList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *metricsv1.PodMetricsList
		err error
	)

	namespace := r.URL.Query().Get("namespace")
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetPodMetricsList(namespace, matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOnePodMetrics(w http.ResponseWriter, r *http.Request) {
	var (
		s   *metricsv1.PodMetrics
		err error
	)

	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")

	if s, err = c.GetOnePodMetrics(namespace, name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetVolumeClaimList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.PersistentVolumeClaimList
		err error
	)
	namespace := r.URL.Query().Get("namespace")
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetVolumeClaimList(namespace, matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOneVolumeClaim(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.PersistentVolumeClaim
		err error
	)

	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")
	if s, err = c.GetOneVolumeClaim(namespace, name); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetVolumeList(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.PersistentVolumeList
		err error
	)
	matchLabels := make(map[string]string)
	if serviceName := r.URL.Query().Get("serviceName"); serviceName != "" {
		matchLabels["io.kompose.service"] = serviceName
	}

	if s, err = c.GetVolumeList(matchLabels); err != nil {
		utils.RespondWithError(w, r, 500, err.Error())
		return
	}
	utils.RespondWithJSON(w, r, 200, s)
	return
}

func (c *Client) handleGetOneVolume(w http.ResponseWriter, r *http.Request) {
	var (
		s   *v1.PersistentVolume
		err error
	)

	name := r.URL.Query().Get("name")
	if s, err = c.GetOneVolume(name); err != nil {
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
