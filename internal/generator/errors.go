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
	ErrRouteNotFound           = errors.New("gpsgend/generator: route not found")
	ErrTrackNotFound           = errors.New("gpsgend/generator: track not found")
	ErrMaxNumRoutesExceeded    = errors.New("gpsgend/generator: max number of routes exceeded")
	ErrMaxNumTracksExceeded    = errors.New("gpsgend/generator: max number of tracks exceeded")
	ErrMaxNumSegmentsExceeded  = errors.New("gpsgend/generator: max number of segments exceeded")
	ErrTrackerNotRunning       = errors.New("gpsgend/generator: tracker not running")
	ErrTrackerNotPaused        = errors.New("gpsgend/generator: tracker not paused")
	ErrTrackerOff              = errors.New("gpsgend/generator: tracker is off")
	ErrParamsEmpty             = errors.New("gpsgend/generator: params empty")
	ErrInvalidTrackerVersion   = errors.New("gpsgend/generator: invalid tracker version")
	ErrTrackerNotFound         = errors.New("gpsgend/generator: tracker not found")
)

func errTrackerOff(t *Tracker) error {
	return fmt.Errorf("%w - Tracker{ID:%s}", ErrTrackerOff, t.ID())
}
