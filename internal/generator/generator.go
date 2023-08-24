package generator

import (
	"context"
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/types"
)

// Generator is a service for managing GPS trackers, routes, and related processes.
type Generator struct {
	trackers    Storage
	processes   Processes
	bootstraper Bootstraper
	query       Query
	eventPub    EventPublisher
}

// New creates a new instance of the Generator.
func New(
	s Storage,
	p Processes,
	b Bootstraper,
	q Query,
	ep EventPublisher,
) *Generator {
	return &Generator{
		trackers:    s,
		processes:   p,
		bootstraper: b,
		query:       q,
		eventPub:    ep,
	}
}

// NewTracker creates a new tracker instance with the given options and inserts it into storage.
func (g *Generator) NewTracker(ctx context.Context, opts NewTrackerOptions) (*Tracker, error) {
	trackerBuilder := NewTrackerBuilder()
	trackerBuilder.ID(types.NewID())

	if opts.Model != nil {
		trackerBuilder.Model(*opts.Model)
	}
	if opts.Color != nil {
		trackerBuilder.Color(*opts.Color)
	}
	if opts.UserID != nil {
		trackerBuilder.CustomID(*opts.UserID)
	}
	if opts.Descr != nil {
		trackerBuilder.Description(*opts.Descr)
	}
	if opts.Props != nil {
		trackerBuilder.Props(*opts.Props)
	}

	trackerBuilder.SkipOffline(opts.SkipOffline)
	trackerBuilder.Offline(opts.Offline)
	trackerBuilder.Elevation(opts.Elevation)
	trackerBuilder.Speed(opts.Speed)
	trackerBuilder.Battery(opts.Battery)
	trackerBuilder.CreatedAt(time.Now())

	newTracker, err := trackerBuilder.Build()
	if err != nil {
		return nil, err
	}

	if err := g.trackers.Insert(ctx, newTracker); err != nil {
		return nil, err
	}

	g.eventPub.PublishTrackerCreated(ctx, newTracker.ID())

	return newTracker, nil
}

// SearchTrackers searches for trackers based on the provided filter.
// It returns a SearchResult and an error if the search operation fails.
func (g *Generator) SearchTrackers(ctx context.Context, f Filter) (SearchResult, error) {
	return g.query.SearchTrackers(ctx, f)
}

// RemoveTracker removes a tracker from storage and detaches it from any associated processes.
func (g *Generator) RemoveTracker(ctx context.Context, trackerID types.ID) error {
	g.processes.Detach(trackerID.String())

	if err := g.trackers.Delete(ctx, trackerID); err != nil {
		return err
	}

	g.eventPub.PublishTrackerRemoved(ctx, trackerID)

	return nil
}

// UpdateTracker updates the information of a tracker, validates options, and manages related processes.
func (g *Generator) UpdateTracker(
	ctx context.Context,
	trackerID types.ID,
	opts UpdateTrackerOptions,
) error {
	if opts.isEmpty() {
		return ErrParamsEmpty
	}

	if err := opts.validate(); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	trackerUpdatedOk, err := tracker.UpdateInfo(opts)
	if err != nil {
		return err
	}

	if trackerUpdatedOk {
		if err := g.trackers.Update(ctx, tracker); err != nil {
			return err
		}
	}

	g.updateTrackerProcess(trackerID.String(), opts)

	g.eventPub.PublishTrackerUpdated(ctx, trackerID)

	return nil
}

// FindTracker retrieves a tracker from storage using its ID.
func (g *Generator) FindTracker(ctx context.Context, trackerID types.ID) (*Tracker, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}
	return g.trackers.Find(ctx, trackerID)
}

// StartTracker starts a tracker process and attaches it.
func (g *Generator) StartTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	if g.processes.HasTracker(trackerID.String()) {
		return ErrTrackerIsAlreadyRunning
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	newProc, err := tracker.NewProcess()
	if err != nil {
		return err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	g.addProc(newProc)

	g.eventPub.PublishTrackerStarted(ctx, trackerID)

	return nil
}

// StopTracker stops a tracker process and detaches it.
func (g *Generator) StopTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.Stop(); err != nil {
		return err
	}

	defer func() {
		_ = g.processes.Detach(trackerID.String())
	}()

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	g.eventPub.PublishTrackerStopped(ctx, trackerID)

	return nil
}

