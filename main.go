package main

import (
	"dim-edge/core/auth"
	"dim-edge/core/k8s"
	"dim-edge/core/node/influxdb"
	"dim-edge/core/prometheus"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var minikubeIP = "127.0.0.1"
var nodeGRPCPort = ":9090"

func connectToK8S(c *k8s.Client) (err error) {
	if err = c.ConnectToInstance(); err != nil {
		logrus.Error("💣 dim-edge/core failed to connect to k8s", err)
		return
	}

	logrus.Info("🥳 dim-edge/core connected to k8s minikube service at ", c.Path)
	return
}

func handleRequests(c *k8s.Client, gc *influxdb.Client, pc *prometheus.Client) (err error) {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "dim-edge/core REST service listening")
	}).Methods("GET")

	c.InitK8SAPI(router)
	gc.InitEdgeNodeAPI(router)
	pc.InitPrometheusAPI(router)
	auth.InitAuthAPI(router)

	addr := ":5000"

	logrus.Info("🤣 dim-edge/core HTTP service started at ", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		logrus.Error("💣 dim-edge/core HTTP service failed to start", err)
		return
	}

	return
}

var (
	g errgroup.Group
)

func main() {
	var (
		err error
	)

	logrus.Info("👀 dim-edge/core service starting")

	// create k8s client
	c := &k8s.Client{
		Path: homeDir(),
	}

	// connect to k8s instance
	err = connectToK8S(c)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	// create prometheus client
	pc := &prometheus.Client{
		Address: "http://" + minikubeIP + ":30090",
	}

	err = pc.ConnectToInstance()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	// create edge-node grpc client
	gc := &influxdb.Client{
		Address: minikubeIP + nodeGRPCPort,
	}

	// connect to edge-node grpc instance
	err = gc.New()
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	// start http service
	err = handleRequests(c, gc, pc)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
