package k8s

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetNodeList get running pod list
func (c *Client) GetNodeList(matchLabels map[string]string) (nodes *v1.NodeList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	nodes, err = c.ClientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneNode get running pod list
func (c *Client) GetOneNode(name string) (nodes *v1.Node, err error) {
	nodes, err = c.ClientSet.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})

	return
}
