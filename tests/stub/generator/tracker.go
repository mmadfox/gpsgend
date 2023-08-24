package stub_generator

import (
	"time"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func Tracker() *generator.Tracker {
	builder := generator.NewTrackerBuilder()
	builder.ID(types.NewID())

	model, err := types.ParseModel("Model")
	if err != nil {
		panic(err)
	}
	builder.Model(model)

	color, err := types.ParseColor("#ff0000")
	if err != nil {
		panic(err)
	}
	builder.Color(color)

	descr, err := types.ParseDescription("some descr")
	if err != nil {
		panic(err)
	}
	builder.Description(descr)

	offline, err := types.ParseOffline(1, 60)
	if err != nil {
		panic(err)
	}
	builder.Offline(offline)

	elevation, err := types.ParseElevation(0, 100, 8, 0)
	if err != nil {
		panic(err)
	}
	builder.Elevation(elevation)

	speed, err := types.ParseSpeed(1, 5, 16)
	if err != nil {
		panic(err)
	}
	builder.Speed(speed)

	battery, err := types.ParseBattery(0, 30, time.Hour)
	if err != nil {
		panic(err)
	}
	builder.Battery(battery)

	builder.Status(types.Running)

	trk, err := builder.Build()
	if err != nil {
		panic(err)
	}

	s1, err := types.NewSensor("s1", 0, 1, 4, 0)
	if err != nil {
		panic(err)
	}
	trk.AddSensor(s1)

	trk.AddRoute(gpsgen.RandomRouteForMoscow())
	trk.AddRoute(gpsgen.RandomRouteForNewYork())

	return trk
}
