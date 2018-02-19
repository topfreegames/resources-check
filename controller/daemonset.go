package controller

import (
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CheckDaemonset checks daemonset on namespace
func CheckDaemonset(
	kubernetesClient kubernetes.Interface,
	config *viper.Viper,
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	daemonsets, err := kubernetesClient.
		Extensions().
		DaemonSets(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedDaemonsets := []string{}
	for _, daemonset := range daemonsets.Items {
		isIgnored := IsIgnored(Daemonset, &daemonset, config)
		hasResources := HasResources(daemonset.Spec.Template.Spec.Containers)

		if !isIgnored && !hasResources {
			failedDaemonsets = append(failedDaemonsets, Name(&daemonset))
		}
	}
	return failedDaemonsets, nil
}
