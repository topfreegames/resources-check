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
		has := hasCPUAndMemory(container.Resources.Limits)
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
	key := fmt.Sprintf("excluding.%s.%s.%s", controllerType, controller.GetNamespace(), controller.GetName())
	return config.GetBool(key)
}

func hasCPUAndMemory(resource v1.ResourceList) bool {
	return len(resource) >= 2
}
