package main

import (
	"dim-edge-core/k8s"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func connectToK8S(c *k8s.Client) (err error) {
	if err = c.ConnectToInstance(); err != nil {
		logrus.Error("dim-edge core failed to connect to k8s", err)
		return
	}

	logrus.Info("dim-edge core connected to k8s")
	return
}

func handleRequests(c *k8s.Client) (err error) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "dim-edge core REST service listening")
	}).Methods("GET")

	c.InitK8SAPI(router)

	if err = http.ListenAndServe(":5000", router); err != nil {
		logrus.Error("dim-edge core HTTP service failed to start", err)
		return
	}

	logrus.Info("dim-edge core HTTP service started")
	return
}

var (
	g errgroup.Group
)

func main() {
	logrus.Info("dim-edge core service starting")

	// create k8s client
	c := &k8s.Client{
		Path: homeDir(),
	}

	var (
		err error
	)

	// connect to k8s instance
	err = connectToK8S(c)
	if err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}

	// start http service
	err = handleRequests(c)
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
