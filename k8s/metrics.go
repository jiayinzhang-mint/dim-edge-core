package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	metricsv1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// GetPodMetricsList get pod metrics
func (c *Client) GetPodMetricsList(namespace string, matchLabels map[string]string) (p *metricsv1.PodMetricsList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}
	p, err = c.MetricsClientSet.MetricsV1beta1().PodMetricses(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})
	return
}

// GetOnePodMetrics get pod metrics
func (c *Client) GetOnePodMetrics(namespace string, name string) (p *metricsv1.PodMetrics, err error) {
	p, err = c.MetricsClientSet.MetricsV1beta1().PodMetricses(namespace).Get(name, metav1.GetOptions{})
	return
}
