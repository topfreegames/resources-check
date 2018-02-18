package model

// MonitorService sends events to a monitor service or any API
type MonitorService interface {
	Send(...string) error
	Name() string
}

// KubernetesController represents a kubernetes controller
// like deployment, daemonset and statefulset
type KubernetesController interface {
	GetName() string
	GetNamespace() string
}
