package grpc

import (
	"time"

	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/go-gpsgen/types"
)

type AddTrackerOptions struct {
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

func NewAddTrackerOptions() *AddTrackerOptions {
	opts := AddTrackerOptions{}
	opts.Speed.Min = 1
	opts.Speed.Max = 6
	opts.Speed.Amplitude = 4
	opts.Elevation.Amplitude = 8
	opts.Elevation.Max = 300
	opts.Elevation.Mode = types.WithSensorRandomMode | types.WithSensorStartMode | types.WithSensorEndMode
	opts.Battery.Max = 100
	opts.Battery.ChargeTime = 7 * time.Hour
	opts.Offline.Min = 5
	opts.Offline.Max = 60
	return &opts
}

type UpdateTrackerOptions struct {
	Model    string
	Color    string
	CustomID string
	Descr    string
}

func (o *UpdateTrackerOptions) isEmpty() bool {
	return len(o.Model) == 0 &&
		len(o.Color) == 0 &&
		len(o.CustomID) == 0 &&
		len(o.Descr) == 0
}
