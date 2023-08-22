package generator

import "github.com/mmadfox/gpsgend/internal/types"

func Debug_InjectInvalidDatatForTracker(trk *Tracker, fieldName string) {
	switch fieldName {
	case "status.paused":
		trk.status = types.Paused
	case "routesSnapshot":
		trk.routesSnapshot = []byte("invalidData")
	default:
		panic("Debug_InjectInvalidDatatForTracker")
	}
}
