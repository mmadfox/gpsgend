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

func ParseOffline(min int, max int) (Offline, error) {
	offline := Offline{min: min, max: max}
	if err := offline.Validate(); err != nil {
		return Offline{}, err
	}
	return offline, nil
}

func (o Offline) IsEmpty() bool {
	return o.min == 0 && o.max == 0
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

func (o Offline) Validate() error {
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
