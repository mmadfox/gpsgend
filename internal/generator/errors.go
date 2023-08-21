package generator

import (
	"errors"
	"fmt"
)

var (
	ErrTrackerIsAlreadyRunning = errors.New("gpsgend/generator: tracker is already running")
	ErrTrackerIsAlreadyStopped = errors.New("gpsgend/generator: tracker is already stopped")
	ErrTrackerHasNoRoutes      = errors.New("gpsgend/generator: tracker has no routes")
	ErrNoRoutes                = errors.New("gpsgend/generator: no routes for tracker")
	ErrNoSensors               = errors.New("gpsgend/generator: no sensors for tracker")
	ErrTrackNotFound           = errors.New("gpsgend/generator: track not found")
	ErrMaxNumRoutesExceeded    = errors.New("gpsgend/generator: max number of routes exceeded")
	ErrMaxNumTracksExceeded    = errors.New("gpsgend/generator: max number of tracks exceeded")
	ErrMaxNumSegmentsExceeded  = errors.New("gpsgend/generator: max number of segments exceeded")
	ErrMaxNumSensorsExceeded   = errors.New("gpsgend/generator: max number of sensors exceeded")
	ErrTrackerNotRunning       = errors.New("gpsgend/generator: tracker not running")
	ErrTrackerNotPaused        = errors.New("gpsgend/generator: tracker not paused")
	ErrTrackerOff              = errors.New("gpsgend/generator: tracker is off")
	ErrParamsEmpty             = errors.New("gpsgend/generator: params empty")
	ErrInvalidTrackerVersion   = errors.New("gpsgend/generator: invalid tracker version")
	ErrTrackerNotFound         = errors.New("gpsgend/generator: tracker not found")
	ErrRouteNotFound           = errors.New("gpsgend/generator: route not found")
	ErrSensorNotFound          = errors.New("gpsgend/generator: sensor not found")
	ErrLoadingTracker          = errors.New("gpsgend/generator: trackers loading error")
	ErrUnloadingTracker        = errors.New("gpsgend/generator: trackers unloading error")
)

func errTrackerOff(t *Tracker) error {
	return fmt.Errorf("%w - Tracker{ID:%s}", ErrTrackerOff, t.ID())
}
