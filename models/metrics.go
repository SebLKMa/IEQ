package models

import "time"

// Metrics data for a device at a point in time
type Metrics struct {
	DeviceID    string
	CreatedOn   time.Time
	Temperature float64
	Humidity    float64
	CO2         float64
	VOC         float64
	PM25        float64
	Lighting    float64
	Noise       float64
	Empty       bool
}

// MetricScore data for a device at a point in time
type MetricScore struct {
	DeviceID    string
	CreatedOn   time.Time
	Temperature float64
	Humidity    float64
	CO2         float64
	VOC         float64
	PM25        float64
	Lighting    float64
	Noise       float64
}

// IeqScore data for a device at a point in time
type IeqScore struct {
	DeviceID          string
	CreatedOn         time.Time
	Scheme            string
	Thermal           float64
	ThermalWeighting  float64
	IAQ               float64
	IAQWeighting      float64
	Lighting          float64
	LightingWeighting float64
	Noise             float64
	NoiseWeighting    float64
	Overall           float64
}
