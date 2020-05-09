package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaceList get namespace list
func (c *Client) GetNamespaceList() (services *v1.NamespaceList, err error) {
	services, err = c.ClientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

	return
}
