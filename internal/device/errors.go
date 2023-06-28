package device

import "errors"

var (
	ErrSensorLimitExceeded  = errors.New("gpsgend: sensor limit exceeded")
	ErrRoutesLimitExceeded  = errors.New("gpsgend: routes limit exceeded")
	ErrSetConfig            = errors.New("gpsgend: unable to change configuration to running device")
	ErrDeviceAlreadyRunning = errors.New("gpsgend: device is already running")
	ErrDeviceAlreadyStopped = errors.New("gpsgend: device is already stopped")
	ErrNoRoutes             = errors.New("gpsgend: no routes")
	ErrNoSnapshot           = errors.New("gpsgend: no snapshot")
	ErrDeviceNotFound       = errors.New("gpsgen: device not found")
)
