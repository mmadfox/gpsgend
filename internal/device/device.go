package device

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/types"
)

const (
	MaxNumSensorsPerDevice = 10
	MaxNumRoutesPerDevice  = 64
)

type Device struct {
	id        uuid.UUID
	color     colorful.Color
	model     types.Model
	userID    string
	descr     types.Description
	speed     types.Speed
	battery   types.Battery
	elevation types.Elevation
	offline   types.Offline
	props     types.Properties
	status    Status
	sensors   []types.Sensor
	routes    []*Route
	snapshot  []byte
	createdAt time.Time
	updatedAt time.Time
	version   int
}

func (d *Device) ID() uuid.UUID {
	return d.id
}

func (d *Device) Model() types.Model {
	return d.model
}

func (d *Device) Color() colorful.Color {
	return d.color
}

func (d *Device) UserID() string {
	return d.userID
}

func (d *Device) Description() types.Description {
	return d.descr
}

func (d *Device) Speed() types.Speed {
	return d.speed
}

func (d *Device) Battery() types.Battery {
	return d.battery
}

func (d *Device) Elevation() types.Elevation {
	return d.elevation
}

func (d *Device) Offline() types.Offline {
	return d.offline
}

func (d *Device) Props() types.Properties {
	return d.props
}

func (d *Device) Status() Status {
	return d.status
}

func (d *Device) Snapshot() []byte {
	return d.snapshot
}

func (d *Device) CreatedAt() time.Time {
	return d.createdAt
}

func (d *Device) UpdatedAt() time.Time {
	return d.updatedAt
}

func (d *Device) IsRunning() bool {
	return d.status == Running
}

func (d *Device) AssignColor() {
	d.color = randomColor()
}

func (d *Device) ResumeDevice() (*gpsgen.Device, error) {
	if d.status == Running {
		return nil, ErrDeviceAlreadyRunning
	}

	if len(d.snapshot) == 0 {
		return nil, ErrNoSnapshot
	}

	proc, err := gpsgen.DeviceFromSnapshot(d.snapshot)
	if err != nil {
		return nil, err
	}

	d.status = Running
	d.snapshot = nil

	return proc, nil
}

func (d *Device) PauseProcess(proc *gpsgen.Device) error {
	if d.status != Running {
		return ErrDeviceAlreadyStopped
	}

	snapshot, err := gpsgen.TakeDeviceSnapshot(proc)
	if err != nil {
		return err
	}

	d.status = Paused
	d.snapshot = snapshot
	return nil
}

func (d *Device) StoreProcess(proc *gpsgen.Device) error {
	if d.status != Running {
		return ErrDeviceAlreadyStopped
	}

	snapshot, err := gpsgen.TakeDeviceSnapshot(proc)
	if err != nil {
		return err
	}

	d.status = Stored
	d.snapshot = snapshot
	return nil
}

func (d *Device) RemoveProcess() error {
	if d.status != Running {
		return ErrDeviceAlreadyStopped
	}

	d.status = Stopped
	return nil
}

func (d *Device) NewProcess() (*gpsgen.Device, error) {
	if d.status == Running {
		return nil, ErrDeviceAlreadyRunning
	}

	if len(d.routes) == 0 {
		return nil, ErrNoRoutes
	}

	d.status = Running

	conf := gpsgen.NewConfig()
	conf.Routes = make([]*navigator.Route, len(d.routes))
	for i := 0; i < len(d.routes); i++ {
		conf.Routes[i] = d.routes[i].Route()
	}
	conf.ID = d.id.String()
	conf.Model = d.model.String()
	conf.UserID = d.userID
	conf.Properties = d.props.ToMap()
	conf.Description = d.descr.String()
	conf.Battery.Min = d.battery.Min()
	conf.Battery.Max = d.battery.Max()
	conf.Battery.ChargeTime = d.battery.ChargeTime()
	conf.Speed.Min = d.speed.Min()
	conf.Speed.Max = d.speed.Max()
	conf.Speed.Amplitude = d.speed.Amplitude()
	conf.Elevation.Min = d.elevation.Min()
	conf.Elevation.Max = d.elevation.Max()
	conf.Elevation.Amplitude = d.elevation.Amplitude()
	conf.Offline.Min = d.offline.Min()
	conf.Offline.Max = d.offline.Max()

	if len(d.sensors) > 0 {
		conf.Sensors = make([]gpsgen.Sensor, len(d.sensors))
		for i := 0; i < len(d.sensors); i++ {
			sensor := d.sensors[i]
			conf.Sensors[i] = gpsgen.Sensor{
				Name:      sensor.Name(),
				Min:       sensor.Min(),
				Max:       sensor.Max(),
				Amplitude: sensor.Amplitude(),
			}
		}
	}

	proc, err := conf.NewDevice()
	if err != nil {
		return nil, err
	}

	return proc, nil
}

