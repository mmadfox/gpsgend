package http

import (
	"github.com/mmadfox/go-gpsgen/geojson"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/device"
)

func toDeviceModel(dev *device.Device) *deviceModel {
	model := deviceModel{
		ID:          dev.ID().String(),
		UserID:      dev.UserID(),
		Color:       dev.Color().Hex(),
		Model:       dev.Model().String(),
		Description: dev.Description().String(),
		Sensors:     zeroSensors,
	}

	model.Status.ID = int(dev.Status())
	model.Status.Text = dev.Status().String()

	speed := dev.Speed()
	model.Speed.Min = speed.Min()
	model.Speed.Max = speed.Max()
	model.Speed.Amplitude = speed.Amplitude()

	battery := dev.Battery()
	model.Battery.Min = battery.Min()
	model.Battery.Max = battery.Max()
	model.Battery.ChargeTime = battery.ChargeTime()

	elevation := dev.Elevation()
	model.Elevation.Min = elevation.Min()
	model.Elevation.Max = elevation.Max()
	model.Elevation.Amplitude = elevation.Amplitude()

	offline := dev.Offline()
	model.Offline.Min = offline.Min()
	model.Offline.Max = offline.Min()

	model.Props = dev.Props().ToMap()
	sensors := dev.Sensors()
	if len(sensors) > 0 {
		model.Sensors = make([]sensorModel, len(sensors))
		for i := 0; i < len(sensors); i++ {
			model.Sensors[i] = sensorModel{
				ID:        sensors[i].ID().String(),
				Name:      sensors[i].Name(),
				Min:       sensors[i].Min(),
				Max:       sensors[i].Max(),
				Amplitude: sensors[i].Amplitude(),
			}
		}
	}

	return &model
}

func toRouteModel(r *device.Route) (routeModel, error) {
	rawRoutes, err := geojson.Encode([]*navigator.Route{r.Route()})
	if err != nil {
		return routeModel{}, err
	}
	return routeModel{
		ID:     r.ID().String(),
		Color:  r.Color().Hex(),
		Routes: rawRoutes,
	}, nil
}
