package k8s

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func TestGetPodMetrics(*testing.T) {

	// create k8s client
	c := &Client{
		Path: homeDir(),
	}

	err := c.ConnectToInstance()
	logrus.Error(err)

	p, err := c.GetPodMetrics()
	logrus.Info(p, err)
}
