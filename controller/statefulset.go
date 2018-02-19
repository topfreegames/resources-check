package controller

import (
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CheckStatefulset checks statefulsets on namespace
func CheckStatefulset(
	kubernetesClient kubernetes.Interface,
	config *viper.Viper,
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	statefulsets, err := kubernetesClient.
		Apps().
		StatefulSets(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedStatefulsets := []string{}
	for _, statefulset := range statefulsets.Items {
		isIgnored := IsIgnored(Statefulset, &statefulset, config)
		hasResources := HasResources(statefulset.Spec.Template.Spec.Containers)

		if !isIgnored && !hasResources {
			failedStatefulsets = append(failedStatefulsets, Name(&statefulset))
		}
	}
	return failedStatefulsets, nil
}
