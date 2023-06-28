package device

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mmadfox/gpsgend/internal/types"
)

type Builder struct {
	device *Device
	err    error
}

func NewBuilder() *Builder {
	return &Builder{
		device: &Device{
			sensors: make([]types.Sensor, 0),
			routes:  make([]*Route, 0),
		},
	}
}

func (db *Builder) ID(id uuid.UUID) *Builder {
	db.device.id = id
	return db
}

func (db *Builder) Color(color colorful.Color) *Builder {
	db.device.color = color
	return db
}

func (db *Builder) Model(val string) *Builder {
	model, err := types.NewModel(val)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device model: %w", err))
		return db
	}
	db.device.model = model
	return db
}

func (db *Builder) Props(val map[string]string) *Builder {
	props, err := types.NewProperties(val)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device props: %w", err))
		return db
	}
	db.device.props = props
	return db
}

func (db *Builder) Description(val string) *Builder {
	descr, err := types.NewDescription(val)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device description: %w", err))
		return db
	}
	db.device.descr = descr
	return db
}

func (db *Builder) Speed(min, max float64, amplitude int) *Builder {
	speed, err := types.NewSpeed(min, max, amplitude)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device speed: %w", err))
		return db
	}
	db.device.speed = speed
	return db
}

func (db *Builder) Battery(min, max float64, chargeTime time.Duration) *Builder {
	battery, err := types.NewBattery(min, max, chargeTime)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device battery: %w", err))
		return db
	}
	db.device.battery = battery
	return db
}

func (db *Builder) Elevation(min, max float64, amplitude int) *Builder {
	elevation, err := types.NewElevation(min, max, amplitude)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device elevation: %w", err))
		return db
	}
	db.device.elevation = elevation
	return db
}

func (db *Builder) Offline(min, max int) *Builder {
	offline, err := types.NewOffline(min, max)
	if err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device offline: %w", err))
		return db
	}
	db.device.offline = offline
	return db
}

func (db *Builder) UserID(uid string) *Builder {
	db.device.userID = uid
	return db
}

func (db *Builder) Status(s Status) *Builder {
	if err := validateStatus(s); err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device status: %w", err))
	}
	db.device.status = s
	return db
}

func (db *Builder) Snapshot(data []byte) *Builder {
	db.device.snapshot = data
	return db
}

func (db *Builder) Sensors(sensors []types.Sensor) *Builder {
	if err := db.device.restoreSensors(sensors); err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device sensors: %w", err))
	}
	return db
}

func (db *Builder) Routes(routes []*Route) *Builder {
	if err := db.device.restoreRoutes(routes); err != nil {
		db.err = errors.Join(db.err, fmt.Errorf("device routes: %w", err))
	}
	return db
}

func (db *Builder) CreatedAt(t time.Time) *Builder {
	db.device.createdAt = t
	return db
}

func (db *Builder) UpdatedAt(t time.Time) *Builder {
	db.device.updatedAt = t
	return db
}

func (db *Builder) Build() (*Device, error) {
	if db.err != nil {
		return nil, db.err
	}
	return db.device, nil
}
