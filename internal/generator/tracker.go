package generator

import (
	"fmt"
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/gpsgend/internal/types"
)

const (
	MaxRoutesPerTracker = 5
	MaxTracksPerRoute   = 10
	MaxSegmentsPerTrack = 128
	MaxSensorsPerDevice = 10
)

// Tracker represents a GPS tracker device with various properties and capabilities.
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
	createdAt       time.Time
	updatedAt       time.Time
	runningAt       time.Time
	stoppedAt       time.Time
}

// ID returns the unique identifier of the tracker.
func (t *Tracker) ID() types.ID {
	return t.id
}

// IsRunning checks if the tracker is currently running.
func (t *Tracker) IsRunning() bool {
	return t.status == types.Running
}

// IsTrackerOff checks if the tracker is in a paused state.
func (t *Tracker) IsTrackerOff() bool {
	return t.status == types.Paused
}

// Model returns the model of the tracker.
func (t *Tracker) Model() types.Model {
	return t.model
}

// Properties returns the custom properties associated with the tracker.
func (t *Tracker) Properties() properties.Properties {
	return t.props
}

// HasNoRoutes checks if the tracker has no routes.
func (t *Tracker) HasNoRoutes() bool {
	return t.numRoutes == 0
}

// CreatedAt returns the timestamp indicating when the tracker was created.
func (t *Tracker) CreatedAt() time.Time {
	return t.createdAt
}

// UpdatedAt returns the timestamp indicating when the tracker was last updated.
func (t *Tracker) UpdatedAt() time.Time {
	return t.updatedAt
}

// RunningAt returns the timestamp indicating when the tracker's GPS process started running.
func (t *Tracker) RunningAt() time.Time {
	return t.runningAt
}

// StoppedAt returns the timestamp indicating when the tracker's GPS process was stopped.
func (t *Tracker) StoppedAt() time.Time {
	return t.stoppedAt
}

// UpdateInfo updates the tracker's information with the provided options.
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

	t.updatedAt = time.Now()
	return true, nil
}

// NewProcess creates a new GPS device process for the tracker.
func (t *Tracker) NewProcess() (newProc *gpsgen.Device, err error) {
	if t.status == types.Running {
		return nil, ErrTrackerIsAlreadyRunning
	}

	if t.status == types.Paused {
		return nil, errTrackerOff(t)
	}

	if t.numRoutes == 0 {
		return nil, ErrTrackerHasNoRoutes
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
	t.runningAt = time.Now()
	return
}

// Stop stops the tracker's GPS device process.
func (t *Tracker) Stop() error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}
	if t.status == types.Stopped {
		return ErrTrackerIsAlreadyStopped
	}
	t.status = types.Stopped
	t.stoppedAt = time.Now()
	return nil
}

// Color returns the color of the tracker.
func (t *Tracker) Color() types.Color {
	return t.color
}

// Description returns the description of the tracker.
func (t *Tracker) Description() types.Description {
	return t.description
}

// UserID returns the custom user identifier associated with the tracker.
func (t *Tracker) UserID() types.CustomID {
	return t.userID
}

// Speed returns the speed information of the tracker.
func (t *Tracker) Speed() types.Speed {
	return t.speed
}

// Battery returns the battery information of the tracker.
func (t *Tracker) Battery() types.Battery {
	return t.battery
}

// Elevation returns the elevation information of the tracker.
func (t *Tracker) Elevation() types.Elevation {
	return t.elevation
}

// Offline returns the offline configuration of the tracker.
func (t *Tracker) Offline() types.Offline {
	return t.offline
}

// SkipOffline returns whether skipping offline data is enabled for the tracker.
func (t *Tracker) SkipOffline() bool {
	return t.skipOffline
}

// AddRoute adds a new route to the tracker.
func (t *Tracker) AddRoute(route *gpsgen.Route) ([]*gpsgen.Route, error) {
	return t.AddRoutes([]*gpsgen.Route{route})
}

// AddRoutes adds multiple new routes to the tracker.
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
	t.updatedAt = time.Now()
	return t.appendRoutes(newRoutes)
}

// RemoveRoute removes a route from the tracker by route ID.
func (t *Tracker) RemoveRoute(routeID types.ID) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if t.numRoutes == 0 {
		return nil
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

		t.updatedAt = time.Now()
		return nil
	}

	return fmt.Errorf("%w Tracker{ID:%s}",
		ErrRouteNotFound, t.id)
}

// RouteAt returns the route at the specified index.
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

// RouteByID returns the route with the specified ID.
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

