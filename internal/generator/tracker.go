package generator

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/gpsgend/internal/types"
)

const (
	MaxRoutesPerTracker = 5
	MaxTracksPerRoute   = 10
	MaxSegmentsPerTrack = 128
)

type Tracker struct {
	id              types.ID
	status          types.DeviceStatus
	model           types.Model
	color           types.Color
	userID          types.CustomID
	description     types.Description
	offline         types.Offline
	elevation       types.Elevation
	battery         types.Battery
	speed           types.Speed
	props           properties.Properties
	numSensors      int
	numRoutes       int
	snapshot        types.Raw
	routesSnapshot  types.Raw
	sensorsSnapshot types.Raw
	skipOffline     bool
}

func (t *Tracker) ID() types.ID {
	return t.id
}

func (t *Tracker) IsRunning() bool {
	return t.status == types.Running
}

func (t *Tracker) IsTrackerOff() bool {
	return t.status == types.Paused
}

func (t *Tracker) Model() types.Model {
	return t.model
}

func (t *Tracker) Properties() properties.Properties {
	return t.props
}

func (t *Tracker) HasNoRoutes() bool {
	return t.numRoutes == 0
}

func (t *Tracker) UpdateInfo(opts UpdateTrackerOptions) (bool, error) {
	if t.status == types.Paused {
		return false, errTrackerOff(t)
	}

	if opts.Color != nil {
		t.color = *opts.Color
	}

	if opts.Descr != nil {
		t.description = *opts.Descr
	}

	if opts.Model != nil {
		t.model = *opts.Model
	}

	if opts.UserID != nil {
		t.userID = *opts.UserID
	}

	return true, nil
}

func (t *Tracker) NewProcess() (newProc *gpsgen.Device, err error) {
	if t.numRoutes == 0 {
		return nil, ErrTrackerHasNoRoutes
	}

	if t.status == types.Paused {
		return nil, errTrackerOff(t)
	}

	opts := gpsgen.NewDeviceOptions()
	opts.ID = t.id.String()
	opts.Model = t.model.String()
	opts.Color = t.color.String()
	opts.UserID = t.userID.String()
	opts.Descr = t.description.String()

	opts.Navigator.Elevation.Min = t.elevation.Min()
	opts.Navigator.Elevation.Max = t.elevation.Max()
	opts.Navigator.Elevation.Amplitude = t.elevation.Amplitude()
	opts.Navigator.Elevation.Mode = t.elevation.Mode()

	opts.Navigator.SkipOffline = t.skipOffline
	opts.Navigator.Offline.Min = t.offline.Min()
	opts.Navigator.Offline.Max = t.offline.Max()

	opts.Battery.Min = t.battery.Min()
	opts.Battery.Max = t.battery.Max()
	opts.Battery.ChargeTime = t.battery.ChargeTime()

	opts.Speed.Min = t.speed.Min()
	opts.Speed.Max = t.speed.Max()
	opts.Speed.Amplitude = t.speed.Amplitude()

	newProc, err = gpsgen.NewDevice(opts)
	if err != nil {
		return nil, err
	}

	routes, err := gpsgen.DecodeRoutes(t.routesSnapshot)
	if err != nil {
		return nil, err
	}
	newProc.AddRoute(routes...)

	if t.numSensors > 0 {
		sensors, err := gpsgen.DecodeSensors(t.sensorsSnapshot)
		if err != nil {
			return nil, err
		}
		newProc.AddSensor(sensors...)
	}

	t.status = types.Running
	return
}

func (t *Tracker) Stop() error {
	if t.status == types.Stopped {
		return errTrackerOff(t)
	}
	if t.status == types.Stopped {
		return ErrTrackerIsAlreadyStopped
	}
	t.status = types.Stopped
	return nil
}

func (t *Tracker) Color() types.Color {
	return t.color
}

func (t *Tracker) Description() types.Description {
	return t.description
}

func (t *Tracker) UserID() types.CustomID {
	return t.userID
}

func (t *Tracker) Speed() types.Speed {
	return t.speed
}

func (t *Tracker) Battery() types.Battery {
	return t.battery
}

func (t *Tracker) Elevation() types.Elevation {
	return t.elevation
}

func (t *Tracker) Offline() types.Offline {
	return t.offline
}

func (t *Tracker) SkipOffline() bool {
	return t.skipOffline
}

func (t *Tracker) AddRoute(route *gpsgen.Route) ([]*gpsgen.Route, error) {
	return t.AddRoutes([]*gpsgen.Route{route})
}

func (t *Tracker) AddRoutes(newRoutes []*gpsgen.Route) ([]*gpsgen.Route, error) {
	if t.status == types.Paused {
		return nil, errTrackerOff(t)
	}
	if len(newRoutes) == 0 {
		return nil, ErrNoRoutes
	}
	if err := t.validateRoutes(newRoutes); err != nil {
		return nil, err
	}
	return t.appendRoutes(newRoutes)
}

