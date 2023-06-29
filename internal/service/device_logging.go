package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/types"
	"golang.org/x/exp/slog"
)

type logging struct {
	logger *slog.Logger
	next   device.UseCase
}

func WithLogging(logger *slog.Logger) device.Middleware {
	return func(s device.UseCase) device.UseCase {
		return logging{
			logger: logger,
			next:   s,
		}
	}
}

func (l logging) NewPresetsDevice(ctx context.Context, p device.NewPresetsDeviceParams) (dev *device.Device, err error) {
	method := "call service.NewPresetsDevice"
	attrs := []any{
		slog.Any("NewPresetsDeviceParams", p),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Group("returns", slog.String("deviceID", dev.ID().String())),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.NewPresetsDevice(ctx, p)
}

func (l logging) NewDevice(ctx context.Context, p device.NewDeviceParams) (dev *device.Device, err error) {
	method := "call service.NewDevice"
	attrs := []any{
		slog.Any("NewDeviceParams", p),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Group("returns", slog.String("deviceID", dev.ID().String())),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.NewDevice(ctx, p)
}

func (l logging) UpdateDevice(ctx context.Context, deviceID uuid.UUID, p device.UpdateDeviceParams) (dev *device.Device, err error) {
	method := "call service.UpdateDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
		slog.Any("UpdateDeviceParams", p),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.UpdateDevice(ctx, deviceID, p)
}

func (l logging) RemoveDevice(ctx context.Context, deviceID uuid.UUID) (err error) {
	method := "call service.RemoveDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.RemoveDevice(ctx, deviceID)
}

func (l logging) RunDevice(ctx context.Context, deviceID uuid.UUID) (err error) {
	method := "call service.RunDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.RunDevice(ctx, deviceID)
}

func (l logging) StopDevice(ctx context.Context, deviceID uuid.UUID) (err error) {
	method := "call service.StopDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.StopDevice(ctx, deviceID)
}

func (l logging) PauseDevice(ctx context.Context, deviceID uuid.UUID) (err error) {
	method := "call service.PauseDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.PauseDevice(ctx, deviceID)
}

func (l logging) ResumeDevice(ctx context.Context, deviceID uuid.UUID) (err error) {
	method := "call service.ResumeDevice"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.ResumeDevice(ctx, deviceID)
}

func (l logging) AddRoutes(ctx context.Context, deviceID uuid.UUID, routes []*navigator.Route) (err error) {
	method := "call service.AddRoutes"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
		slog.Int("routes", len(routes)),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.AddRoutes(ctx, deviceID, routes)
}

func (l logging) Routes(ctx context.Context, deviceID uuid.UUID) (routes []*device.Route, err error) {
	method := "call service.Routes"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Int("returns.routes", len(routes)),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.Routes(ctx, deviceID)
}

func (l logging) RemoveRoute(ctx context.Context, deviceID uuid.UUID, routeID uuid.UUID) (err error) {
	method := "call service.RemoveRoute"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
		slog.String("routeID", routeID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.RemoveRoute(ctx, deviceID, routeID)
}

func (l logging) Sensors(ctx context.Context, deviceID uuid.UUID) (sensors []types.Sensor, err error) {
	method := "call service.Sensors"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Group("returns", slog.Int("sensors", len(sensors))),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.Sensors(ctx, deviceID)
}

func (l logging) AddSensor(ctx context.Context, deviceID uuid.UUID, p device.AddSensorParams) (sensors []types.Sensor, err error) {
	method := "call service.AddSensor"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
		slog.Any("AddSensorsParams", p),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Group("returns", slog.Int("sensors", len(sensors))),
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.AddSensor(ctx, deviceID, p)
}

func (l logging) RemoveSensor(ctx context.Context, deviceID uuid.UUID, sensorID uuid.UUID) (err error) {
	method := "call service.RemoveSensor"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
		slog.String("sensorID", sensorID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.RemoveSensor(ctx, deviceID, sensorID)
}

func (l logging) DeviceByID(ctx context.Context, deviceID uuid.UUID) (dev *device.Device, err error) {
	method := "call service.DeviceByID"
	attrs := []any{
		slog.String("deviceID", deviceID.String()),
	}

	reqID, ok := ctx.Value("requestid").(string)
	if ok {
		attrs = append(attrs, slog.String("requestID", reqID))
	}
	params := slog.Group("params", attrs...)

	defer func(start time.Time) {
		if err != nil {
			l.logger.Error(method,
				params,
				slog.Duration("took", time.Since(start)),
				"err",
				err,
			)
			return
		}
		l.logger.Info(method,
			params,
			slog.Duration("took", time.Since(start)),
		)
	}(time.Now())

	return l.next.DeviceByID(ctx, deviceID)
}
