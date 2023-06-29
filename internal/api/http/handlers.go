package http

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/go-gpsgen/geojson"
	"github.com/mmadfox/go-gpsgen/gpx"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/go-gpsgen/route"
	"github.com/mmadfox/gpsgend/internal/device"
)

func createDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		req, err := decodeNewDeviceRequest(c)
		if err != nil {
			return err
		}

		var newDevice *device.Device

		switch req.UsePresets {
		case false:
			newDevice, err = duc.NewDevice(
				c.Context(),
				req.ToNewDeviceParams(),
			)
		case true:
			newDevice, err = duc.NewPresetsDevice(
				c.Context(),
				req.ToNewPresetsDeviceParams(),
			)
		}
		if err != nil {
			return err
		}

		return c.JSON(toDeviceModel(newDevice))
	}
}

func findDeviceByIDHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		device, err := duc.DeviceByID(ctx.Context(), deviceID)
		if err != nil {
			return err
		}

		return ctx.JSON(toDeviceModel(device))
	}
}

func updateDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		req, err := decodeUpdateDeviceRequest(ctx)
		if err != nil {
			return err
		}

		device, err := duc.UpdateDevice(
			ctx.Context(),
			req.DeviceID,
			req.ToUpdateDeviceParams(),
		)
		if err != nil {
			return err
		}

		return ctx.JSON(toDeviceModel(device))
	}
}

func removeDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		if err := duc.RemoveDevice(ctx.Context(), deviceID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func runDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		if err := duc.RunDevice(ctx.Context(), deviceID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func stopDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		if err := duc.StopDevice(ctx.Context(), deviceID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func pauseDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		if err := duc.PauseDevice(ctx.Context(), deviceID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func resumeDeviceHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		if err := duc.ResumeDevice(ctx.Context(), deviceID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func addSensorHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		req, err := decodeAddSensorRequest(ctx)
		if err != nil {
			return err
		}

		sensors, err := duc.AddSensor(ctx.Context(), deviceID, req.ToAddSensorParams())
		if err != nil {
			return err
		}

		return writeSensorsResponse(ctx, sensors)
	}
}

func sensorsHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		sensors, err := duc.Sensors(ctx.Context(), deviceID)
		if err != nil {
			return err
		}

		return writeSensorsResponse(ctx, sensors)
	}
}

func removeSensorHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		sensorID, err := decodeSensorID(ctx)
		if err != nil {
			return err
		}

		if err := duc.RemoveSensor(ctx.Context(), deviceID, sensorID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func uploadRouteHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		fileHeader, err := ctx.FormFile(uploadFilename)
		if err != nil {
			return err
		}
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer file.Close()

		if fileHeader.Size < minUploadFileSize {
			return fmt.Errorf("file[%q] is empty", fileHeader.Filename)
		}

		if fileHeader.Size > maxUploadFileSize {
			return fmt.Errorf("file[%q] too large. maximum file size %d bytes",
				fileHeader.Filename, maxUploadFileSize)
		}

		var routes []*navigator.Route
		buf := bytes.NewBuffer(nil)
		io.CopyN(buf, file, minUploadFileSize)
		str := strings.TrimSpace(buf.String())
		io.Copy(buf, file)
		isGPX := strings.IndexByte(str, '<') >= 0
		isGeoJSON := strings.IndexByte(str, '{') >= 0

		switch {
		case isGeoJSON:
			routes, err = geojson.Decode(buf.Bytes())
		case isGPX:
			routes, err = gpx.Decode(buf.Bytes())
		default:
			err = fmt.Errorf("undefined file %q data format", fileHeader.Filename)
		}
		if err != nil {
			return err
		}

		if err := duc.AddRoutes(ctx.Context(), deviceID, routes); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func importRouteHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		body := ctx.Request().Body()
		bodySize := len(body)

		if bodySize < minUploadFileSize {
			return fmt.Errorf("routes data not found. got %d bytes, expected > %d bytes",
				bodySize, minUploadFileSize)
		}

		header := body[:minUploadFileSize]
		isURL := bytes.Contains(header, []byte("http"))
		if isURL {
			addr := string(body)
			body, err = downloadFileFromURL(addr)
			if err != nil {
				return err
			}
			bodySize = len(body)
			if bodySize < minUploadFileSize {
				return fmt.Errorf("routes data not found. got %d bytes, expected > %d bytes",
					bodySize, minUploadFileSize)
			}
			header = body[:minUploadFileSize]
		}

		if bodySize > maxUploadFileSize {
			return fmt.Errorf("routes data too large. got %d bytes, expected < %d bytes",
				bodySize, maxUploadFileSize)
		}

		var routes []*navigator.Route
		isGPX := bytes.IndexByte(header, '<') >= 0
		isGeoJSON := bytes.IndexByte(header, '{') >= 0

		switch {
		case isGeoJSON:
			routes, err = geojson.Decode(body)
		case isGPX:
			routes, err = gpx.Decode(body)
		default:
			err = fmt.Errorf("undefined data format")
		}
		if err != nil {
			return err
		}

		if err := duc.AddRoutes(ctx.Context(), deviceID, routes); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func downloadRouteHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		routes, err := duc.Routes(ctx.Context(), deviceID)
		if err != nil {
			return err
		}

		return writeDownloadRoutesResponse(ctx, deviceID, routes)
	}
}

func routesHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}

		routes, err := duc.Routes(ctx.Context(), deviceID)
		if err != nil {
			return err
		}

		return writeRoutesResponse(ctx, routes)
	}
}

func removeRouteHandler(duc device.UseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		deviceID, err := decodeDeviceID(ctx)
		if err != nil {
			return err
		}
		routeID, err := decodeRouteID(ctx)
		if err != nil {
			return err
		}

		if err := duc.RemoveRoute(ctx.Context(), deviceID, routeID); err != nil {
			return err
		}

		return writeSuccessResponse(ctx)
	}
}

func randomRouteHandler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var routes []*navigator.Route
		countryID, err := decodeCountryParam(ctx, true)
		withCountry := err == nil

		count, err := strconv.Atoi(ctx.Query("count", "1"))
		if err != nil {
			return err
		}
		if count < 1 {
			count = 1
		}

		for i := 0; i < count; i++ {
			var newRoute *navigator.Route
			var err error
			switch withCountry {
			case true:
				newRoute, err = route.GenerateFor(countryID)
			case false:
				newRoute, err = route.Generate()
			}
			if err != nil {
				return err
			}
			routes = append(routes, newRoute)
		}

		data, err := geojson.Encode(routes)
		if err != nil {
			return err
		}

		_, err = ctx.Write(data)
		return err
	}
}

func searchDevicesHandler(deviceQuery device.Query) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		qf, err := decodeDeviceQueryFilter(ctx)
		if err != nil {
			return fmt.Errorf("query filter decode error: %w", err)
		}

		result, err := deviceQuery.Search(ctx.Context(), qf)
		if err != nil {
			return err
		}

		return ctx.JSON(result)
	}
}
