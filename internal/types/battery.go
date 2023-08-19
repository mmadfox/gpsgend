package types

import (
	"fmt"
	"time"
)

const (
	MinBatteryValue      = 0
	MaxBatteryValue      = 100
	MinBatteryChargeTime = time.Hour
	MaxBatteryChargeTime = 10 * time.Hour
)

type Battery struct {
	min        float64
	max        float64
	chargeTime time.Duration
}

func ParseBattery(min float64, max float64, chargeTime time.Duration) (Battery, error) {
	battery := Battery{
		min:        min,
		max:        max,
		chargeTime: chargeTime,
	}
	if err := battery.Validate(); err != nil {
		return Battery{}, err
	}
	return battery, nil
}

func (b Battery) String() string {
	return fmt.Sprintf("Battery{min:%.2f, max:%.2f, chargeTime:%s}",
		b.min, b.max, b.chargeTime)
}

func (b Battery) Min() float64 {
	return b.min
}

func (b Battery) Max() float64 {
	return b.max
}

func (b Battery) ChargeTime() time.Duration {
	return b.chargeTime
}

func (b Battery) Validate() error {
	if b.min < MinBatteryValue {
		return ErrInvalidMinValue
	}
	if b.max > MaxBatteryValue {
		return ErrInvalidMaxValue
	}
	if b.min > b.max {
		return ErrInvalidRangeValue
	}
	if b.chargeTime < MinBatteryChargeTime {
		return ErrInvalidMinChargeTime
	}
	if b.chargeTime > MaxBatteryChargeTime {
		return ErrInvalidMaxChargeTime
	}
	return nil
}
