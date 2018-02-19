package controller

// Controller represents a kubernetes controller
// like deployment, daemonset and statefulset
type Controller interface {
	GetName() string
	GetNamespace() string
}