func (t *Tracker) RemoveRoute(routeID types.ID) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if t.numRoutes == 0 {
		return ErrNoRoutes
	}

	routes, err := gpsgen.DecodeRoutes(t.routesSnapshot)
	if err != nil {
		return err
	}

	var ok bool
	rid := routeID.String()
	for i := 0; i < len(routes); i++ {
		if routes[i].ID() == rid {
			routes = append(routes[:i], routes[i+1:]...)
			ok = true
			break
		}
	}

	if ok {
		t.numRoutes = len(routes)
		if t.numRoutes > 0 {
			data, err := gpsgen.EncodeRoutes(routes)
			if err != nil {
				return err
			}
			t.routesSnapshot = data
		} else {
			t.routesSnapshot = make(types.Raw, 0)
		}
		return nil
	}

	return fmt.Errorf("%w Tracker{ID:%s}",
		ErrRouteNotFound, t.id)
}

func (t *Tracker) RouteAt(index int) (*gpsgen.Route, error) {
	if index <= 0 || index > t.numRoutes {
		return nil, fmt.Errorf("%w by index for Tracker{ID:%s}",
			ErrRouteNotFound, t.id)
	}

	routes, err := gpsgen.DecodeRoutes(t.routesSnapshot)
	if err != nil {
		return nil, err
	}

	return routes[index-1], nil
}

func (t *Tracker) RouteByID(routeID types.ID) (route *gpsgen.Route, err error) {
	if t.numRoutes == 0 {
		return nil, ErrNoRoutes
	}

	routes, err := gpsgen.DecodeRoutes(t.routesSnapshot)
	if err != nil {
		return nil, err
	}

	rid := routeID.String()
	for i := 0; i < len(routes); i++ {
		if routes[i].ID() == rid {
			route = routes[i]
			break
		}
	}
	if route == nil {
		err = fmt.Errorf("%w by id Tracker{ID:%s}",
			ErrRouteNotFound, t.id)
	}
	return
}

func (t *Tracker) ResetRoutes() error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}
	t.numRoutes = 0
	t.routesSnapshot = make(types.Raw, 0)
	return nil
}

func (t *Tracker) NumRoutes() int {
	return t.numRoutes
}

func (t *Tracker) Routes() ([]*gpsgen.Route, error) {
	if t.numRoutes == 0 {
		return []*gpsgen.Route{}, nil
	}
	return gpsgen.DecodeRoutes(t.routesSnapshot)
}

func (t *Tracker) appendRoutes(newRoutes []*gpsgen.Route) ([]*gpsgen.Route, error) {
	seen := make(map[string]struct{})
	new := make([]*gpsgen.Route, 0, len(newRoutes)+t.numRoutes)

	if t.numRoutes > 0 {
		prevRoutes, err := gpsgen.DecodeRoutes(t.routesSnapshot)
		if err != nil {
			return nil, err
		}
		for i := 0; i < len(prevRoutes); i++ {
			route := prevRoutes[i]
			_, ok := seen[route.ID()]
			if ok {
				continue
			}
			seen[route.ID()] = struct{}{}
			new = append(new, route)
		}
	}

	for i := 0; i < len(newRoutes); i++ {
		route := newRoutes[i]
		_, ok := seen[route.ID()]
		if ok {
			continue
		}
		seen[route.ID()] = struct{}{}
		new = append(new, route)
	}

	data, err := gpsgen.EncodeRoutes(new)
	if err != nil {
		return nil, err
	}

	t.routesSnapshot = data
	t.numRoutes = len(new)
	return new, nil
}

func (t *Tracker) validateRoutes(routes []*gpsgen.Route) error {
	if t.numRoutes+len(routes) > MaxRoutesPerTracker {
		return ErrMaxNumRoutesExceeded
	}

	var err error

loop:
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		if route.NumTracks() == 0 || route.NumTracks() > MaxTracksPerRoute {
			err = fmt.Errorf("%w routeIndex: %d",
				ErrMaxNumTracksExceeded, i)
			break loop
		}
		for j := 0; j < route.NumTracks(); j++ {
			track := route.TrackAt(j)
			if track.NumSegments() == 0 || track.NumSegments() > MaxSegmentsPerTrack {
				err = fmt.Errorf("%w routeIndex: %d, trackIndex: %d",
					ErrMaxNumSegmentsExceeded, i, j)
				break loop
			}
		}
	}

	return err
}

func (t *Tracker) AddSensor(sensor *gpsgen.Sensor) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}
	return nil
}

func (t *Tracker) RemoveSensorByID(id types.ID) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}
	return nil
}

func (t *Tracker) Sensors() []*gpsgen.Sensor {
	return nil
}

func (t *Tracker) ResetStatus() {
	t.status = types.Stopped
}

func (t *Tracker) TakeSnapshot(tracker *gpsgen.Device) error {
	data, err := gpsgen.EncodeTracker(tracker)
	if err != nil {
		return err
	}
	t.snapshot = data
	t.status = types.Paused
	return nil
}

func (t *Tracker) FromSnapshot() (*gpsgen.Device, error) {
	if t.status != types.Paused || len(t.snapshot) == 0 {
		return nil, ErrTrackerNotPaused
	}
	return nil, nil
}