// TrackerState retrieves the state of a running tracker process.
func (g *Generator) TrackerState(ctx context.Context, trackerID types.ID) (*proto.Device, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return nil, ErrTrackerNotRunning
	}

	return proc.State(), nil
}

// AddRoutes adds new routes to a tracker and updates related processes.
func (g *Generator) AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*gpsgen.Route) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	currRoutes, err := tracker.AddRoutes(newRoutes)
	if err != nil {
		return err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.findProc(trackerID)
	if ok {
		proc.ResetRoutes()
		proc.AddRoute(currRoutes...)
	}

	g.eventPub.PublishTrackerRoutesAdded(ctx, trackerID, newRoutes)

	return nil
}

// RemoveRoute removes a route from a tracker and updates related processes if needed.
func (g *Generator) RemoveRoute(ctx context.Context, trackerID types.ID, routeID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.RemoveRoute(routeID); err != nil {
		return err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.findProc(trackerID)
	if ok {
		proc.RemoveRoute(routeID.String())
		if tracker.HasNoRoutes() {
			g.removeProc(proc)
		}
	}

	g.eventPub.PublishTrackerRouteRemoved(ctx, trackerID, routeID)

	return nil
}

// Routes retrieves the list of routes associated with a tracker.
func (g *Generator) Routes(ctx context.Context, trackerID types.ID) ([]*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	return tracker.Routes()
}

// RouteAt retrieves a route at a specific index for a tracker.
func (g *Generator) RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	route, err := tracker.RouteAt(routeIndex)
	if err != nil {
		return nil, err
	}

	return route, nil
}

// RouteByID retrieves a route by its ID for a tracker.
func (g *Generator) RouteByID(ctx context.Context, trackerID, routeID types.ID) (*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return nil, err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	route, err := tracker.RouteByID(routeID)
	if err != nil {
		return nil, err
	}

	return route, nil
}

// ResetRoutes resets all routes for a tracker and updates related processes.
func (g *Generator) ResetRoutes(ctx context.Context, trackerID types.ID) (bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return false, err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return false, err
	}

	if err := tracker.ResetRoutes(); err != nil {
		return false, err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return false, err
	}

	proc, ok := g.findProc(trackerID)
	if ok {
		proc.ResetRoutes()
	}

	g.eventPub.PublishTrackerRoutesReseted(ctx, trackerID)

	return true, nil
}

// ResetNavigator resets the navigation state of a running tracker process.
func (g *Generator) ResetNavigator(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return ErrTrackerNotRunning
	}

	proc.ResetNavigator()

	g.eventPub.PublishTrackerNavigatorReseted(ctx, trackerID)

	return nil
}

