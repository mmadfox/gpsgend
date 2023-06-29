package device

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/types"
)

type UseCase interface {
	NewPresetsDevice(ctx context.Context, p NewPresetsDeviceParams) (*Device, error)
	NewDevice(ctx context.Context, p NewDeviceParams) (*Device, error)
	UpdateDevice(ctx context.Context, deviceID uuid.UUID, p UpdateDeviceParams) (*Device, error)
	RemoveDevice(ctx context.Context, deviceID uuid.UUID) error
	RunDevice(ctx context.Context, deviceID uuid.UUID) error
	StopDevice(ctx context.Context, deviceID uuid.UUID) error
	PauseDevice(ctx context.Context, deviceID uuid.UUID) error
	ResumeDevice(ctx context.Context, deviceID uuid.UUID) error
	AddRoutes(ctx context.Context, deviceID uuid.UUID, routes []*navigator.Route) error
	Routes(ctx context.Context, deviceID uuid.UUID) ([]*Route, error)
	RemoveRoute(ctx context.Context, deviceID uuid.UUID, routeID uuid.UUID) error
	Sensors(ctx context.Context, deviceID uuid.UUID) ([]types.Sensor, error)
	AddSensor(ctx context.Context, deviceID uuid.UUID, p AddSensorParams) ([]types.Sensor, error)
	RemoveSensor(ctx context.Context, deviceID uuid.UUID, sensorID uuid.UUID) error
	DeviceByID(ctx context.Context, deviceID uuid.UUID) (*Device, error)
}

type Middleware func(UseCase) UseCase

type NewDeviceParams struct {
	Model       string            `json:"model"`
	UserID      string            `json:"userId,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Description string            `json:"description,omitempty"`
	Speed       struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Battery struct {
		Max        float64 `json:"max"`
		Min        float64 `json:"min"`
		ChargeTime int64   `json:"chargeTime"`
	} `json:"battery"`
	Elevation struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"elevation"`
	Offline struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
	WithRandomRoutes bool `json:"withRandomRoutes"`
}

type NewPresetsDeviceParams struct {
	UserID           string            `json:"userId,omitempty"`
	Description      string            `json:"description,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
	WithRandomRoutes bool              `json:"withRandomRoutes"`
}

type UpdateDeviceParams struct {
	Model       *string            `json:"model"`
	Description *string            `json:"description"`
	UserID      *string            `json:"userId,omitempty"`
	Properties  *map[string]string `json:"properties,omitempty"`
	Speed       *struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Battery *struct {
		Max        float64 `json:"max"`
		Min        float64 `json:"min"`
		ChargeTime int64   `json:"chargeTime"`
	} `json:"battery"`
	Elevation *struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"elevation"`
	Offline *struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
}

type AddSensorParams struct {
	Name      string  `json:"name"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Amplitude int     `json:"amplitude"`
}
