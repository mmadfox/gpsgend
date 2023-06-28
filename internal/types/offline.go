package types

import "fmt"

const (
	MinOfflineValue = 0
	MaxOfflineValue = 900
)

type Offline struct {
	min int
	max int
}

func NewOffline(min int, max int) (Offline, error) {
	offline := Offline{min: min, max: max}
	if err := offline.validate(); err != nil {
		return Offline{}, err
	}
	return offline, nil
}

func (o Offline) String() string {
	return fmt.Sprintf("Offline{min:%d, max:%d}", o.min, o.max)
}

func (o Offline) Min() int {
	return o.min
}

func (o Offline) Max() int {
	return o.max
}

func (o Offline) validate() error {
	if o.min < MinOfflineValue {
		return ErrInvalidMinValue
	}
	if o.max > MaxOfflineValue {
		return ErrInvalidMaxValue
	}
	if o.min > o.max {
		return ErrInvalidRangeValue
	}
	return nil
}
