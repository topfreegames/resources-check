package model

// ControllerType names the Kubernetes controller types
type ControllerType string

// Deployment names the controller Deployment
var Deployment ControllerType = "deployment"

// Statefulset names the controller Statefulset
var Statefulset ControllerType = "statefulset"

// Daemonset names the controller Daemonset
var Daemonset ControllerType = "daemonset"
