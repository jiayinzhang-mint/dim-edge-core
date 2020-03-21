package k8s

import (
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetReplicaSetList get running replicaset list
func (c *Client) GetReplicaSetList(namespace string, matchLabels map[string]string) (rs *appv1.ReplicaSetList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	rs, err = c.ClientSet.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneReplicaSet get one running replicaset
func (c *Client) GetOneReplicaSet(namespace string, name string) (rs *appv1.ReplicaSet, err error) {
	rs, err = c.ClientSet.AppsV1().ReplicaSets(namespace).Get(name, metav1.GetOptions{})

	return
}