func (d *Device) Sensors() []types.Sensor {
	sensors := make([]types.Sensor, len(d.sensors))
	copy(sensors, d.sensors)
	return sensors
}

func (d *Device) Routes() []*Route {
	routes := make([]*Route, len(d.routes))
	copy(routes, d.routes)
	return routes
}

func (d *Device) ChangeModel(val types.Model) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.model = val
	return nil
}

func (d *Device) ChangeDescription(val types.Description) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.descr = val
	return nil
}

func (d *Device) ChangeUserID(userID string) {
	d.userID = userID
}

func (d *Device) ChangeSpeed(val types.Speed) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.speed = val
	return nil
}

func (d *Device) ChangeProps(props types.Properties) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.props = props
	return nil
}

func (d *Device) ChangeBattery(val types.Battery) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.battery = val
	return nil
}

func (d *Device) ChangeElevation(val types.Elevation) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.elevation = val
	return nil
}

func (d *Device) ChangeOfflineDuraiton(val types.Offline) error {
	if err := d.checkState(); err != nil {
		return err
	}

	d.offline = val
	return nil
}

func (d *Device) RemoveSensor(sensorID uuid.UUID) error {
	// TODO:
	return nil
}

func (d *Device) AddSensor(name string, min float64, max float64, amplitude int) error {
	if err := d.checkState(); err != nil {
		return err
	}

	if len(d.sensors)+1 > MaxNumSensorsPerDevice {
		return ErrSensorLimitExceeded
	}

	newSensor, err := types.NewSensor(name, min, max, amplitude)
	if err != nil {
		return err
	}

	d.sensors = append(d.sensors, newSensor)
	return nil
}

func (d *Device) AddRoute(r *Route) error {
	if err := d.checkState(); err != nil {
		return err
	}

	if len(d.routes)+1 > MaxNumRoutesPerDevice {
		return ErrRoutesLimitExceeded
	}

	d.routes = append(d.routes, r)
	return nil
}

func (d *Device) RemoveRoute(routeID uuid.UUID) error {
	// TODO:
	return nil
}

func (d *Device) checkState() error {
	if d.status == Running || d.status == Paused {
		return ErrSetConfig
	}
	return nil
}

var colorRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomColor() colorful.Color {
	target := colorful.Color{
		R: 243, G: 243, B: 243,
	}

	for i := 0; i < 3; i++ {
		h := colorRand.Float64() * 360.0
		s := 0.7 + colorRand.Float64()*0.3
		v := 0.6 + colorRand.Float64()*0.3
		color := colorful.Hsv(h, s, v)
		if !color.IsValid() {
			continue
		}
		target = color
		break
	}

	return target
}

func (d *Device) restoreRoutes(routes []*Route) error {
	if len(routes) > MaxNumRoutesPerDevice {
		return ErrRoutesLimitExceeded
	}

	d.routes = routes
	return nil
}

func (d *Device) restoreSensors(sensors []types.Sensor) error {
	if len(sensors) > MaxNumSensorsPerDevice {
		return ErrSensorLimitExceeded
	}

	d.sensors = sensors
	return nil
}
