package types

import (
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
)

const MaxNameLenValue = 512

type Sensor struct {
	id        uuid.UUID
	name      string
	min       float64
	max       float64
	amplitude int
}

func NewSensor(name string, min float64, max float64, amplitude int) (Sensor, error) {
	s := Sensor{id: uuid.New(), name: name, min: min, max: max, amplitude: amplitude}
	if err := s.validate(); err != nil {
		return Sensor{}, err
	}
	return s, nil
}

func SensorFrom(id uuid.UUID, name string, min float64, max float64, amplitude int) (Sensor, error) {
	s := Sensor{id: id, name: name, min: min, max: max, amplitude: amplitude}
	if err := s.validate(); err != nil {
		return Sensor{}, nil
	}
	return s, nil
}

func (s Sensor) ID() uuid.UUID {
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

func (s Sensor) validate() error {
	if len(s.name) == 0 {
		return ErrInvalidName
	}
	if len(s.name) > MaxNameLenValue {
		return ErrInvalidName
	}
	if s.amplitude < gpsgen.Amplitude4 {
		return ErrInvalidMinAmplitude
	}
	if s.amplitude > gpsgen.Amplitude512 {
		return ErrInvalidMaxAmplitude
	}
	return nil
}
