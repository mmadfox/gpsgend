package grpc

import "errors"

var (
	ErrInvalidTrackerID    = errors.New("gpsgend/grpc: invalid tracker id")
	ErrNoRoutes            = errors.New("gpsgend/grpc: no routes")
	ErrNoSensor            = errors.New("gpsgend/grpc: no sensor")
	ErrNoDataForUpdate     = errors.New("gpsgend/grpc: no data for update")
	ErrInvalidRouteIndex   = errors.New("gpsgend/grpc: invalid route index")
	ErrInvalidTrackIndex   = errors.New("gpsgend/grpc: invalid track index")
	ErrInvalidSegmentIndex = errors.New("gpsgend/grpc: invalid segment index")
)
