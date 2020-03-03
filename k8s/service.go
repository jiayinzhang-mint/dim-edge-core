package k8s

import (
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetServiceList get running service list
func (c *Client) GetServiceList() (services *v1.ServiceList, err error) {
	services, err = c.ClientSet.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		return
	}
	logrus.Infof("There are %d services in the cluster\n", len(services.Items))

	return
}
