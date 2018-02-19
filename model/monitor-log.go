package model

import "github.com/sirupsen/logrus"

// LogMonitor just logs whenever a unset controller is found
type LogMonitor struct {
	logger logrus.FieldLogger
}

// NewLogMonitor returns a new LogMonitor
func NewLogMonitor(logger logrus.FieldLogger) *LogMonitor {
	return &LogMonitor{
		logger: logger.WithField("source", "LogMonitor"),
	}
}

// Send just logs the unset controllers
func (l *LogMonitor) Send(controllers ...string) error {
	l.logger.WithField("message", "failed controllers").Info(controllers)
	return nil
}

// Name returns this monitor's name
func (l *LogMonitor) Name() string {
	return "Log"
}
