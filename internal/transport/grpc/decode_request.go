package grpc

import (
	"encoding/json"
	"fmt"
	"time"

	gpsgen "github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/properties"
	stdtypes "github.com/mmadfox/go-gpsgen/types"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

func decodeNewTrackerRequest(req *gpsgendproto.NewTrackerRequest) (*generator.NewTrackerOptions, error) {
	opts := generator.NewTrackerOptions{}
	if len(req.Model) > 0 {
		model, err := types.ParseModel(req.Model)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.model: %w", err)
		}
		opts.Model = &model
	}

	if len(req.Color) > 0 {
		color, err := types.ParseColor(req.Color)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.color: %w", err)
		}
		opts.Color = &color
	}

	if len(req.CustomId) > 0 {
		cid, err := types.ParseCustomID(req.CustomId)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.customID: %w", err)
		}
		opts.UserID = &cid
	}

	if len(req.Descr) > 0 {
		descr, err := types.ParseDescription(req.Descr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.description: %w", err)
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
			return nil, fmt.Errorf("failed to decode tracker.elevation: %w", err)
		}
		opts.Elevation = elevation
	}

	if req.Offline != nil {
		offline, err := types.ParseOffline(
			int(req.Offline.Min),
			int(req.Offline.Max),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.offline: %w", err)
		}
		opts.Offline = offline
	}

	if req.Battery != nil {
		battery, err := types.ParseBattery(
			req.Battery.Min,
			req.Battery.Max,
			time.Duration(req.Battery.ChargeTime),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.battery: %w", err)
		}
		opts.Battery = battery
	}

	if len(req.Props) > 0 {
		var props properties.Properties
		if err := json.Unmarshal(req.Props, &props); err != nil {
			return nil, fmt.Errorf("failed to decode tracker.props: %v", err)
		}
		opts.Props = &props
	}

	if req.Speed != nil {
		speed, err := types.ParseSpeed(
			req.Speed.Min,
			req.Speed.Max,
			int(req.Speed.Amplitude),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to decode tracker.speed: %w", err)
		}
		opts.Speed = speed
	}

	return &opts, nil
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

func decodeAddSensorRequest(req *gpsgendproto.AddSensorRequest) (types.ID, *types.Sensor, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return trackerID, nil, err
	}
	sensor, err := types.NewSensor(
		req.Name,
		req.Min,
		req.Max,
		int(req.Amplitude),
		int(req.Mode),
	)
	if err != nil {
		return trackerID, nil, err
	}
	return trackerID, sensor, nil
}
