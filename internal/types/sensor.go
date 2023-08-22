package types

import (
	"fmt"

	stdtypes "github.com/mmadfox/go-gpsgen/types"
)

type SensorMode = stdtypes.SensorMode

type Sensor struct {
	id        ID
	name      string
	min, max  float64
	amplitude int
	mode      SensorMode
}

func ParseSensor(id ID, name string, min, max float64, amplitude int, mode int) (*Sensor, error) {
	sensor := Sensor{
		id:        id,
		name:      name,
		min:       min,
		max:       max,
		amplitude: amplitude,
		mode:      SensorMode(mode),
	}
	if err := sensor.Validate(); err != nil {
		return nil, err
	}
	return &sensor, nil
}

func NewSensor(name string, min, max float64, amplitude int, mode int) (*Sensor, error) {
	return ParseSensor(NewID(), name, min, max, amplitude, mode)
}

func (s Sensor) String() string {
	return fmt.Sprintf("Sensor{id:%s, name:%s, min:%.6f, max: %.6f, amplitude: %d}",
		s.id, s.name, s.min, s.max, s.amplitude)
}

func (s Sensor) ID() ID {
	return s.id
}

func (s Sensor) Name() string {
	return s.name
}

func (s Sensor) Min() float64 {
	return s.min
}

func (s Sensor) Max() float64 {
	return s.max
}

func (s Sensor) Amplitude() int {
	return s.amplitude
}

func (s Sensor) Mode() SensorMode {
	return s.mode
}

func (s Sensor) Validate() error {
	if len(s.id.value) == 0 {
		return ErrInvalidID
	}

	if len(s.name) == 0 {
		return ErrInvalidName
	}

	if s.amplitude < 4 {
		return ErrInvalidMinAmplitude
	}

	if s.amplitude > 512 {
		return ErrInvalidMaxAmplitude
	}
	return nil
}
