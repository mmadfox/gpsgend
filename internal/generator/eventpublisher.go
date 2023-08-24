package generator

import (
	"context"

	gpsgen "github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/types"
)

type EventPublisher interface {
	PublishTrackerCreated(ctx context.Context, trackerID types.ID)
	PublishTrackerRemoved(ctx context.Context, trackerID types.ID)
	PublishTrackerUpdated(ctx context.Context, trackerUD types.ID)
	PublishTrackerStarted(ctx context.Context, trackerID types.ID)
	PublishTrackerStopped(ctx context.Context, trackerID types.ID)
	PublishTrackerRoutesAdded(ctx context.Context, trackerID types.ID, routes []*gpsgen.Route)
	PublishTrackerRouteRemoved(ctx context.Context, trackerID, routeID types.ID)
	PublishTrackerRoutesReseted(ctx context.Context, trackerID types.ID)
	PublishTrackerNavigatorReseted(ctx context.Context, trackerID types.ID)
	PublishTrackerNavigatorJumped(ctx context.Context, trackerID types.ID)
	PublishTrackerSensorAdded(ctx context.Context, trackerID, sensorID types.ID)
	PublishTrackerSensorRemoved(ctx context.Context, trackerID, sensorID types.ID)
	PublishTrackerShutdowned(ctx context.Context, trackerID types.ID)
	PublishTrackerResumed(ctx context.Context, trackerID types.ID)
}
