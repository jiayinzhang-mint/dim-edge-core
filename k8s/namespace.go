package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaceList get namespace list
func (c *Client) GetNamespaceList() (services *v1.NamespaceList, err error) {
	services, err = c.ClientSet.CoreV1().Namespaces().List(metav1.ListOptions{})

	return
}
