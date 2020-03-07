package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	metricsv1 "k8s.io/metrics/pkg/apis/metrics/v1beta1"
)

// GetPodMetrics get pod metrics
func (c *Client) GetPodMetrics() (p *metricsv1.PodMetricsList, err error) {

	p, err = c.MetricsClientSet.MetricsV1beta1().PodMetricses("").List(metav1.ListOptions{})

	return
}
