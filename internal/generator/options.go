package generator

import (
	"fmt"

	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/gpsgend/internal/types"
)

type UpdateTrackerOptions struct {
	Model  *types.Model       // Model of the device.
	Color  *types.Color       // Color of the device.
	UserID *types.CustomID    // User ID associated with the device.
	Descr  *types.Description // Description of the device.
}

func (o UpdateTrackerOptions) isEmpty() bool {
	return o.Model == nil &&
		o.Color == nil &&
		o.UserID == nil &&
		o.Descr == nil
}

func (o UpdateTrackerOptions) validate() error {
	if o.Model != nil {
		if err := validateType(*o.Model, "tracker.model"); err != nil {
			return err
		}
	}
	if o.Color != nil {
		if err := validateType(*o.Color, "tracker.color"); err != nil {
			return err
		}
	}
	if o.UserID != nil {
		if err := validateType(*o.UserID, "tracker.userID"); err != nil {
			return err
		}
	}
	if o.Descr != nil {
		if err := validateType(*o.Descr, "tracker.description"); err != nil {
			return err
		}
	}
	return nil
}

type NewTrackerOptions struct {
	Model  *types.Model       // Model of the device.
	Color  *types.Color       // Color of the device.
	UserID *types.CustomID    // User ID associated with the device.
	Descr  *types.Description // Description of the device.
	Props  *properties.Properties

	SkipOffline bool // Skip offline mode.
	Offline     types.Offline
	Elevation   types.Elevation
	Battery     types.Battery
	Speed       types.Speed
}

type validater interface {
	Validate() error
}

var _ = []validater{
	types.Speed{}, types.Elevation{}, types.Battery{},
	types.Offline{}, types.Description{}, types.CustomID{},
	types.Color{}, types.Model{}, types.DeviceStatus(0),
}

func validateType(t any, ctx string) error {
	v, ok := t.(validater)
	if !ok {
		return nil
	}
	if err := v.Validate(); err != nil {
		return fmt.Errorf("%w: %s", err, ctx)
	}
	return nil
}
