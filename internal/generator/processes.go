package generator

import (
	gpsgen "github.com/mmadfox/go-gpsgen"
)

type Processes interface {
	HasTracker(deviceID string) bool
	Attach(d *gpsgen.Device) error
	Detach(deviceID string) error
	Lookup(deviceID string) (*gpsgen.Device, bool)
	Each(func(n int, p *gpsgen.Device) bool)
}
