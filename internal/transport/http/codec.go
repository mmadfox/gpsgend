package http

import (
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func tracker2model(trk *generator.Tracker, model *generator.TrackerView) {
	model.ID = trk.ID().String()
	model.CustomID = trk.UserID().String()
	model.Status.ID = int(trk.Status())
	model.Status.Name = trk.Status().String()
	model.Model = trk.Model().String()
	model.Color = trk.Color().String()
	model.Descr = trk.Description().String()
	model.Offline.Min = trk.Offline().Min()
	model.Offline.Max = trk.Offline().Max()
	model.Elevation.Min = trk.Elevation().Min()
	model.Elevation.Max = trk.Elevation().Max()
	model.Elevation.Amplitude = trk.Elevation().Amplitude()
	model.Elevation.Mode = int(trk.Elevation().Mode())
	model.Battery.Min = trk.Battery().Min()
	model.Battery.Max = trk.Battery().Max()
	model.Battery.ChargeTime = trk.Battery().ChargeTime().Seconds()
	model.Speed.Min = trk.Speed().Min()
	model.Speed.Max = trk.Speed().Max()
	model.Speed.Amplitude = trk.Speed().Amplitude()
	model.Props = trk.Properties()
	model.NumRoutes = trk.NumRoutes()
	model.NumSensors = trk.NumSensors()
	model.SkipOffline = trk.SkipOffline()
	model.CreatedAt = trk.CreatedAt().Unix()
	model.UpdatedAt = trk.UpdatedAt().Unix()

	if !trk.RunningAt().IsZero() {
		model.RunningAt = trk.RunningAt().Unix()
	}
	if !trk.StoppedAt().IsZero() {
		model.StoppedAt = trk.StoppedAt().Unix()
	}
}

func sensor2model(s *types.Sensor, model *sensor) {
	model.ID = s.ID().String()
	model.Name = s.Name()
	model.Max = s.Max()
	model.Min = s.Min()
	model.Amplitude = s.Amplitude()
	model.Mode = int(s.Mode())
}
