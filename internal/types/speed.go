package types

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen"
)

const (
	MinSpeedValue = 1
	MaxSpeedValue = 100
)

type Speed struct {
	min       float64
	max       float64
	amplitude int
}

func NewSpeed(min float64, max float64, amplitude int) (Speed, error) {
	speed := Speed{min: min, max: max, amplitude: amplitude}
	if err := speed.validate(); err != nil {
		return Speed{}, err
	}
	return speed, nil
}

func (s Speed) String() string {
	return fmt.Sprintf("Speed{min:%.2f, max:%.2f, amplitude:%d}",
		s.min, s.max, s.amplitude)
}

func (s Speed) Min() float64 {
	return s.min
}

func (s Speed) Max() float64 {
	return s.max
}

func (s Speed) Amplitude() int {
	return s.amplitude
}

func (s Speed) validate() error {
	if s.min < MinSpeedValue {
		return ErrInvalidMinValue
	}
	if s.max > MaxSpeedValue {
		return ErrInvalidMaxValue
	}
	if s.min > s.max {
		return ErrInvalidRangeValue
	}
	if s.amplitude < gpsgen.Amplitude4 {
		return ErrInvalidMinAmplitude
	}
	if s.amplitude > gpsgen.Amplitude512 {
		return ErrInvalidMaxAmplitude
	}
	return nil
}
