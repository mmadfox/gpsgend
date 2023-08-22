package codes

import (
	"net/http"

	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

type code struct {
	Err  error
	Code int
}

var table = []code{
	// http.StatusBadRequest
	{Err: types.ErrInvalidMinValue, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidMaxValue, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidRangeValue, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidMinChargeTime, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidMaxChargeTime, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidMinAmplitude, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidMaxAmplitude, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidName, Code: http.StatusBadRequest},
	{Err: types.ErrInvalidID, Code: http.StatusBadRequest},
	{Err: generator.ErrParamsEmpty, Code: http.StatusBadRequest},

	// http.StatusUnprocessableEntity
	{Err: generator.ErrTrackerIsAlreadyRunning, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrTrackerIsAlreadyStopped, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrTrackerHasNoRoutes, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrNoRoutes, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrNoSensor, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrMaxNumRoutesExceeded, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrMaxNumTracksExceeded, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrMaxNumSegmentsExceeded, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrMaxNumSensorsExceeded, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrTrackerNotRunning, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrTrackerNotPaused, Code: http.StatusUnprocessableEntity},
	{Err: generator.ErrTrackerOff, Code: http.StatusUnprocessableEntity},

	// http.StatusNotFound
	{Err: generator.ErrTrackNotFound, Code: http.StatusNotFound},
	{Err: generator.ErrTrackerNotFound, Code: http.StatusNotFound},
	{Err: generator.ErrRouteNotFound, Code: http.StatusNotFound},
	{Err: generator.ErrSensorNotFound, Code: http.StatusNotFound},

	// http.StatusInsufficientStorage
	{Err: generator.ErrInvalidTrackerVersion, Code: http.StatusInsufficientStorage},
	{Err: generator.ErrLoadingTracker, Code: http.StatusInsufficientStorage},
	{Err: generator.ErrUnloadingTracker, Code: http.StatusInsufficientStorage},
}
