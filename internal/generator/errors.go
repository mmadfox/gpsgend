package generator

import "errors"

var (
	ErrTrackerIsAlreadyRunning = errors.New("gpsgend/generator: tracker is already running")
	ErrTrackerIsAlreadyStopped = errors.New("gpsgend/generator: tracker is already stopped")
	ErrTrackerHasNoRoutes      = errors.New("gpsgend/generator: tracker has no routes")
	ErrNoRoutes                = errors.New("gpsgend/generator: no routes for tracker")
	ErrTrackNotFound           = errors.New("gpsgend/generator: track not found")
	ErrMaxNumRoutesExceeded    = errors.New("gpsgend/generator: max number of routes exceeded")
	ErrMaxNumTracksExceeded    = errors.New("gpsgend/generator: max number of tracks exceeded")
	ErrMaxNumSegmentsExceeded  = errors.New("gpsgend/generator: max number of segments exceeded")
	ErrTrackerNotRunning       = errors.New("gpsgend/generator: tracker not running")
)
