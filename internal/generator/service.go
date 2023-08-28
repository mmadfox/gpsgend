package generator

import (
	"context"

	gpsgen "github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/types"
)

type Service interface {
	NewTracker(ctx context.Context, opts *NewTrackerOptions) (*Tracker, error)
	SearchTrackers(ctx context.Context, f Filter) (SearchResult, error)
	RemoveTracker(ctx context.Context, trackerID types.ID) error
	UpdateTracker(ctx context.Context, trackerID types.ID, opts UpdateTrackerOptions) error
	FindTracker(ctx context.Context, trackerID types.ID) (*Tracker, error)
	StartTracker(ctx context.Context, trackerID types.ID) error
	StopTracker(ctx context.Context, trackerID types.ID) error
	TrackerState(ctx context.Context, trackerID types.ID) (*proto.Device, error)
	AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*gpsgen.Route) error
	RemoveRoute(ctx context.Context, trackerID types.ID, routeID types.ID) error
	Routes(ctx context.Context, trackerID types.ID) ([]*gpsgen.Route, error)
	RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (*gpsgen.Route, error)
	RouteByID(ctx context.Context, trackerID, routeID types.ID) (*gpsgen.Route, error)
	ResetRoutes(ctx context.Context, trackerID types.ID) (bool, error)
	ResetNavigator(ctx context.Context, trackerID types.ID) error
	ToNextRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error)
	ToPrevRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error)
	MoveToRoute(ctx context.Context, trackerID types.ID, routeIndex int) (types.Navigator, bool, error)
	MoveToRouteByID(ctx context.Context, trackerID types.ID, routeID types.ID) (types.Navigator, bool, error)
	MoveToTrack(ctx context.Context, trackerID types.ID, routeIndex, trackIndex int) (types.Navigator, bool, error)
	MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID types.ID) (types.Navigator, bool, error)
	MoveToSegment(ctx context.Context, trackerID types.ID, routeIndex, trackIndex, segmentIndex int) (types.Navigator, bool, error)
	AddSensor(ctx context.Context, trackerID types.ID, sensor *types.Sensor) error
	RemoveSensor(ctx context.Context, trackerID types.ID, sensorID types.ID) error
	Sensors(ctx context.Context, trackerID types.ID) ([]*types.Sensor, error)
	ShutdownTracker(ctx context.Context, trackerID types.ID) error
	ResumeTracker(ctx context.Context, trackerID types.ID) error
	Stats(ctx context.Context) ([]StatsItem, error)
	Sync(ctx context.Context, trackerID types.ID) error
}
