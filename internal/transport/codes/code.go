package codes

import (
	stderrors "errors"
	"net/http"
)

const (
	// default
	CodeUnknown = 0

	// block: 1 .. 100
	CodeInvalidMinValue      = 1
	CodeInvalidMaxValue      = 2
	CodeInvalidRangeValue    = 3
	CodeInvalidMinChargeTime = 4
	CodeInvalidMaxChargeTime = 5
	CodeInvalidMinAmplitude  = 6
	CodeInvalidMaxAmplitude  = 7
	CodeInvalidName          = 8
	CodeInvalidID            = 9
	CodeParamsEmpty          = 10

	// block: 100 .. 300
	CodeTrackerIsAlreadyRunning = 101
	CodeTrackerIsAlreadyStopped = 102
	CodeTrackerHasNoRoutes      = 103
	CodeNoRoutes                = 104
	CodeNoSensor                = 105
	CodeMaxNumRoutesExceeded    = 106
	CodeMaxNumTracksExceeded    = 107
	CodeMaxNumSegmentsExceeded  = 108
	CodeMaxNumSensorsExceeded   = 109
	CodeTrackerNotRunning       = 110
	CodeTrackerNotPaused        = 111
	CodeTrackerOff              = 112
	CodeInvalidTrackerVersion   = 113
	CodeLoadingTracker          = 114
	CodeUnloadingTracker        = 115

	// block: 300 .. 400
	CodeTrackNotFound   = 300
	CodeTrackerNotFound = 301
	CodeRouteNotFound   = 302
	CodeSensorNotFound  = 303
)

func FromError(err error) int {
	for i := 0; i < len(table); i++ {
		code := table[i]
		if stderrors.Is(err, code.Err) {
			return code.Code
		}
	}
	return 0
}

func ToHTTP(code int) int {
	if code == 0 {
		return http.StatusInternalServerError
	}

	if code >= 1 && code <= 100 {
		return http.StatusBadRequest
	}

	if code > 100 && code <= 300 {
		return http.StatusUnprocessableEntity
	}

	if code > 300 && code <= 400 {
		return http.StatusNotFound
	}
	return 0
}
