package generator

import (
	"fmt"
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/curve"
	"github.com/mmadfox/go-gpsgen/properties"
	stdtypes "github.com/mmadfox/go-gpsgen/types"
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
	id             types.ID
	status         types.DeviceStatus
	model          types.Model
	color          types.Color
	userID         types.CustomID
	description    types.Description
	offline        types.Offline
	elevation      types.Elevation
	battery        types.Battery
	speed          types.Speed
	props          properties.Properties
	numRoutes      int
	snapshot       types.Raw
	routesSnapshot types.Raw
	sensors        map[types.ID]*types.Sensor
	skipOffline    bool
	createdAt      time.Time
	updatedAt      time.Time
	runningAt      time.Time
	stoppedAt      time.Time
	version        int
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

// Status returns the status of the tracker.
func (t *Tracker) Status() types.DeviceStatus {
	return t.status
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
		return nil, fmt.Errorf("%w, please add routes and try again",
			ErrTrackerHasNoRoutes)
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

	if len(t.sensors) > 0 {
		sensors, err := t.makeTrackerSensorsFromConf()
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

// NumSensors returns the number of sensors associated with the tracker.
func (t *Tracker) NumSensors() int {
	return len(t.sensors)
}

// Routes returns a list of all routes associated with the tracker.
func (t *Tracker) Routes() ([]*gpsgen.Route, error) {
	if t.numRoutes == 0 {
		return []*gpsgen.Route{}, nil
	}
	return gpsgen.DecodeRoutes(t.routesSnapshot)
}

func (t *Tracker) TakeSnapshot(snap *TrackerSnapshot) {
	if snap == nil {
		return
	}
	if !t.id.IsEmpty() {
		snap.ID = t.id.String()
	}
	snap.CustomID = t.userID.String()
	snap.Status.ID = int(t.status)
	snap.Status.Name = t.status.String()
	snap.Model = t.model.String()
	snap.Descr = t.description.String()
	snap.Color = t.color.String()
	snap.Offline.Min = t.offline.Min()
	snap.Offline.Max = t.offline.Max()
	snap.Elevation.Min = t.elevation.Min()
	snap.Elevation.Max = t.elevation.Max()
	snap.Elevation.Amplitude = t.Elevation().Amplitude()
	snap.Elevation.Mode = int(t.Elevation().Mode())
	snap.Battery.Min = t.battery.Min()
	snap.Battery.Max = t.battery.Max()
	snap.Battery.ChargeTime = t.Battery().ChargeTime()
	snap.Speed.Min = t.speed.Min()
	snap.Speed.Max = t.speed.Max()
	snap.Speed.Amplitude = t.speed.Amplitude()
	snap.Props = t.props
	snap.NumRoutes = t.NumRoutes()
	snap.NumSensors = t.NumSensors()
	snap.SkipOffline = t.skipOffline
	snap.Snapshot = t.snapshot
	snap.Routes = t.routesSnapshot
	snap.CreatedAt = t.createdAt.Unix()
	snap.UpdatedAt = t.updatedAt.Unix()
	snap.RunningAt = t.runningAt.Unix()
	snap.StoppedAt = t.stoppedAt.Unix()
	if len(t.sensors) > 0 {
		snap.Sensors = make([]SensorSnapshot, 0, len(t.sensors))
		for _, sensor := range t.sensors {
			snap.Sensors = append(snap.Sensors, SensorSnapshot{
				ID:        sensor.ID().String(),
				Name:      sensor.Name(),
				Min:       sensor.Min(),
				Max:       sensor.Max(),
				Amplitude: sensor.Amplitude(),
				Mode:      int(sensor.Mode()),
			})
		}
	}
	snap.Version = t.version
}

func (t *Tracker) FromSnapshot(snap *TrackerSnapshot) error {
	if snap == nil {
		return nil
	}

	t.routesSnapshot = snap.Routes
	t.version = snap.Version
	t.numRoutes = snap.NumRoutes
	t.skipOffline = snap.SkipOffline
	t.props = snap.Props
	t.updatedAt = time.Unix(snap.UpdatedAt, 0)
	t.createdAt = time.Unix(snap.CreatedAt, 0)
	t.runningAt = time.Unix(snap.RunningAt, 0)
	t.stoppedAt = time.Unix(snap.StoppedAt, 0)
	t.snapshot = snap.Snapshot

	trackeID, err := types.ParseID(snap.ID)
	if err != nil {
		return fmt.Errorf("%w: can't parse type tracker.id", err)
	}
	t.id = trackeID

	if len(snap.CustomID) > 0 {
		customID, err := types.ParseCustomID(snap.CustomID)
		if err != nil {
			return fmt.Errorf("%w: can't parse type tracker.customID", err)
		}
		t.userID = customID
	}

	status, err := types.ParseDeviceStatus(snap.Status.ID)
	if err != nil {
		return fmt.Errorf("%w: can't parse type tracker.status", err)
	}
	t.status = status

	if len(snap.Model) > 0 {
		model, err := types.ParseModel(snap.Model)
		if err != nil {
			return fmt.Errorf("%w: can't parse type tracker.model", err)
		}
		t.model = model
	}

	if len(snap.Color) > 0 {
		color, err := types.ParseColor(snap.Color)
		if err != nil {
			return fmt.Errorf("%w: can't parse type tracker.color", err)
		}
		t.color = color
	}

	if len(snap.Descr) > 0 {
		descr, err := types.ParseDescription(snap.Descr)
		if err != nil {
			return fmt.Errorf("%w: can't parse type tracker.description", err)
		}
		t.description = descr
	}

	if !snap.SkipOffline {
		offline, err := types.ParseOffline(snap.Offline.Min, snap.Offline.Max)
		if err != nil {
			return fmt.Errorf("%w: can't parse type tracker.offline", err)
		}
		t.offline = offline
	}

	elevation, err := types.ParseElevation(
		snap.Elevation.Min,
		snap.Elevation.Max,
		snap.Elevation.Amplitude,
		curve.CurveMode(snap.Elevation.Mode),
	)
	if err != nil {
		return fmt.Errorf("%w: can't parse type tracker.elevation", err)
	}
	t.elevation = elevation

	battery, err := types.ParseBattery(snap.Battery.Min, snap.Battery.Max, snap.Battery.ChargeTime)
	if err != nil {
		return fmt.Errorf("%w: can't parse type tracker.battery", err)
	}
	t.battery = battery

	speed, err := types.ParseSpeed(snap.Speed.Min, snap.Speed.Max, snap.Speed.Amplitude)
	if err != nil {
		return fmt.Errorf("%w: can't parse type tracker.speed", err)
	}
	t.speed = speed

	if len(snap.Sensors) > 0 {
		t.sensors = make(map[types.ID]*types.Sensor, len(snap.Sensors))
		for i := 0; i < len(snap.Sensors); i++ {
			sn := snap.Sensors[i]
			sid, err := types.ParseID(sn.ID)
			if err != nil {
				return fmt.Errorf("%w: can't parse type tracker.sensor[%d].id", err, i)
			}
			sensor, err := types.ParseSensor(sid, sn.Name, sn.Min, sn.Max, sn.Amplitude, sn.Mode)
			if err != nil {
				return fmt.Errorf("%w: can't parse type tracker.sensor[%d]", err, i)
			}
			t.sensors[sid] = sensor
		}
	}

	return nil
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

// AddSensor adds one sensors to the tracker.
func (t *Tracker) AddSensor(newSensor *types.Sensor) error {
	if t.sensors == nil {
		t.sensors = make(map[types.ID]*types.Sensor)
	}

	if newSensor == nil {
		return ErrNoSensor
	}

	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if len(t.sensors)+1 > MaxSensorsPerDevice {
		return ErrMaxNumSensorsExceeded
	}

	_, ok := t.sensors[newSensor.ID()]
	if ok {
		return ErrSensorAlreadyExists
	}

	t.sensors[newSensor.ID()] = newSensor
	t.updatedAt = time.Now()
	return nil
}

func (t *Tracker) ResetSensors() {
	t.sensors = make(map[types.ID]*types.Sensor)
	t.updatedAt = time.Now()
}

// RemoveSensorByID removes a sensor from the tracker by sensor ID.
func (t *Tracker) RemoveSensorByID(id types.ID) error {
	if t.status == types.Paused {
		return errTrackerOff(t)
	}

	if len(t.sensors) == 0 {
		return nil
	}

	_, ok := t.sensors[id]
	if !ok {
		return fmt.Errorf("%w for Tracker{ID:%s}", ErrSensorNotFound, t.id)
	}

	delete(t.sensors, id)

	t.updatedAt = time.Now()
	return nil
}

// Sensors returns a list of all sensors associated with the tracker.
func (t *Tracker) Sensors() []*types.Sensor {
	if len(t.sensors) == 0 {
		return []*types.Sensor{}
	}
	sensors := make([]*types.Sensor, 0, len(t.sensors))
	for _, sensor := range t.sensors {
		sensors = append(sensors, sensor)
	}
	return sensors
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

func (t *Tracker) makeSensorFromConf(conf *types.Sensor) (*gpsgen.Sensor, error) {
	return stdtypes.RestoreSensor(
		conf.ID().String(),
		conf.Name(),
		conf.Min(),
		conf.Max(),
		conf.Amplitude(),
		conf.Mode(),
	)
}

func (t *Tracker) makeTrackerSensorsFromConf() ([]*gpsgen.Sensor, error) {
	sensors := make([]*gpsgen.Sensor, 0, len(t.sensors))
	for _, conf := range t.sensors {
		sensor, err := t.makeSensorFromConf(conf)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}
