package grpc

import (
	"encoding/json"

	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func sensors2model(sensors []*types.Sensor) []*gpsgendproto.Sensor {
	sensorsModel := make([]*gpsgendproto.Sensor, len(sensors))
	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		sensorsModel[i] = &gpsgendproto.Sensor{
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

func trackerView2model(v *generator.TrackerView) (*gpsgendproto.Tracker, error) {
	var rawProps []byte
	if len(v.Props) > 0 {
		props, err := json.Marshal(v.Props)
		if err != nil {
			return nil, err
		}
		rawProps = props
	}
	return &gpsgendproto.Tracker{
		Id:       v.ID,
		CustomId: v.CustomID,
		Status: &gpsgendproto.Status{
			Id:   int64(v.Status.ID),
			Name: v.Status.Name,
		},
		Model: v.Model,
		Color: v.Color,
		Descr: v.Descr,
		Offline: &gpsgendproto.Offline{
			Min: int64(v.Offline.Min),
			Max: int64(v.Offline.Max),
		},
		Elevation: &gpsgendproto.Elevation{
			Min:       v.Elevation.Min,
			Max:       v.Elevation.Max,
			Amplitude: int64(v.Elevation.Amplitude),
			Mode:      int64(v.Elevation.Mode),
		},
		Battery: &gpsgendproto.Battery{
			Min:        v.Battery.Min,
			Max:        v.Battery.Max,
			ChargeTime: int64(v.Battery.ChargeTime),
		},
		Speed: &gpsgendproto.Speed{
			Min:       v.Speed.Min,
			Max:       v.Speed.Max,
			Amplitude: int64(v.Speed.Amplitude),
		},
		Props:       rawProps,
		NumSensors:  int64(v.NumSensors),
		NumRoutes:   int64(v.NumRoutes),
		SkipOffline: v.SkipOffline,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		RunningAt:   v.RunningAt,
		StoppedAt:   v.StoppedAt,
	}, nil
}

func tracker2model(t *generator.Tracker) (*gpsgendproto.Tracker, error) {
	var rawProps []byte
	if len(t.Properties()) > 0 {
		props, err := json.Marshal(t.Properties())
		if err != nil {
			return nil, err
		}
		rawProps = props
	}
	return &gpsgendproto.Tracker{
		Id:       t.ID().String(),
		CustomId: t.UserID().String(),
		Status: &gpsgendproto.Status{
			Id:   int64(t.Status()),
			Name: t.Status().String(),
		},
		Model: t.Model().String(),
		Color: t.Color().String(),
		Descr: t.Description().String(),
		Offline: &gpsgendproto.Offline{
			Min: int64(t.Offline().Min()),
			Max: int64(t.Offline().Max()),
		},
		Elevation: &gpsgendproto.Elevation{
			Min:       t.Elevation().Min(),
			Max:       t.Elevation().Max(),
			Amplitude: int64(t.Elevation().Amplitude()),
			Mode:      int64(t.Elevation().Mode()),
		},
		Battery: &gpsgendproto.Battery{
			Min:        t.Battery().Min(),
			Max:        t.Battery().Max(),
			ChargeTime: int64(t.Battery().ChargeTime()),
		},
		Speed: &gpsgendproto.Speed{
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

func navigator2model(n *types.Navigator) *gpsgendproto.Navigator {
	return &gpsgendproto.Navigator{
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
