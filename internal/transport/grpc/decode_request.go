package grpc

import (
	"fmt"
	"time"

	gpsgen "github.com/mmadfox/go-gpsgen"
	stdtypes "github.com/mmadfox/go-gpsgen/types"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func decodeNewTrackerRequest(req *gpsgendproto.NewTrackerRequest) (generator.NewTrackerOptions, error) {
	opts := generator.NewTrackerOptions{}
	if len(req.Model) > 0 {
		model, err := types.ParseModel(req.Model)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.model", err)
		}
		opts.Model = &model
	}

	if len(req.Color) > 0 {
		color, err := types.ParseColor(req.Color)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.color", err)
		}
		opts.Color = &color
	}

	if len(req.CustomId) > 0 {
		cid, err := types.ParseCustomID(req.CustomId)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.customID", err)
		}
		opts.UserID = &cid
	}

	if len(req.Descr) > 0 {
		descr, err := types.ParseDescription(req.Descr)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.description", err)
		}
		opts.Descr = &descr
	}

	opts.SkipOffline = req.SkipOffline

	if req.Elevation != nil {
		elevation, err := types.ParseElevation(
			req.Elevation.Min,
			req.Elevation.Max,
			int(req.Elevation.Amplitude),
			stdtypes.SensorMode(req.Elevation.Mode),
		)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.elevation", err)
		}
		opts.Elevation = elevation
	}

	if req.Battery != nil {
		battery, err := types.ParseBattery(
			req.Battery.Min,
			req.Battery.Max,
			time.Duration(req.Battery.ChargeTime),
		)
		if err != nil {
			return opts, err
		}
		opts.Battery = battery
	}

	if req.Speed != nil {
		speed, err := types.ParseSpeed(
			req.Speed.Min,
			req.Speed.Max,
			int(req.Speed.Amplitude),
		)
		if err != nil {
			return opts, fmt.Errorf("%w: tracker.speed", err)
		}
		opts.Speed = speed
	}

	return opts, nil
}

func decodeSearchTrackersRequest(req *gpsgendproto.SearchTrackersRequest) (generator.Filter, error) {
	filter := generator.Filter{}
	if req.Filter == nil {
		return filter, nil
	}
	return generator.Filter{
		Model:  req.Filter.Model,
		Descr:  req.Filter.Descr,
		Color:  req.Filter.Color,
		Status: int(req.Filter.Status),
		Limit:  req.Filter.Limit,
		Offset: req.Filter.Offset,
	}, nil
}

func decodeUpdateTrackerRequest(req *gpsgendproto.UpdateTrackerRequest) (types.ID, generator.UpdateTrackerOptions, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return types.ID{}, generator.UpdateTrackerOptions{}, err
	}

	opts := generator.UpdateTrackerOptions{}

	if len(req.Model) > 0 {
		model, err := types.ParseModel(req.Model)
		if err != nil {
			return trackerID, opts, err
		}
		opts.Model = &model
	}

	if len(req.Color) > 0 {
		color, err := types.ParseColor(req.Color)
		if err != nil {
			return trackerID, opts, err
		}
		opts.Color = &color
	}

	if len(req.CustomId) > 0 {
		cid, err := types.ParseCustomID(req.CustomId)
		if err != nil {
			return trackerID, opts, err
		}
		opts.UserID = &cid
	}

	if len(req.Descr) > 0 {
		descr, err := types.ParseDescription(req.Descr)
		if err != nil {
			return trackerID, opts, err
		}
		opts.Descr = &descr
	}

	return trackerID, opts, nil
}

func decodeAddRoutesRequest(req *gpsgendproto.AddRoutesRequest) (types.ID, []*gpsgen.Route, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return trackerID, nil, err
	}
	routes, err := gpsgen.DecodeRoutes(req.Routes)
	if err != nil {
		return trackerID, nil, err
	}
	return trackerID, routes, nil
}

func decodeAddSensorRequest(req *gpsgendproto.AddSensorRequest) (types.ID, *gpsgen.Sensor, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return trackerID, nil, err
	}
	sensor, err := gpsgen.NewSensor(
		req.Name,
		req.Min,
		req.Max,
		int(req.Amplitude),
		stdtypes.SensorMode(req.Mode),
	)
	if err != nil {
		return trackerID, nil, err
	}
	return trackerID, sensor, nil
}
