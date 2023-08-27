package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/curve"
	"github.com/mmadfox/go-gpsgen/properties"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	"github.com/valyala/fasthttp"
)

func decodeID(c *fiber.Ctx, name string) (types.ID, error) {
	uid, err := types.ParseID(c.Params(name))
	if err != nil {
		return types.ID{}, fmt.Errorf("%w: failed to decode tracker.%s: %w",
			errBadRequest, name, err)
	}
	return uid, err
}

func decodeIndex(c *fiber.Ctx, name string) (int, error) {
	index, err := strconv.Atoi(c.Params(name))
	if err != nil {
		return -1, fmt.Errorf("%w: failed to decode tracker.%s: %w",
			errBadRequest, name, err)
	}
	return index, err
}

func decodeAddTrackerRequest(r *fasthttp.Request) (*generator.NewTrackerOptions, error) {
	req := new(addTrackerRequest)
	if err := json.Unmarshal(r.Body(), &req); err != nil {
		return nil, fmt.Errorf("%w: %v", errBadRequest, err)
	}

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

	if len(req.CustomID) > 0 {
		cid, err := types.ParseCustomID(req.CustomID)
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

	elevation, err := types.ParseElevation(
		req.Elevation.Min,
		req.Elevation.Max,
		int(req.Elevation.Amplitude),
		curve.CurveMode(req.Elevation.Mode),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tracker.elevation: %w", err)
	}
	opts.Elevation = elevation

	offline, err := types.ParseOffline(
		int(req.Offline.Min),
		int(req.Offline.Max),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tracker.offline: %w", err)
	}
	opts.Offline = offline

	battery, err := types.ParseBattery(
		req.Battery.Min,
		req.Battery.Max,
		time.Duration(req.Battery.ChargeTime),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tracker.battery: %w", err)
	}
	opts.Battery = battery

	if len(req.Props) > 0 {
		opts.Props = (*properties.Properties)(&req.Props)
	}

	speed, err := types.ParseSpeed(
		req.Speed.Min,
		req.Speed.Max,
		int(req.Speed.Amplitude),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to decode tracker.speed: %w", err)
	}
	opts.Speed = speed

	return &opts, nil
}

func decodeUpdateTrackerRequest(c *fiber.Ctx) (generator.UpdateTrackerOptions, error) {
	opts := generator.UpdateTrackerOptions{}
	req := updateTrackerRequest{}
	if err := json.NewDecoder(c.Request().BodyStream()).Decode(&req); err != nil {
		return opts, fmt.Errorf("%w: %v", errBadRequest, err)
	}

	if len(req.Model) > 0 {
		model, err := types.ParseModel(req.Model)
		if err != nil {
			return opts, fmt.Errorf("%w: failed to decode tracker.model: %v",
				errBadRequest, err)
		}
		opts.Model = &model
	}

	if len(req.Color) > 0 {
		color, err := types.ParseColor(req.Color)
		if err != nil {
			return opts, fmt.Errorf("%w: failed to decode tracker.color: %v",
				errBadRequest, err)
		}
		opts.Color = &color
	}

	if len(req.UserID) > 0 {
		cid, err := types.ParseCustomID(req.UserID)
		if err != nil {
			return opts, fmt.Errorf("%w: failed to decode tracker.customID: %v",
				errBadRequest, err)
		}
		opts.UserID = &cid
	}

	if len(req.Descr) > 0 {
		descr, err := types.ParseDescription(req.Descr)
		if err != nil {
			return opts, fmt.Errorf("%w: failed to decode tracker.description: %v",
				errBadRequest, err)
		}
		opts.Descr = &descr
	}

	return opts, nil
}

func decodeAddRoutesRequest(c *fiber.Ctx) ([]*gpsgen.Route, error) {
	if len(c.Body()) < 60 {
		return nil, fmt.Errorf("%w: no data", errBadRequest)
	}
	var routes []*gpsgen.Route
	var err error
	data := c.Body()
	header := data[:60]
	if bytes.Contains(header, []byte("{")) {
		routes, err = gpsgen.DecodeGeoJSONRoutes(data)
	} else if bytes.Contains(header, []byte("<")) {
		routes, err = gpsgen.DecodeGPXRoutes(data)
	} else {
		err = fmt.Errorf("%w: invalid data format",
			errBadRequest)
	}
	if err != nil {
		return nil, fmt.Errorf("%w: %v", errBadRequest, err)
	}
	return routes, nil
}

func decodeAddSensorRequest(c *fiber.Ctx) (*types.Sensor, error) {
	req := sensor{}
	if err := json.Unmarshal(c.Body(), &req); err != nil {
		return nil, fmt.Errorf("%w: %v", errBadRequest, err)
	}
	return types.NewSensor(req.Name, req.Min, req.Max, req.Amplitude, req.Mode)
}
