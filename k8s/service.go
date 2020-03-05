package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetServiceList get running service list
func (c *Client) GetServiceList(namespace string) (services *v1.ServiceList, err error) {
	services, err = c.ClientSet.CoreV1().Services(namespace).List(metav1.ListOptions{})

	return
}

// GetSingleService get single service by name
func (c *Client) GetSingleService(namespace string, name string) (service *v1.Service, err error) {
	service, err = c.ClientSet.CoreV1().Services(namespace).Get(name, metav1.GetOptions{})

	return
}
