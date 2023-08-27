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
	ErrNoSensor                = errors.New("gpsgend/generator: no sensor for tracker")
	ErrTrackNotFound           = errors.New("gpsgend/generator: track not found")
	ErrMaxNumRoutesExceeded    = errors.New("gpsgend/generator: max number of routes exceeded")
	ErrMaxNumTracksExceeded    = errors.New("gpsgend/generator: max number of tracks exceeded")
	ErrMaxNumSegmentsExceeded  = errors.New("gpsgend/generator: max number of segments exceeded")
	ErrMaxNumSensorsExceeded   = errors.New("gpsgend/generator: max number of sensors exceeded")
	ErrTrackerNotRunning       = errors.New("gpsgend/generator: tracker not running")
	ErrTrackerNotPaused        = errors.New("gpsgend/generator: tracker not paused")
	ErrTrackerOff              = errors.New("gpsgend/generator: tracker is off")
	ErrParamsEmpty             = errors.New("gpsgend/generator: params empty")
	ErrTrackerNotFound         = errors.New("gpsgend/generator: tracker not found")
	ErrRouteNotFound           = errors.New("gpsgend/generator: route not found")
	ErrSensorNotFound          = errors.New("gpsgend/generator: sensor not found")
	ErrLoadingTracker          = errors.New("gpsgend/generator: trackers loading error")
	ErrUnloadingTracker        = errors.New("gpsgend/generator: trackers unloading error")
	ErrSensorAlreadyExists     = errors.New("gpsgned/generator: sensor already exists")
	ErrNoTracker               = errors.New("gpsgend/generator: no tracker")
	ErrInvalidParams           = errors.New("gpsgend/generator: invalid params")

	ErrBrokenTracker         = errors.New("gpsgend/generator: broken tracker data")
	ErrInvalidTrackerVersion = errors.New("gpsgend/generator: invalid tracker version")

	ErrStorageInsert = errors.New("gpsgend/generator: insert into storage")
	ErrStorageUpdate = errors.New("gpsgend/generator: update in storage")
	ErrStorageFind   = errors.New("gpsgned/generator: find in storage")
	ErrStorageDelete = errors.New("gpsgned/generator: delete from storage")
	ErrStorageSearch = errors.New("gpsgned/generator: search in storage")
)

func errTrackerOff(t *Tracker) error {
	return fmt.Errorf("%w - Tracker{ID:%s}", ErrTrackerOff, t.ID())
}
