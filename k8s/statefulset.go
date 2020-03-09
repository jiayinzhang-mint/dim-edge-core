package k8s

import (
	appv1 "k8s.io/api/apps/v1"
	scalev1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetStatefulSetList get running stateful set list
func (c *Client) GetStatefulSetList(namespace string, matchLabels map[string]string) (s *appv1.StatefulSetList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	s, err = c.ClientSet.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneStatefulSet get single pod by name
func (c *Client) GetOneStatefulSet(namespace string, name string) (s *appv1.StatefulSet, err error) {
	s, err = c.ClientSet.AppsV1().StatefulSets(namespace).Get(name, metav1.GetOptions{})

	return
}

// CreateStatefulSet create one statefulset
func (c *Client) CreateStatefulSet(namespace string, ori *appv1.StatefulSet) (s *appv1.StatefulSet, err error) {
	s, err = c.ClientSet.AppsV1().StatefulSets(namespace).Create(ori)
	return
}

// UpdateStatefulSetScale update statefulset scale
func (c *Client) UpdateStatefulSetScale(namespace string, name string, replicas int32) (s *scalev1.Scale, err error) {
	s, err = c.ClientSet.AppsV1().StatefulSets(namespace).UpdateScale(name, &scalev1.Scale{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: scalev1.ScaleSpec{
			Replicas: replicas,
		},
	})
	return
}
