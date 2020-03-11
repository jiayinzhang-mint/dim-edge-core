package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

// GetVolumeClaimList get persistent volume claims
func (c *Client) GetVolumeClaimList(namespace string, matchLabels map[string]string) (v *v1.PersistentVolumeClaimList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	v, err = c.ClientSet.CoreV1().PersistentVolumeClaims(namespace).List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneVolumeClaim get one persistent volume claim
func (c *Client) GetOneVolumeClaim(namespace string, name string) (v *v1.PersistentVolumeClaim, err error) {

	v, err = c.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Get(name, metav1.GetOptions{})

	return
}

// GetVolumeList get persistent volumes
func (c *Client) GetVolumeList(matchLabels map[string]string) (v *v1.PersistentVolumeList, err error) {
	labelSelector := metav1.LabelSelector{MatchLabels: matchLabels}

	v, err = c.ClientSet.CoreV1().PersistentVolumes().List(metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	return
}

// GetOneVolume get one persistent volume
func (c *Client) GetOneVolume(name string) (v *v1.PersistentVolume, err error) {

	v, err = c.ClientSet.CoreV1().PersistentVolumes().Get(name, metav1.GetOptions{})

	return
}
