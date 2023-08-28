package grpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen/properties"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"google.golang.org/grpc/status"
)

func DecodeSensor(m *gpsgendproto.Sensor) *Sensor {
	return &Sensor{
		ID:        m.Id,
		Name:      m.Name,
		Min:       m.Min,
		Max:       m.Max,
		Amplitude: int(m.Amplitude),
		Mode:      int(m.Mode),
	}
}

func DecodeNavigator(m *gpsgendproto.Navigator) *Navigator {
	if m == nil {
		return new(Navigator)
	}

	return &Navigator{
		Lon:             m.Lon,
		Lat:             m.Lat,
		Distance:        m.Distance,
		RouteiD:         m.RouteId,
		RouteDistance:   m.RouteDistance,
		RouteIndex:      int(m.RouteIndex),
		TrackID:         m.TrackId,
		TrackDistance:   m.TrackDistance,
		TrackIndex:      int(m.TrackIndex),
		SegmentDistance: m.SegmentDistance,
		SegmentIndex:    int(m.SegmentIndex),
	}
}

func DecodeStats(m []*gpsgendproto.StatsItem) []StatsItem {
	items := make([]StatsItem, len(m))
	for i := 0; i < len(m); i++ {
		items[i] = StatsItem{
			Status: m[i].Status,
			Total:  int(m[i].Total),
		}
	}
	return items
}

func DecodeTracker(m *gpsgendproto.Tracker) (*Tracker, error) {
	if m == nil {
		return new(Tracker), nil
	}

	var props properties.Properties
	if len(m.Props) > 0 {
		if err := json.Unmarshal(m.Props, &props); err != nil {
			return nil, err
		}
	}

	trk := Tracker{}
	trk.ID = m.Id
	trk.Model = m.Model
	trk.CustomID = m.CustomId
	trk.Color = m.Color
	trk.Descr = m.Descr

	if m.Status != nil {
		trk.Status.ID = int(m.Status.Id)
		trk.Status.Name = m.Status.Name
	}

	if m.Offline != nil {
		trk.Offline.Min = int(m.Offline.Min)
		trk.Offline.Max = int(m.Offline.Max)
	}

	if m.Elevation != nil {
		trk.Elevation.Min = m.Elevation.Min
		trk.Elevation.Max = m.Elevation.Max
		trk.Elevation.Amplitude = int(m.Elevation.Amplitude)
		trk.Elevation.Mode = int(m.Elevation.Mode)
	}

	if m.Battery != nil {
		trk.Battery.Min = m.Battery.Min
		trk.Battery.Max = m.Battery.Max
		trk.Battery.ChargeTime = time.Duration(m.Battery.ChargeTime)
	}

	if m.Speed != nil {
		trk.Speed.Min = m.Speed.Min
		trk.Speed.Max = m.Speed.Max
		trk.Speed.Amplitude = int(m.Speed.Amplitude)
	}

	trk.Props = props
	trk.NumSensors = int(m.NumSensors)
	trk.NumRoutes = int(m.NumRoutes)
	trk.SkipOffline = m.SkipOffline
	trk.CreatedAt = time.Unix(m.CreatedAt, 0)
	trk.RunningAt = time.Unix(m.RunningAt, 0)
	trk.StoppedAt = time.Unix(m.StoppedAt, 0)
	return &trk, nil
}

func toError(err error) error {
	status, ok := status.FromError(err)
	if !ok {
		return err
	}
	return errors.New(status.Message())
}

func validateID(id string, msg string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("%w %s", err, msg)
	}
	return nil
}
