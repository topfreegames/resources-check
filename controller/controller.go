package controller

import (
	"fmt"

	"github.com/spf13/viper"
	"k8s.io/client-go/pkg/api/v1"
)

// Name returns controller's namespace/name
func Name(controller Controller) string {
	return fmt.Sprintf("%s/%s",
		controller.GetNamespace(),
		controller.GetName(),
	)
}

// HasResources returns true if controller has
// request and limit setted
func HasResources(containers []v1.Container) bool {
	for _, container := range containers {
		has := hasCPUAndMemory(container.Resources.Requests) &&
			hasCPUAndMemory(container.Resources.Limits)
		if !has {
			return false
		}
	}
	return true
}

// IsIgnored returns true if
// controller is chosen to be ignored
func IsIgnored(
	controllerType Type,
	controller Controller,
	config *viper.Viper,
) bool {
	key := fmt.Sprintf("app.ignored.%s.%s", controllerType, controller.GetNamespace())
	controllerNames := config.GetStringSlice(key)
	for _, controllerName := range controllerNames {
		if controllerName == controller.GetName() {
			return true
		}
	}
	return false
}

func hasCPUAndMemory(resource v1.ResourceList) bool {
	return len(resource) >= 2
}