// ResetRoutes removes all routes from the tracker.
func (t *Tracker) ResetRoutes() error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}
	t.numRoutes = 0
	t.routesSnapshot = make(types.Raw, 0)
	t.updatedAt = time.Now()
	return nil
}

// NumRoutes returns the number of routes associated with the tracker.
func (t *Tracker) NumRoutes() int {
	return t.numRoutes
}

// Routes returns a list of all routes associated with the tracker.
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
			if _, ok := seen[route.ID()]; ok {
				continue
			}
			seen[route.ID()] = struct{}{}
			new = append(new, route)
		}
	}

	for i := 0; i < len(newRoutes); i++ {
		route := newRoutes[i]
		if _, ok := seen[route.ID()]; ok {
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
	t.updatedAt = time.Now()
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

// AddSensor adds one or more sensors to the tracker.
func (t *Tracker) AddSensor(newSensors ...*gpsgen.Sensor) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if len(newSensors) == 0 {
		return ErrNoSensors
	}

	if t.numSensors+len(newSensors) > MaxSensorsPerDevice {
		return ErrMaxNumSensorsExceeded
	}

	seen := make(map[string]struct{})
	new := make([]*gpsgen.Sensor, 0, t.numSensors+len(newSensors))

	if t.numSensors > 0 {
		prevSensors, err := gpsgen.DecodeSensors(t.sensorsSnapshot)
		if err != nil {
			return err
		}
		for i := 0; i < len(prevSensors); i++ {
			sensor := prevSensors[i]
			if _, ok := seen[sensor.ID()]; ok {
				continue
			}
			seen[sensor.ID()] = struct{}{}
			new = append(new, sensor)
		}
	}

	for i := 0; i < len(newSensors); i++ {
		sensor := newSensors[i]
		if sensor == nil {
			continue
		}
		if _, ok := seen[sensor.ID()]; ok {
			continue
		}
		seen[sensor.ID()] = struct{}{}
		new = append(new, sensor)
	}

	if len(new) > 0 {
		data, err := gpsgen.EncodeSensors(new)
		if err != nil {
			return err
		}
		t.sensorsSnapshot = data
	} else {
		t.sensorsSnapshot = nil
	}

	t.numSensors = len(new)
	t.updatedAt = time.Now()
	return nil
}

// RemoveSensorByID removes a sensor from the tracker by sensor ID.
func (t *Tracker) RemoveSensorByID(id types.ID) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if t.numSensors == 0 {
		return nil
	}

	sensors, err := gpsgen.DecodeSensors(t.sensorsSnapshot)
	if err != nil {
		return err
	}

	var ok bool
	sid := id.String()
	for i := 0; i < len(sensors); i++ {
		if sensors[i].ID() == sid {
			sensors = append(sensors[:i], sensors[i+1:]...)
			ok = true
			break
		}
	}

	if ok {
		t.numSensors = len(sensors)
		if t.numSensors > 0 {
			data, err := gpsgen.EncodeSensors(sensors)
			if err != nil {
				return err
			}
			t.sensorsSnapshot = data
		} else {
			t.sensorsSnapshot = make(types.Raw, 0)
		}
		t.updatedAt = time.Now()
		return nil
	}

	return fmt.Errorf("%w Tracker{ID:%s}",
		ErrSensorNotFound, t.id)
}

// Sensors returns a list of all sensors associated with the tracker.
func (t *Tracker) Sensors() ([]*gpsgen.Sensor, error) {
	if t.numSensors == 0 {
		return []*gpsgen.Sensor{}, nil
	}
	return gpsgen.DecodeSensors(t.sensorsSnapshot)
}

// ResetStatus resets the tracker's status to stopped.
func (t *Tracker) ResetStatus() {
	t.status = types.Stopped
	t.updatedAt = time.Now()
}

// ShutdownProcess saves the current GPS device state and pauses the tracker's process.
func (t *Tracker) ShutdownProcess(tracker *gpsgen.Device) error {
	data, err := gpsgen.EncodeTracker(tracker)
	if err != nil {
		return err
	}
	t.snapshot = data
	t.status = types.Paused
	t.updatedAt = time.Now()
	return nil
}

// ResumeProcess resumes a paused tracker's process with the provided GPS device state.
func (t *Tracker) ResumeProcess() (*gpsgen.Device, error) {
	if t.status != types.Paused {
		return nil, ErrTrackerNotPaused
	}

	proc, err := gpsgen.DecodeTracker(t.snapshot)
	if err != nil {
		return nil, err
	}

	t.status = types.Running
	t.snapshot = nil
	t.updatedAt = time.Now()
	return proc, nil
}
