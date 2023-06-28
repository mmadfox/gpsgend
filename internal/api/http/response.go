package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/types"
)

func writeRoutesResponse(c *fiber.Ctx, routes []*device.Route) error {
	response := struct {
		Routes []routeModel `json:"routes"`
	}{
		Routes: make([]routeModel, len(routes)),
	}

	for i := 0; i < len(routes); i++ {
		route, err := toRouteModel(routes[i])
		if err != nil {
			return err
		}
		response.Routes[i] = route
	}

	return c.JSON(response)
}

func writeSensorsResponse(ctx *fiber.Ctx, sensors []types.Sensor) error {
	response := struct {
		Sensors []sensorModel `json:"sensors"`
	}{
		Sensors: make([]sensorModel, len(sensors)),
	}

	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		response.Sensors[i] = sensorModel{
			ID:        sensor.ID().String(),
			Name:      sensor.Name(),
			Min:       sensor.Min(),
			Max:       sensor.Max(),
			Amplitude: sensor.Amplitude(),
		}
	}

	return ctx.JSON(response)
}

var successResponse = []byte(`{"ok": true}`)

var errorsIndex = map[error]int{
	// types
	types.ErrInvalidMinValue:      fiber.StatusBadRequest,
	types.ErrInvalidMaxValue:      fiber.StatusBadRequest,
	types.ErrInvalidRangeValue:    fiber.StatusBadRequest,
	types.ErrInvalidMinChargeTime: fiber.StatusBadRequest,
	types.ErrInvalidMaxChargeTime: fiber.StatusBadRequest,
	types.ErrInvalidMinAmplitude:  fiber.StatusBadRequest,
	types.ErrInvalidMaxAmplitude:  fiber.StatusBadRequest,
	types.ErrInvalidName:          fiber.StatusBadRequest,
	// http
	errDecodeRequest: fiber.StatusBadRequest,
	// device service
	device.ErrSensorLimitExceeded:  fiber.StatusUnprocessableEntity,
	device.ErrRoutesLimitExceeded:  fiber.StatusUnprocessableEntity,
	device.ErrSetConfig:            fiber.StatusUnprocessableEntity,
	device.ErrDeviceAlreadyRunning: fiber.StatusUnprocessableEntity,
	device.ErrDeviceAlreadyStopped: fiber.StatusUnprocessableEntity,
	device.ErrNoRoutes:             fiber.StatusUnprocessableEntity,
	device.ErrNoSnapshot:           fiber.StatusUnprocessableEntity,
}

func errorHandler(ctx *fiber.Ctx, err error) error {
	httpErr := fiber.ErrInternalServerError
	for k, v := range errorsIndex {
		if errors.Is(err, k) {
			httpErr = fiber.NewError(v, err.Error())
			break
		}
	}
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	ctx.Status(httpErr.Code)
	return ctx.JSON(httpErr)
}

func writeSuccessResponse(ctx *fiber.Ctx) error {
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	_, err := ctx.Write(successResponse)
	return err
}
