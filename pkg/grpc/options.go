package grpc

import (
	"time"

	"github.com/mmadfox/go-gpsgen/properties"
)

type NewTrackerOptions struct {
	Model       string // Model of the device.
	Color       string // Color of the device.
	CustomID    string // Custom ID associated with the device.
	Descr       string // Description of the device.
	Props       properties.Properties
	SkipOffline bool // Skip offline mode.
	Offline     struct {
		Min int
		Max int
	}
	Elevation struct {
		Min       float64
		Max       float64
		Amplitude int
		Mode      int
	}
	Battery struct {
		Min        float64
		Max        float64
		ChargeTime time.Duration
	}
	Speed struct {
		Min       float64
		Max       float64
		Amplitude int
	}
}
