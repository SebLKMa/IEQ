package models

import "time"

// DeviceInfo holds the device information
type DeviceInfo struct {
	DeviceID          string
	DisplayName       string
	MacAddress        string
	SerialNumber      string
	VendorDeviceID    string
	VendorName        string
	Org               string
	CreatedOn         time.Time
	Status            int
	StatusDescription string
}
