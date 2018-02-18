package model

import (
	"fmt"

	"k8s.io/client-go/pkg/api/v1"

	"github.com/spf13/viper"
)

func isIgnoredKey(
	controllerType ControllerType,
	name, namespace string,
	config *viper.Viper,
) string {
	return fmt.Sprintf("app.ignored.%s.%s", controllerType, namespace)
}

func isIgnored(
	controllerType ControllerType,
	controller KubernetesController,
	config *viper.Viper,
) bool {
	key := isIgnoredKey(controllerType,
		controller.GetName(), controller.GetNamespace(), config)

	controllerNames := config.GetStringSlice(key)
	for _, controllerName := range controllerNames {
		if controllerName == controller.GetName() {
			return true
		}
	}
	return false
}

func hasResources(containers []v1.Container) bool {
	for _, container := range containers {
		has := hasCPUAndMemory(container.Resources.Requests) &&
			hasCPUAndMemory(container.Resources.Limits)
		if !has {
			return false
		}
	}
	return true
}

func hasCPUAndMemory(resource v1.ResourceList) bool {
	return len(resource) >= 2
}

func name(controller KubernetesController) string {
	return fmt.Sprintf("%s/%s",
		controller.GetNamespace(),
		controller.GetName(),
	)
}
