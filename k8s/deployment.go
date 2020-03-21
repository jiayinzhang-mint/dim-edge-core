package k8s

import (
	appv1 "k8s.io/api/apps/v1"
	scalev1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetDeploymentList get running replicaset list
func (c *Client) GetDeploymentList(namespace string, matchLabels map[string]string) (rs *appv1.DeploymentList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	rs, err = c.ClientSet.AppsV1().Deployments(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneDeployment get one running replicaset
func (c *Client) GetOneDeployment(namespace string, name string) (rs *appv1.Deployment, err error) {
	rs, err = c.ClientSet.AppsV1().Deployments(namespace).Get(name, metav1.GetOptions{})

	return
}

// UpdateDeploymentScale update one replicaset's scale
func (c *Client) UpdateDeploymentScale(namespace string, name string, replicas int32) (scale *scalev1.Scale, err error) {
	scale, err = c.ClientSet.AppsV1().Deployments(namespace).UpdateScale(name, &scalev1.Scale{
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
