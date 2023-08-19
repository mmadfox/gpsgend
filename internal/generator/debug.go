package generator

func Debug_InjectInvalidDatatForTracker(trk *Tracker, fieldName string) {
	switch fieldName {
	case "routesSnapshot":
		trk.routesSnapshot = []byte("invalidData")
	case "sensorsSnapshot":
		trk.sensorsSnapshot = []byte("invalidData")
	default:
		panic("DebugInvalidInvariantFroTracker")
	}
}
