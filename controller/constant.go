package controller

// Type names the Kubernetes controller types
type Type string

// Deployment names the controller Deployment
var Deployment Type = "deployment"

// Statefulset names the controller Statefulset
var Statefulset Type = "statefulset"

// Daemonset names the controller Daemonset
var Daemonset Type = "daemonset"
