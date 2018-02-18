package model

// DatadogMonitor implements MonitorService interface.
// Sends an event to Datadog when worker finds a
// Kubernetes controller that has no limit and/or
// request setted
type DatadogMonitor struct {
}

// NewDatadogMonitor ...
func NewDatadogMonitor() *DatadogMonitor {
	return &DatadogMonitor{}
}

// Send sends controllers names to Datadog's StatsD
func (d *DatadogMonitor) Send(controllers ...string) error {
	return nil
}

// Name returns this monitor's name
func (d *DatadogMonitor) Name() string {
	return "Datadog"
}