// ToNextRoute moves the tracker's navigation to the next route
// and provides the updated navigation state.
func (g *Generator) ToNextRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.ToNextRoute()

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// ToPrevRoute moves the tracker's navigation to the previous route
// and provides the updated navigation state.
func (g *Generator) ToPrevRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.ToPrevRoute()

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// MoveToRoute moves the tracker's navigation to a specific route index
// and provides the updated navigation state.
func (g *Generator) MoveToRoute(ctx context.Context, trackerID types.ID, routeIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToRoute(routeIndex)

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// MoveToRouteByID moves the tracker's navigation to a specific route ID
// and provides the updated navigation state.
func (g *Generator) MoveToRouteByID(ctx context.Context, trackerID types.ID, routeID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToRouteByID(routeID.String())

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// MoveToTrack moves the tracker's navigation to a specific track index
// within a route and provides the updated navigation state.
func (g *Generator) MoveToTrack(ctx context.Context, trackerID types.ID, routeIndex, trackIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToTrack(routeIndex, trackIndex)

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// MoveToTrackByID moves the tracker's navigation to a specific track ID
// within a route and provides the updated navigation state.
func (g *Generator) MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(trackID, "track.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToTrackByID(routeID.String(), trackID.String())

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// MoveToSegment moves the tracker's navigation to a specific segment index
// within a route and provides the updated navigation state.
func (g *Generator) MoveToSegment(ctx context.Context, trackerID types.ID, routeIndex, trackIndex, segmentIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.findProc(trackerID)
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToSegment(routeIndex, trackIndex, segmentIndex)

	if next {
		g.eventPub.PublishTrackerNavigatorJumped(ctx, trackerID)
	}

	return types.NavigatorFromProc(proc), next, nil
}

// AddSensor adds a sensor to a tracker's list of sensors and
// updates related processes.
func (g *Generator) AddSensor(
	ctx context.Context,
	trackerID types.ID,
	sensorConf *types.Sensor,
) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.AddSensor(sensorConf); err != nil {
		return err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.findProc(trackerID)
	if ok {
		newSensor, err := tracker.makeSensorFromConf(sensorConf)
		if err != nil {
			return err
		}
		proc.AddSensor(newSensor)
	}

	g.eventPub.PublishTrackerSensorAdded(ctx, trackerID, sensorConf.ID())

	return nil
}

// RemoveSensor removes a sensor from a tracker's list
// of sensors and updates related processes.
func (g *Generator) RemoveSensor(
	ctx context.Context,
	trackerID types.ID,
	sensorID types.ID,
) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}
	if err := validateType(sensorID, "sensor.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.RemoveSensorByID(sensorID); err != nil {
		return err
	}

	proc, ok := g.findProc(trackerID)
	if ok {
		proc.RemoveSensor(sensorID.String())
	}

	g.eventPub.PublishTrackerSensorRemoved(ctx, trackerID, sensorID)

	return g.trackers.Update(ctx, tracker)
}

// Sensors retrieves the list of sensors associated with a tracker.
func (g *Generator) Sensors(
	ctx context.Context,
	trackerID types.ID,
) ([]*types.Sensor, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	return tracker.Sensors(), nil
}

// Shutdown shuts down a tracker, takes snapshots if needed, and detaches processes.
func (g *Generator) ShutdownTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	if !tracker.IsRunning() {
		return ErrTrackerNotRunning
	}

	invalidProc := false
	proc, ok := g.findProc(trackerID)
	if ok {
		tracker.ShutdownProcess(proc)
		g.removeProc(proc)
	} else {
		invalidProc = true
		tracker.ResetStatus()
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	if invalidProc {
		return ErrTrackerNotRunning
	}

	g.eventPub.PublishTrackerShutdowned(ctx, trackerID)

	return nil
}

// ResumeTracker resumes a tracker from a previously taken snapshot and attaches the process// Resume resumes a tracker from a previously taken snapshot and attaches the process.
func (g *Generator) ResumeTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.trackers.Find(ctx, trackerID)
	if err != nil {
		return err
	}

	proc, err := tracker.ResumeProcess()
	if err != nil {
		return err
	}

	if err := g.trackers.Update(ctx, tracker); err != nil {
		return err
	}

	g.addProc(proc)
	g.eventPub.PublishTrackerResumed(ctx, trackerID)

	return nil
}

func (g *Generator) Run(ctx context.Context) error {
	return g.bootstraper.LoadTrackers(ctx, g.processes)
}

func (g *Generator) Close(ctx context.Context) error {
	return g.bootstraper.UnloadTrackers(ctx, g.processes)
}

func (g *Generator) findProc(trackerID types.ID) (*gpsgen.Device, bool) {
	return g.processes.Lookup(trackerID.String())
}

func (g *Generator) removeProc(proc *gpsgen.Device) {
	g.processes.Detach(proc.ID())
}

func (g *Generator) addProc(proc *gpsgen.Device) {
	g.processes.Attach(proc)
}

func (g *Generator) updateTrackerProcess(trackerID string, opts UpdateTrackerOptions) {
	proc, ok := g.processes.Lookup(trackerID)
	if !ok {
		return
	}
	if opts.Color != nil {
		proc.SetColor((*opts.Color).RGB())
	}
	if opts.Descr != nil {
		proc.SetDescription((*opts.Descr).String())
	}
	if opts.Model != nil {
		proc.SetModel((*opts.Model).String())
	}
	if opts.UserID != nil {
		proc.SetUserID((*opts.UserID).String())
	}
}
