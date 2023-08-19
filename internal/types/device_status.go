package types

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen"
)

type DeviceStatus gpsgen.Status

const (
	Running = DeviceStatus(gpsgen.Running)
	Stopped = DeviceStatus(gpsgen.Stopped)
)

func ParseDeviceStatus(status int) (DeviceStatus, error) {
	s := DeviceStatus(status)
	if err := s.Validate(); err != nil {
		return DeviceStatus(0), err
	}
	return s, nil
}

func (t DeviceStatus) Validate() error {
	switch t {
	case Running, Stopped:
		return nil
	default:
		return fmt.Errorf("gpsgend/types: invalid device status")
	}
}
