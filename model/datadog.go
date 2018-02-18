package model

// DatadogService implements MonitorService interface.
// Sends an event to Datadog when worker finds a
// Kubernetes controller that has no limit and/or
// request setted
type DatadogService struct {
}
