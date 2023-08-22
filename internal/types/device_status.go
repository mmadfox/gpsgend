package types

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen"
)

type DeviceStatus gpsgen.Status

const (
	Running = DeviceStatus(gpsgen.Running)
	Stopped = DeviceStatus(gpsgen.Stopped)
	Paused  = DeviceStatus(gpsgen.Stopped + 1)
)

func (t DeviceStatus) String() string {
	switch t {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Paused:
		return "Paused"
	default:
		return "Unknown device status"
	}
}

func ParseDeviceStatus(status int) (DeviceStatus, error) {
	s := DeviceStatus(status)
	if err := s.Validate(); err != nil {
		return DeviceStatus(0), err
	}
	return s, nil
}

func (t DeviceStatus) Validate() error {
	switch t {
	case Running, Stopped, Paused:
		return nil
	default:
		return fmt.Errorf("gpsgend/types: invalid device status")
	}
}
