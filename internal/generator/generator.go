package generator

import (
	"context"

	"github.com/mmadfox/go-gpsgen/navigator"
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

	if opts.Color != nil {
		tracker.ChangeColor(*opts.Color)
	}
	if opts.Descr != nil {
		tracker.ChangeDescription(*opts.Descr)
	}
	if opts.Model != nil {
		tracker.ChangeModel(*opts.Model)
	}
	if opts.UserID != nil {
		tracker.ChangeUserID(*opts.UserID)
	}

	if err := g.storage.Update(ctx, tracker); err != nil {
		return err
	}

	if !tracker.IsRunning() {
		return nil
	}

	proc, ok := g.processes.Lookup(trackerID.String())
	if ok {
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

	if tracker.HasNoRoutes() {
		return ErrNoRoutes
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

	if g.processes.HasTracker(trackerID.String()) {
		g.processes.Detach(trackerID.String())
	}

	isNotRunning := !tracker.IsRunning()
	if isNotRunning {
		return ErrTrackerIsAlreadyStopped
	}

	tracker.Stop()

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

func (g *Generator) AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*navigator.Route) error {
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
		proc.Update()
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
		switch tracker.HasNoRoutes() {
		case false:
			proc.RemoveRoute(routeID.String())
			proc.Update()
		case true:
			g.processes.Detach(trackerID.String())
		}
	}
	return nil
}

func (g *Generator) Routes(ctx context.Context, trackerID, routeID types.ID) ([]*navigator.Route, error) {
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

func (g *Generator) RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (*navigator.Route, error) {
	return nil, nil
}

func (g *Generator) RouteByID(ctx context.Context, trackerID, routeID types.ID) (*navigator.Route, error) {
	return nil, nil
}

func (g *Generator) ResetRoutes(ctx context.Context, trackerID types.ID) (bool, error) {
	if err := validateType(trackerID, "tracker.id"); err != nil {
		return false, err
	}

	tracker, err := g.storage.FindTracker(ctx, trackerID)
	if err != nil {
		return false, err
	}

	tracker.ResetRoutes()
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
