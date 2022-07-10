package interfaces

import (
	mdl "github.com/seblkma/ieq/models"
)

// Device represents a device containing sensors to provide measurements.
// Authentication should be done by the concrete implementation.
type Device interface {
	GetState(id string) (result string, err error)
	GetDeviceInfo(id string) (result mdl.DeviceInfo, err error)
	GetRawMetrics(id string) (result string, err error)
	GetLatestMetrics(deviceID string) (result mdl.Metrics, err error)
}
