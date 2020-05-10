package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetServiceList get running service list
func (c *Client) GetServiceList(namespace string, matchLabels map[string]string) (services *v1.ServiceList, err error) {

	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	services, err = c.ClientSet.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetSingleService get single service by name
func (c *Client) GetSingleService(namespace string, name string) (service *v1.Service, err error) {
	service, err = c.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	return
}
