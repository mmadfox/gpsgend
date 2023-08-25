package broker

import (
	"time"

	"github.com/mmadfox/go-gpsgen"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/types"
)

func makeTrackerCreatedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_CREATED,
		Payload: &gpsgendproto.Event_TrackerCreatedEvent{
			TrackerCreatedEvent: &gpsgendproto.Event_TrackerCreated{
				TrackerId: trackerID.String(),
				CreatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerRemovedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_REMOVED,
		Payload: &gpsgendproto.Event_TrackerRemovedEvent{
			TrackerRemovedEvent: &gpsgendproto.Event_TrackerRemoved{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerUpdatedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_UPDATED,
		Payload: &gpsgendproto.Event_TrackerUpdatedEvent{
			TrackerUpdatedEvent: &gpsgendproto.Event_TrackerUpdated{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerStartedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_STARTED,
		Payload: &gpsgendproto.Event_TrackerStartedEvent{
			TrackerStartedEvent: &gpsgendproto.Event_TrackerStarted{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerStoppedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_STOPPED,
		Payload: &gpsgendproto.Event_TrackerStoppedEvent{
			TrackerStoppedEvent: &gpsgendproto.Event_TrackerStopped{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerRoutesAddedEvent(trackerID types.ID, routes []*gpsgen.Route) *gpsgendproto.Event {
	routesIDs := make([]string, len(routes))
	for i := 0; i < len(routes); i++ {
		routesIDs[i] = routes[i].ID()
	}
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_ROUTE_ADDED,
		Payload: &gpsgendproto.Event_TrackerRoutesAddedEvent{
			TrackerRoutesAddedEvent: &gpsgendproto.Event_TrackerRoutesAdded{
				TrackerId: trackerID.String(),
				RoutesId:  routesIDs,
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerRoutesRemoveEvent(trackerID, routeID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_ROUTE_REMOVED,
		Payload: &gpsgendproto.Event_TrackerRouteRemovedEvent{
			TrackerRouteRemovedEvent: &gpsgendproto.Event_TrackerRouteRemoved{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerRoutesResetedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_ROUTES_RESETED,
		Payload: &gpsgendproto.Event_TrackerRoutesResetedEvent{
			TrackerRoutesResetedEvent: &gpsgendproto.Event_TrackerRoutesReseted{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerNavigatorResetedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_NAVIGATOR_RESETED,
		Payload: &gpsgendproto.Event_TrackerNavigatorResetedEvent{
			TrackerNavigatorResetedEvent: &gpsgendproto.Event_TrackerNavigatorReseted{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerNavigatorJumpedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_NAVIGATOR_JUMPED,
		Payload: &gpsgendproto.Event_TrackerNavigatorJumpedEvent{
			TrackerNavigatorJumpedEvent: &gpsgendproto.Event_TrackerNavigatorJumped{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerSensorAddedEvent(trackerID, sensorID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_SENSOR_ADDED,
		Payload: &gpsgendproto.Event_TrackerSensorAddedEvent{
			TrackerSensorAddedEvent: &gpsgendproto.Event_TrackerSensorAdded{
				TrackerId: trackerID.String(),
				SenesorId: sensorID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerSensorRemovedEvent(trackerID, sensorID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_SENSOR_REMOVED,
		Payload: &gpsgendproto.Event_TrackerSensorRemovedEvent{
			TrackerSensorRemovedEvent: &gpsgendproto.Event_TrackerSensorRemoved{
				TrackerId: trackerID.String(),
				SenesorId: sensorID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerShutdownedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_SHUTDOWNED,
		Payload: &gpsgendproto.Event_TrackerShutdownedEvent{
			TrackerShutdownedEvent: &gpsgendproto.Event_TrackerShutdowned{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}

func makeTrackerResumedEvent(trackerID types.ID) *gpsgendproto.Event {
	return &gpsgendproto.Event{
		Kind: gpsgendproto.Event_TRACKER_RESUMED,
		Payload: &gpsgendproto.Event_TrackerResumedEvent{
			TrackerResumedEvent: &gpsgendproto.Event_TrackerResumed{
				TrackerId: trackerID.String(),
				UpdatedAt: time.Now().Unix(),
			},
		},
	}
}
