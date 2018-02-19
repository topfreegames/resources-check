package controller

import (
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CheckDeployments checks deployments on namespace
func CheckDeployments(
	kubernetesClient kubernetes.Interface,
	config *viper.Viper,
	namespace string,
) ([]string, error) {
	listOptions := v1.ListOptions{}
	deployments, err := kubernetesClient.
		Extensions().
		Deployments(namespace).
		List(listOptions)
	if err != nil {
		return nil, err
	}
	failedDeployments := []string{}
	for _, deployment := range deployments.Items {
		isIgnored := IsIgnored(Deployment, &deployment, config)
		hasResources := HasResources(deployment.Spec.Template.Spec.Containers)

		if !isIgnored && !hasResources {
			failedDeployments = append(failedDeployments, Name(&deployment))
		}
	}
	return failedDeployments, nil
}
