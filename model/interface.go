package model

// MonitorService sends events to a monitor service or any API
type MonitorService interface {
	Send(...string) error
}
