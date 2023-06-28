package stubservice

import (
	"time"

	"github.com/mmadfox/gpsgend/internal/device"
)

func UpdateDeviceParams() device.UpdateDeviceParams {
	params := device.UpdateDeviceParams{
		Model:       str2ptr("model"),
		Description: str2ptr("description"),
		UserID:      str2ptr("userID"),
		Properties:  &map[string]string{"foo": "bar"},
	}
	params.Speed = &struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	}{
		Min:       1,
		Max:       7,
		Amplitude: 4,
	}
	params.Battery = &struct {
		Max        float64 `json:"max"`
		Min        float64 `json:"min"`
		ChargeTime int64   `json:"chargeTime"`
	}{
		Min:        1,
		Max:        100,
		ChargeTime: int64(time.Hour.Seconds()),
	}
	params.Elevation = &struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	}{
		Min:       1,
		Max:       100,
		Amplitude: 4,
	}
	params.Offline = &struct {
		Min int `json:"min"`
		Max int `json:"max"`
	}{
		Min: 1,
		Max: 2,
	}
	return params
}

func str2ptr(str string) *string {
	return &str
}
