package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetPodList get running pod list
func (c *Client) GetPodList(namespace string, matchLabels map[string]string) (pods *v1.PodList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	pods, err = c.ClientSet.CoreV1().Pods(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetSinglePod get single pod by name
func (c *Client) GetSinglePod(namespace string, name string) (pod *v1.Pod, err error) {
	pod, err = c.ClientSet.CoreV1().Pods(namespace).Get(name, metav1.GetOptions{})

	return
}
