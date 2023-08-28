package logging

import (
	"context"
	"time"

	"log/slog"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/middleware"
	"github.com/mmadfox/gpsgend/internal/transport"
	"github.com/mmadfox/gpsgend/internal/types"
)

type logging struct {
	logger  *slog.Logger
	service generator.Service
}

func With(logger *slog.Logger) middleware.Middleware {
	return func(svc generator.Service) generator.Service {
		return logging{
			logger:  logger,
			service: svc,
		}
	}
}

func (l logging) NewTracker(ctx context.Context, opts *generator.NewTrackerOptions) (trk *generator.Tracker, err error) {
	call := "call generator.NewTracker"
	attrs := newAttrs(ctx)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.NewTracker(ctx, opts)
}

func (l logging) SearchTrackers(ctx context.Context, f generator.Filter) (res generator.SearchResult, err error) {
	call := "call generator.SearchTrackers"
	attrs := newAttrs(ctx)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "found", len(res.Trackers), slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.SearchTrackers(ctx, f)
}

func (l logging) RemoveTracker(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.RemoveTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.RemoveTracker(ctx, trackerID)
}

func (l logging) UpdateTracker(ctx context.Context, trackerID types.ID, opts generator.UpdateTrackerOptions) (err error) {
	call := "call generator.UpdateTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.UpdateTracker(ctx, trackerID, opts)
}

func (l logging) FindTracker(ctx context.Context, trackerID types.ID) (trk *generator.Tracker, err error) {
	call := "call generator.FindTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.FindTracker(ctx, trackerID)
}

func (l logging) StartTracker(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.StartTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.StartTracker(ctx, trackerID)
}

func (l logging) StopTracker(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.StopTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.StopTracker(ctx, trackerID)
}

func (l logging) TrackerState(ctx context.Context, trackerID types.ID) (state *proto.Device, err error) {
	call := "call generator.TrackerState"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.TrackerState(ctx, trackerID)
}

func (l logging) AddRoutes(ctx context.Context, trackerID types.ID, newRoutes []*gpsgen.Route) (err error) {
	call := "call generator.AddRoutes"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routes", len(newRoutes))
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.AddRoutes(ctx, trackerID, newRoutes)
}

func (l logging) RemoveRoute(ctx context.Context, trackerID types.ID, routeID types.ID) (err error) {
	call := "call generator.RemoveRoute"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeID", routeID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.RemoveRoute(ctx, trackerID, routeID)
}

func (l logging) Routes(ctx context.Context, trackerID types.ID) (routes []*gpsgen.Route, err error) {
	call := "call generator.Routes"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "found", len(routes), slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.Routes(ctx, trackerID)
}

func (l logging) RouteAt(ctx context.Context, trackerID types.ID, routeIndex int) (route *gpsgen.Route, err error) {
	call := "call generator.RouteAt"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeIndex", routeIndex)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.RouteAt(ctx, trackerID, routeIndex)
}

func (l logging) RouteByID(ctx context.Context, trackerID, routeID types.ID) (route *gpsgen.Route, err error) {
	call := "call generator.RouteByID"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeID", routeID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.RouteByID(ctx, trackerID, routeID)
}

func (l logging) ResetRoutes(ctx context.Context, trackerID types.ID) (ok bool, err error) {
	call := "call generator.ResetRoutes"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ResetRoutes(ctx, trackerID)
}

func (l logging) ResetNavigator(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.ResetNavigator"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ResetNavigator(ctx, trackerID)
}

func (l logging) ToNextRoute(ctx context.Context, trackerID types.ID) (nav types.Navigator, ok bool, err error) {
	call := "call generator.ToNextRoute"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ToNextRoute(ctx, trackerID)
}

func (l logging) ToPrevRoute(ctx context.Context, trackerID types.ID) (nav types.Navigator, ok bool, err error) {
	call := "call generator.ToPrevRoute"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ToPrevRoute(ctx, trackerID)
}

func (l logging) MoveToRoute(ctx context.Context, trackerID types.ID, routeIndex int) (nav types.Navigator, ok bool, err error) {
	call := "call generator.MoveToRoute"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeIndex", routeIndex)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.MoveToRoute(ctx, trackerID, routeIndex)
}

func (l logging) MoveToRouteByID(ctx context.Context, trackerID types.ID, routeID types.ID) (nav types.Navigator, ok bool, err error) {
	call := "call generator.MoveToRouteByID"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeID", routeID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.MoveToRouteByID(ctx, trackerID, routeID)
}

func (l logging) MoveToTrack(ctx context.Context, trackerID types.ID, routeIndex, trackIndex int) (nav types.Navigator, ok bool, err error) {
	call := "call generator.MoveToTrack"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeIndex", routeIndex)
	attrs = append(attrs, "trackIndex", trackIndex)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.MoveToTrack(ctx, trackerID, routeIndex, trackIndex)
}

func (l logging) MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID types.ID) (nav types.Navigator, ok bool, err error) {
	call := "call generator.MoveToTrackByID"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeID", routeID)
	attrs = append(attrs, "trackID", trackID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.MoveToTrackByID(ctx, trackerID, routeID, trackID)
}

func (l logging) MoveToSegment(ctx context.Context, trackerID types.ID, routeIndex, trackIndex, segmentIndex int) (nav types.Navigator, ok bool, err error) {
	call := "call generator.MoveToSegment"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "routeIndex", routeIndex)
	attrs = append(attrs, "trackIndex", trackIndex)
	attrs = append(attrs, "segmentIndex", segmentIndex)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "status", ok, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.MoveToSegment(ctx, trackerID, routeIndex, trackIndex, segmentIndex)
}

func (l logging) AddSensor(ctx context.Context, trackerID types.ID, sensor *types.Sensor) (err error) {
	call := "call generator.AddSensor"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "sensorID", sensor.ID())
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.AddSensor(ctx, trackerID, sensor)
}

func (l logging) RemoveSensor(ctx context.Context, trackerID types.ID, sensorID types.ID) (err error) {
	call := "call generator.RemoveSensor"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	attrs = append(attrs, "sensorID", sensorID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.RemoveSensor(ctx, trackerID, sensorID)
}

func (l logging) Sensors(ctx context.Context, trackerID types.ID) (sensors []*types.Sensor, err error) {
	call := "call generator.Sensors"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "found", len(sensors), slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.Sensors(ctx, trackerID)
}

func (l logging) ShutdownTracker(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.ShutdownTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ShutdownTracker(ctx, trackerID)
}

func (l logging) ResumeTracker(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.ResumeTracker"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.ResumeTracker(ctx, trackerID)
}

func (l logging) Stats(ctx context.Context) (item []generator.StatsItem, err error) {
	call := "call generator.Stats"
	attrs := newAttrs(ctx)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, "found", len(item), slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.Stats(ctx)
}

func (l logging) Sync(ctx context.Context, trackerID types.ID) (err error) {
	call := "call generator.Sync"
	attrs := newAttrs(ctx)
	attrs = append(attrs, "trackerID", trackerID)
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(call, params, slog.Duration("took", time.Since(start)), "err", err)
		} else {
			l.logger.Info(call, params, slog.Duration("took", time.Since(start)))
		}
	}(time.Now())

	return l.service.Sync(ctx, trackerID)
}

func newAttrs(ctx context.Context) []any {
	attrs := []any{}

	reqID := transport.RequestIDFromContext(ctx)
	if len(reqID) > 0 {
		attrs = append(attrs, slog.String("requestid", reqID))
	}

	return attrs
}
