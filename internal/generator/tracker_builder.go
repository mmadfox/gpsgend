package generator

import (
	"errors"
	"time"

	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/gpsgend/internal/types"
)

type TrackerBuilder struct {
	errors error
	device *Tracker
}

func NewTrackerBuilder() *TrackerBuilder {
	return &TrackerBuilder{
		device: &Tracker{
			sensors: make(map[types.ID]*types.Sensor),
		},
	}
}

func (b *TrackerBuilder) ID(id types.ID) *TrackerBuilder {
	if err := validateType(id, "can't parse type tracker.id"); err != nil {
		b.appendError(err)
	} else {
		b.device.id = id
	}
	return b
}

func (b *TrackerBuilder) Status(status types.DeviceStatus) *TrackerBuilder {
	if err := validateType(status, "can't parse type tracker.status"); err != nil {
		b.appendError(err)
	} else {
		b.device.status = status
	}
	return b
}

func (b *TrackerBuilder) Model(model types.Model) *TrackerBuilder {
	if err := validateType(model, "can't parse type tracker.model"); err != nil {
		b.appendError(err)
	} else {
		b.device.model = model
	}
	return b
}

func (b *TrackerBuilder) Color(color types.Color) *TrackerBuilder {
	if err := validateType(color, "can't parse type tracker.color"); err != nil {
		b.appendError(err)
	} else {
		b.device.color = color
	}
	return b
}

func (b *TrackerBuilder) CustomID(id types.CustomID) *TrackerBuilder {
	if err := validateType(id, "can't parse type tracker.userID"); err != nil {
		b.appendError(err)
	} else {
		b.device.userID = id
	}
	return b
}

func (b *TrackerBuilder) Props(p properties.Properties) *TrackerBuilder {
	if b.device.props == nil {
		b.device.props = properties.Make()
	}
	b.device.props.Merge(p)
	return b
}

func (b *TrackerBuilder) Description(descr types.Description) *TrackerBuilder {
	if err := validateType(descr, "can't parse type tracker.description"); err != nil {
		b.appendError(err)
	} else {
		b.device.description = descr
	}
	return b
}

func (b *TrackerBuilder) Offline(offline types.Offline) *TrackerBuilder {
	if err := validateType(offline, "can't parse type offline"); err != nil {
		b.appendError(err)
	} else {
		b.device.offline = offline
	}
	return b
}

func (b *TrackerBuilder) SkipOffline(flag bool) *TrackerBuilder {
	b.device.skipOffline = flag
	return b
}

func (b *TrackerBuilder) Elevation(elevation types.Elevation) *TrackerBuilder {
	if err := validateType(elevation, "can't parse type tracker.elevation"); err != nil {
		b.appendError(err)
	} else {
		b.device.elevation = elevation
	}
	return b
}

func (b *TrackerBuilder) Battery(battery types.Battery) *TrackerBuilder {
	if err := validateType(battery, "can't parse type tracker.battery"); err != nil {
		b.appendError(err)
	} else {
		b.device.battery = battery
	}
	return b
}

func (b *TrackerBuilder) Speed(speed types.Speed) *TrackerBuilder {
	if err := validateType(speed, "can't parse type tracker.speed"); err != nil {
		b.appendError(err)
	} else {
		b.device.speed = speed
	}
	return b
}

func (b *TrackerBuilder) CreatedAt(t time.Time) *TrackerBuilder {
	b.device.createdAt = t
	return b
}

func (b *TrackerBuilder) UpdatedAt(t time.Time) *TrackerBuilder {
	b.device.updatedAt = t
	return b
}

func (b *TrackerBuilder) RunningAt(t time.Time) *TrackerBuilder {
	b.device.runningAt = t
	return b
}

func (b *TrackerBuilder) StoppedAt(t time.Time) *TrackerBuilder {
	b.device.stoppedAt = t
	return b
}

func (b *TrackerBuilder) Build() (*Tracker, error) {
	if b.errors != nil {
		return nil, b.errors
	}
	return b.device, nil
}

func (b *TrackerBuilder) appendError(err error) {
	b.errors = errors.Join(b.errors, err)
}
