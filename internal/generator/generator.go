package generator

import (
	"context"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/types"
)

type Generator struct {
	storage   Storage
	processes Processes
}

func New(
	storage Storage,
	processes Processes,
) *Generator {
	return &Generator{
		storage:   storage,
		processes: processes,
	}
}

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

	newTracker, err := trackerBuilder.Build()
	if err != nil {
		return nil, err
	}

	if err := g.storage.Insert(ctx, newTracker); err != nil {
		return nil, err
	}

	return newTracker, nil
}

func (g *Generator) RemoveTracker(ctx context.Context, trackID types.ID) error {
	if g.processes.HasTracker(trackID.String()) {
		_ = g.processes.Detach(trackID.String())
	}
	return g.storage.Delete(ctx, trackID)
}

func (g *Generator) UpdateTracker(
	ctx context.Context,
	trackerID types.ID,
	opts UpdateTrackerOptions,
) error {
	if opts.isEmpty() {
		return nil
	}

	if err := opts.validate(); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	trackerUpdatedOk, err := tracker.UpdateInfo(opts)
	if err != nil {
		return err
	}

	if trackerUpdatedOk {
		if err := g.storage.Update(ctx, tracker); err != nil {
			return err
		}
	}

	g.updateTrackerProcess(trackerID.String(), opts)
	return nil
}

func (g *Generator) FindTracker(ctx context.Context, trackerID types.ID) (*Tracker, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}
	return g.storage.FindTracker(ctx, trackerID)
}

func (g *Generator) StartTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	if g.processes.HasTracker(trackerID.String()) {
		return ErrTrackerIsAlreadyRunning
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	newProc, err := tracker.NewProcess()
	if err != nil {
		return err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	g.processes.Attach(newProc)
	return nil
}

func (g *Generator) StopTracker(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	defer func() {
		_ = g.processes.Detach(trackerID.String())
	}()

	if err := tracker.Stop(); err != nil {
		return err
	}

	return g.storage.Update(ctx, tracker)
}

func (g *Generator) TrackerState(ctx context.Context, trackerID types.ID) (*proto.Device, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return nil, ErrTrackerNotRunning
	}

	return proc.State(), nil
}

func (g *Generator) AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*gpsgen.Route) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	currRoutes, err := tracker.AddRoutes(newRoutes)
	if err != nil {
		return err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		proc.ResetRoutes()
		proc.AddRoute(currRoutes...)
	}

	return nil
}

func (g *Generator) RemoveRoute(ctx context.Context, trackerID types.ID, routeID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.RemoveRoute(routeID); err != nil {
		return err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		proc.RemoveRoute(routeID.String())
		if tracker.HasNoRoutes() {
			g.processes.Detach(trackerID.String())
		}
	}
	return nil
}

func (g *Generator) Routes(ctx context.Context, trackerID, routeID types.ID) ([]*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	if err := validateType(routeID, "route.id"); err != nil {
		return nil, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	return tracker.Routes()
}

func (g *Generator) RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	route, err := tracker.RouteAt(routeIndex)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (g *Generator) RouteByID(ctx context.Context, trackerID, routeID types.ID) (*gpsgen.Route, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	route, err := tracker.RouteByID(routeID)
	if err != nil {
		return nil, err
	}

	return route, nil
}

func (g *Generator) ResetRoutes(ctx context.Context, trackerID types.ID) (bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return false, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return false, err
	}

	if err := tracker.ResetRoutes(); err != nil {
		return false, err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		proc.ResetRoutes()
	}

	return true, nil
}

func (g *Generator) ResetNavigator(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return ErrTrackerNotRunning
	}

	proc.ResetNavigator()
	return nil
}

func (g *Generator) ToNextRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.ToNextRoute()

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) ToPrevRoute(ctx context.Context, trackerID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.ToPrevRoute()

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) MoveToRoute(ctx context.Context, trackerID types.ID, routeIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToRoute(routeIndex)

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) MoveToRouteByID(ctx context.Context, trackerID types.ID, routeID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(trackerID, "route.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToRouteByID(routeID.String())

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) MoveToTrack(ctx context.Context, trackerID types.ID, routeIndex, trackIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToTrack(routeIndex, trackIndex)

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID types.ID) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(trackerID, "route.id"); err != nil {
		return types.Navigator{}, false, err
	}

	if err := validateType(trackerID, "track.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToTrackByID(routeID.String(), trackID.String())

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) MoveToSegment(ctx context.Context, trackerID types.ID, routeIndex, trackIndex, segmentIndex int) (types.Navigator, bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return types.Navigator{}, false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if !ok {
		return types.Navigator{}, false, ErrTrackerNotRunning
	}

	next := proc.MoveToSegment(routeIndex, trackIndex, segmentIndex)

	return types.NavigatorFromProc(proc), next, nil
}

func (g *Generator) AddSensor(ctx context.Context, trackerID types.ID, sensor *gpsgen.Sensor) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	if err := tracker.AddSensor(sensor); err != nil {
		return err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		proc.AddSensor(sensor)
	}

	return nil
}

func (g *Generator) RemoveSensor(ctx context.Context, trackerID types.ID, sensorID types.ID) (bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return false, err
	}
	if err := validateType(sensorID, "sensor.id"); err != nil {
		return false, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return false, err
	}

	if err := tracker.RemoveSensorByID(sensorID); err != nil {
		return false, nil
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return false, err
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		proc.RemoveSensor(sensorID.String())
	}

	return false, nil
}

func (g *Generator) Sensors(ctx context.Context, trackerID types.ID) ([]*gpsgen.Sensor, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return nil, err
	}

	return tracker.Sensors(), nil
}

func (g *Generator) Shutdown(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	isNotRunning := tracker.IsRunning()
	if isNotRunning {
		return ErrTrackerNotRunning
	}

	invalidProc := false
	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
		tracker.TakeSnapshot(proc)
		g.processes.Detach(proc.ID())
	} else {
		invalidProc = true
		tracker.ResetStatus()
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	if invalidProc {
		return ErrTrackerNotRunning
	}

	return nil
}

func (g *Generator) Resume(ctx context.Context, trackerID types.ID) error {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return err
	}

	proc, err := tracker.FromSnapshot()
	if err != nil {
		return err
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	g.processes.Attach(proc)
	return nil
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
