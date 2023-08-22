package grpc

import (
	"encoding/json"
	"errors"

	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func sensors2model(sensors []*types.Sensor) []*Sensor {
	sensorsModel := make([]*Sensor, len(sensors))
	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		sensorsModel[i] = &Sensor{
			Id:        sensor.ID().String(),
			Name:      sensor.Name(),
			Min:       sensor.Min(),
			Max:       sensor.Max(),
			Amplitude: int64(sensor.Amplitude()),
			Mode:      int64(sensor.Mode()),
		}
	}
	return sensorsModel
}

func trackerView2model(v *generator.TrackerView) (*Tracker, error) {
	var rawProps []byte
	if len(v.Props) > 0 {
		props, err := json.Marshal(v.Props)
		if err != nil {
			return nil, err
		}
		rawProps = props
	}
	return &Tracker{
		Id:       v.ID,
		CustomId: v.CustomID,
		Status: &Status{
			Id:   int64(v.Status.ID),
			Name: v.Status.Name,
		},
		Model: v.Model,
		Color: v.Color,
		Descr: v.Descr,
		Offline: &Offline{
			Min: int64(v.Offline.Min),
			Max: int64(v.Offline.Max),
		},
		Elevation: &Elevation{
			Min:       v.Elevation.Min,
			Max:       v.Elevation.Max,
			Amplitude: int64(v.Elevation.Amplitude),
			Mode:      int64(v.Elevation.Mode),
		},
		Battery: &Battery{
			Min:        v.Battery.Min,
			Max:        v.Battery.Max,
			ChargeTime: int64(v.Battery.ChargeTime),
		},
		Speed: &Speed{
			Min:       v.Speed.Min,
			Max:       v.Speed.Max,
			Amplitude: int64(v.Speed.Amplitude),
		},
		Props:       rawProps,
		NumSensors:  int64(v.NumSensors),
		NumRoutes:   int64(v.NumRoutes),
		SkipOffline: v.SkipOffline,
		CreatedAt:   v.CreatedAt.Unix(),
		UpdatedAt:   v.UpdatedAt.Unix(),
		RunningAt:   v.RunningAt.Unix(),
		StoppedAt:   v.StoppedAt.Unix(),
	}, nil
}

func tracker2model(t *generator.Tracker) (*Tracker, error) {
	var rawProps []byte
	if len(t.Properties()) > 0 {
		props, err := json.Marshal(t.Properties())
		if err != nil {
			return nil, err
		}
		rawProps = props
	}
	return &Tracker{
		Id:       t.ID().String(),
		CustomId: t.UserID().String(),
		Status: &Status{
			Id:   int64(t.Status()),
			Name: t.Status().String(),
		},
		Model: t.Model().String(),
		Color: t.Color().String(),
		Descr: t.Description().String(),
		Offline: &Offline{
			Min: int64(t.Offline().Min()),
			Max: int64(t.Offline().Max()),
		},
		Elevation: &Elevation{
			Min:       t.Elevation().Min(),
			Max:       t.Elevation().Max(),
			Amplitude: int64(t.Elevation().Amplitude()),
			Mode:      int64(t.Elevation().Mode()),
		},
		Battery: &Battery{
			Min:        t.Battery().Min(),
			Max:        t.Battery().Max(),
			ChargeTime: int64(t.Battery().ChargeTime()),
		},
		Speed: &Speed{
			Min:       t.Speed().Min(),
			Max:       t.Speed().Max(),
			Amplitude: int64(t.Speed().Amplitude()),
		},
		Props:       rawProps,
		NumSensors:  int64(t.NumSensors()),
		NumRoutes:   int64(t.NumRoutes()),
		SkipOffline: t.SkipOffline(),
		CreatedAt:   t.CreatedAt().Unix(),
		UpdatedAt:   t.UpdatedAt().Unix(),
		RunningAt:   t.RunningAt().Unix(),
		StoppedAt:   t.StoppedAt().Unix(),
	}, nil
}

func navigator2model(n *types.Navigator) *Navigator {
	return &Navigator{
		Lon:             n.Lon,
		Lat:             n.Lat,
		Distance:        n.Distance,
		RouteDistance:   n.RouteDistance,
		RouteIndex:      int64(n.RouteIndex),
		TrackDistance:   n.TrackDistance,
		TrackIndex:      int64(n.TrackIndex),
		SegmentDistance: n.SegmentDistance,
		SegmentIndex:    int64(n.SegmentIndex),
		Units:           "meters",
	}
}

func toErr(e *Error) error {
	return errors.New(e.Msg)
}
