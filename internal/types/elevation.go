package types

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen/types"
)

const (
	MinElevationValue = 0
	MaxElevationValue = 10000
)

type Elevation struct {
	min       float64
	max       float64
	amplitude int
	mode      types.SensorMode
}

func ParseElevation(min, max float64, amplitude int, mode types.SensorMode) (Elevation, error) {
	elevation := Elevation{min: min, max: max, amplitude: amplitude, mode: mode}
	if err := elevation.Validate(); err != nil {
		return Elevation{}, err
	}
	return elevation, nil
}

func (e Elevation) String() string {
	return fmt.Sprintf("Elevation{min:%.2f, max:%.2f, amplitude:%d}",
		e.min, e.max, e.amplitude)
}

func (e Elevation) Min() float64 {
	return e.min
}

func (e Elevation) Max() float64 {
	return e.max
}

func (e Elevation) Amplitude() int {
	return e.amplitude
}

func (e Elevation) Mode() types.SensorMode {
	return e.mode
}

func (e Elevation) Validate() error {
	if e.min < MinElevationValue {
		return ErrInvalidMinValue
	}
	if e.max > MaxElevationValue {
		return ErrInvalidMaxValue
	}
	if e.min > e.max {
		return ErrInvalidRangeValue
	}
	if e.amplitude < 4 {
		return ErrInvalidMinAmplitude
	}
	if e.amplitude > 512 {
		return ErrInvalidMaxAmplitude
	}
	return nil
}
