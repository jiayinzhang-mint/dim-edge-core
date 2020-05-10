package k8s

import (
	"flag"
	"io"
	"path/filepath"

	ot "github.com/opentracing/opentracing-go"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

// Client k8s client instance
type Client struct {
	Path             string `json:"path"`
	ClientSet        *kubernetes.Clientset
	MetricsClientSet *metricsv.Clientset
	Tracer           ot.Tracer
	TraceCloser      io.Closer
}

// ConnectToInstance connect to k8s
func (c *Client) ConnectToInstance() (err error) {
	var kubeconfig *string

	kubeconfig = flag.String("kubeconfig", filepath.Join(c.Path, ".kube", "config"), "")

	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return
	}

	// create the clientset
	c.ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	c.MetricsClientSet, err = metricsv.NewForConfig(config)
	if err != nil {
		return
	}

	return
}
