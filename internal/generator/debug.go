package generator

import "github.com/mmadfox/gpsgend/internal/types"

func Debug_InjectInvalidDatatForTracker(trk *Tracker, fieldName string) {
	switch fieldName {
	case "status.paused":
		trk.status = types.Paused
	case "routesSnapshot":
		trk.routesSnapshot = []byte("invalidData")
	case "sensorsSnapshot":
		trk.sensorsSnapshot = []byte("invalidData")
	default:
		panic("DebugInvalidInvariantFroTracker")
	}
}
