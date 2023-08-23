package codes

import (
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

type code struct {
	Err  error
	Code int
}

var table = []code{
	{Err: types.ErrInvalidMinValue, Code: 1},
	{Err: types.ErrInvalidMaxValue, Code: 2},
	{Err: types.ErrInvalidRangeValue, Code: 3},
	{Err: types.ErrInvalidMinChargeTime, Code: 4},
	{Err: types.ErrInvalidMaxChargeTime, Code: 5},
	{Err: types.ErrInvalidMinAmplitude, Code: 6},
	{Err: types.ErrInvalidMaxAmplitude, Code: 7},
	{Err: types.ErrInvalidName, Code: 8},
	{Err: types.ErrInvalidID, Code: 9},

	{Err: generator.ErrParamsEmpty, Code: 10},
	{Err: generator.ErrTrackerIsAlreadyRunning, Code: 11},
	{Err: generator.ErrTrackerIsAlreadyStopped, Code: 12},
	{Err: generator.ErrTrackerHasNoRoutes, Code: 13},
	{Err: generator.ErrNoRoutes, Code: 14},
	{Err: generator.ErrNoSensor, Code: 15},
	{Err: generator.ErrMaxNumRoutesExceeded, Code: 16},
	{Err: generator.ErrMaxNumTracksExceeded, Code: 17},
	{Err: generator.ErrMaxNumSegmentsExceeded, Code: 18},
	{Err: generator.ErrMaxNumSensorsExceeded, Code: 19},
	{Err: generator.ErrTrackerNotRunning, Code: 20},
	{Err: generator.ErrTrackerNotPaused, Code: 21},
	{Err: generator.ErrTrackerOff, Code: 22},
	{Err: generator.ErrTrackNotFound, Code: 23},
	{Err: generator.ErrTrackerNotFound, Code: 24},
	{Err: generator.ErrRouteNotFound, Code: 25},
	{Err: generator.ErrSensorNotFound, Code: 26},
	{Err: generator.ErrInvalidTrackerVersion, Code: 27},
	{Err: generator.ErrLoadingTracker, Code: 28},
	{Err: generator.ErrUnloadingTracker, Code: 29},
}
