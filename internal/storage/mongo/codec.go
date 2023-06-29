package mongo

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func encodeDevice(d *device.Device) (*deviceModel, error) {
	model := deviceModel{
		ID:          d.ID().String(),
		UserID:      d.UserID(),
		Model:       d.Model().String(),
		Color:       d.Color().Hex(),
		Description: d.Description().String(),
		Snapshot:    d.Snapshot(),
	}

	model.Speed.Amplitude = d.Speed().Amplitude()
	model.Speed.Min = d.Speed().Min()
	model.Speed.Max = d.Speed().Max()

	model.Battery.ChargeTime = d.Battery().ChargeTime().String()
	model.Battery.Min = d.Battery().Min()
	model.Battery.Max = d.Battery().Max()

	model.Elevation.Amplitude = d.Elevation().Amplitude()
	model.Elevation.Min = d.Elevation().Min()
	model.Elevation.Max = d.Elevation().Max()

	model.Offline.Min = d.Offline().Min()
	model.Offline.Max = d.Offline().Max()

	model.Props = d.Props().ToMap()
	model.Status.ID = int(d.Status())
	model.Status.Text = d.Status().String()

	sensors := d.Sensors()
	model.NumSensors = len(sensors)
	model.Sensors = make([]sensorModel, len(sensors))
	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		model.Sensors[i] = sensorModel{
			ID:        sensor.ID().String(),
			Name:      sensor.Name(),
			Min:       sensor.Min(),
			Max:       sensor.Max(),
			Amplitude: sensor.Amplitude(),
		}
	}

	routes := d.Routes()
	model.NumRoutes = len(routes)
	model.Routes = make([]routeModel, len(routes))
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		rawRoute, err := route.Route().MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("route marshal error: %w", err)
		}
		model.Routes[i] = routeModel{
			ID:    route.ID().String(),
			Color: route.Color().Hex(),
			Route: rawRoute,
		}
	}

	model.CreatedAt = d.CreatedAt().UnixNano()
	model.UpdatedAt = d.UpdatedAt().UnixNano()
	model.Version = device.Version(d)

	return &model, nil
}

func decodeDevice(d *deviceModel) (*device.Device, error) {
	builder := device.NewBuilder()

	id, err := uuid.Parse(d.ID)
	if err != nil {
		return nil, err
	}
	builder.ID(id)
	builder.UserID(d.UserID)
	builder.Model(d.Model)

	color, err := colorful.Hex(d.Color)
	if err != nil {
		return nil, err
	}

	batteryChargeTime, err := time.ParseDuration(d.Battery.ChargeTime)
	if err != nil {
		return nil, err
	}

	builder.Color(color)
	builder.Description(d.Description)
	builder.Speed(d.Speed.Min, d.Speed.Max, d.Speed.Amplitude)
	builder.Battery(d.Battery.Min, d.Battery.Max, batteryChargeTime)
	builder.Elevation(d.Elevation.Min, d.Elevation.Max, d.Elevation.Amplitude)
	builder.Offline(d.Offline.Min, d.Offline.Max)
	builder.Props(d.Props)
	builder.Status(device.Status(d.Status.ID))
	builder.Snapshot(d.Snapshot)

	if len(d.Sensors) > 0 {
		sensors := make([]types.Sensor, len(d.Sensors))
		for i := 0; i < len(d.Sensors); i++ {
			sid, err := uuid.Parse(d.Sensors[i].ID)
			if err != nil {
				return nil, err
			}
			sensor, err := types.SensorFrom(
				sid,
				d.Sensors[i].Name,
				d.Sensors[i].Min,
				d.Sensors[i].Max,
				d.Sensors[i].Amplitude,
			)
			if err != nil {
				return nil, err
			}
			sensors[i] = sensor
		}
		builder.Sensors(sensors)
	}

	if len(d.Routes) > 0 {
		routes := make([]*device.Route, len(d.Routes))
		for i := 0; i < len(d.Routes); i++ {
			originRoute := new(navigator.Route)
			if err := originRoute.UnmarshalBinary(d.Routes[i].Route); err != nil {
				return nil, err
			}
			rid, err := uuid.Parse(d.Routes[i].ID)
			if err != nil {
				return nil, err
			}
			rc, err := colorful.Hex(d.Routes[i].Color)
			if err != nil {
				return nil, err
			}
			routes[i] = device.RouteFrom(rid, rc, originRoute)
		}
		builder.Routes(routes)
	}

	if d.CreatedAt > 0 {
		builder.CreatedAt(time.Unix(0, d.CreatedAt))
	}

	if d.UpdatedAt > 0 {
		builder.UpdatedAt(time.Unix(0, d.UpdatedAt))
	}

	dev, err := builder.Build()
	if err != nil {
		return nil, err
	}

	device.SetVersion(dev, d.Version)

	return dev, nil
}

func encodeQueryFilter(qf *device.QueryFilter) bson.D {
	filter := bson.D{}

	if qf.Model != nil {
		filter = append(filter, bson.E{
			Key: "model",
			Value: bson.M{
				"$regex": primitive.Regex{Pattern: "^" + *qf.Model, Options: "i"},
			},
		})
	}

	if qf.Status != nil {
		filter = append(filter, bson.E{
			Key:   "status.id",
			Value: bson.M{"$in": *qf.Status},
		})
	}

	if qf.ID != nil {
		filter = append(filter, bson.E{
			Key:   "device_id",
			Value: bson.M{"$in": *qf.ID},
		})
	}

	if qf.User != nil {
		filter = append(filter, bson.E{
			Key:   "user_id",
			Value: bson.M{"$in": *qf.User},
		})
	}

	if qf.Sensor != nil {
		filter = append(filter, bson.E{
			Key:   "sensors.id",
			Value: bson.M{"$in": *qf.Sensor},
		})
	}

	return filter
}
