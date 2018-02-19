package model

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/topfreegames/extensions/dogstatsd"
)

// DatadogMonitor implements MonitorService interface.
// Sends an event to Datadog when worker finds a
// Kubernetes controller that has no limit and/or
// request setted
type DatadogMonitor struct {
	client dogstatsd.Client
	region string
}

// NewDatadogMonitor sends metrics to StatsD
func NewDatadogMonitor(config *viper.Viper) (*DatadogMonitor, error) {
	host := config.GetString("monitors.datadog.host")
	prefix := config.GetString("monitors.datadog.prefix")

	ddClient, err := dogstatsd.New(host, prefix)
	if err != nil {
		return nil, err
	}

	return &DatadogMonitor{
		client: ddClient,
		region: config.GetString("monitors.datadog.region"),
	}, nil
}

// Send sends controllers names to Datadog's StatsD
func (d *DatadogMonitor) Send(controllers ...string) error {
	gauge := float64(1)
	for _, controller := range controllers {
		d.client.Gauge(
			"resources-check",
			gauge,
			d.tags(controller),
			1,
		)
	}
	return nil
}

func (d *DatadogMonitor) tags(controller string) []string {
	return []string{
		fmt.Sprintf("unset:%s", controller),
		fmt.Sprintf("region:%s", d.region),
	}
}

// Name returns this monitor's name
func (d *DatadogMonitor) Name() string {
	return "Datadog"
}
