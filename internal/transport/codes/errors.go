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
	{Err: types.ErrInvalidMinValue, Code: CodeInvalidMinValue},
	{Err: types.ErrInvalidMaxValue, Code: CodeInvalidMaxValue},
	{Err: types.ErrInvalidRangeValue, Code: CodeInvalidRangeValue},
	{Err: types.ErrInvalidMinChargeTime, Code: CodeInvalidMinChargeTime},
	{Err: types.ErrInvalidMaxChargeTime, Code: CodeInvalidMaxChargeTime},
	{Err: types.ErrInvalidMinAmplitude, Code: CodeInvalidMinAmplitude},
	{Err: types.ErrInvalidMaxAmplitude, Code: CodeInvalidMaxAmplitude},
	{Err: types.ErrInvalidName, Code: CodeInvalidName},
	{Err: types.ErrInvalidID, Code: CodeInvalidID},
	{Err: generator.ErrParamsEmpty, Code: CodeParamsEmpty},
	{Err: generator.ErrTrackerIsAlreadyRunning, Code: CodeTrackerIsAlreadyRunning},
	{Err: generator.ErrTrackerIsAlreadyStopped, Code: CodeTrackerIsAlreadyStopped},
	{Err: generator.ErrTrackerHasNoRoutes, Code: CodeTrackerHasNoRoutes},
	{Err: generator.ErrNoRoutes, Code: CodeNoRoutes},
	{Err: generator.ErrNoSensor, Code: CodeNoSensor},
	{Err: generator.ErrMaxNumRoutesExceeded, Code: CodeMaxNumRoutesExceeded},
	{Err: generator.ErrMaxNumTracksExceeded, Code: CodeMaxNumTracksExceeded},
	{Err: generator.ErrMaxNumSegmentsExceeded, Code: CodeMaxNumSegmentsExceeded},
	{Err: generator.ErrMaxNumSensorsExceeded, Code: CodeMaxNumSensorsExceeded},
	{Err: generator.ErrTrackerNotRunning, Code: CodeTrackerNotRunning},
	{Err: generator.ErrTrackerNotPaused, Code: CodeTrackerNotPaused},
	{Err: generator.ErrTrackerOff, Code: CodeTrackerOff},
	{Err: generator.ErrTrackNotFound, Code: CodeTrackNotFound},
	{Err: generator.ErrTrackerNotFound, Code: CodeTrackerNotFound},
	{Err: generator.ErrRouteNotFound, Code: CodeRouteNotFound},
	{Err: generator.ErrSensorNotFound, Code: CodeSensorNotFound},
	{Err: generator.ErrInvalidTrackerVersion, Code: CodeInvalidTrackerVersion},
	{Err: generator.ErrLoadingTracker, Code: CodeLoadingTracker},
	{Err: generator.ErrUnloadingTracker, Code: CodeUnloadingTracker},
	{Err: generator.ErrSensorAlreadyExists, Code: CodeSensorAlreadyExists},
	{Err: generator.ErrNoTracker, Code: CodeNoTracker},
	{Err: generator.ErrInvalidParams, Code: CodeInvalidParams},
	{Err: generator.ErrStorageInsert, Code: CodeStorageInsert},
	{Err: generator.ErrStorageUpdate, Code: CodeStorageUpdate},
	{Err: generator.ErrStorageDelete, Code: CodeStorageDelete},
	{Err: generator.ErrStorageSearch, Code: CodeStorageSearch},
}
