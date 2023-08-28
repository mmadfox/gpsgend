package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/go-gpsgen"
	gpsgenproto "github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
	"google.golang.org/protobuf/encoding/protojson"
)

type successResponse struct {
	Ok bool `json:"ok"`
}

type navigatorResponse struct {
	Navigator types.Navigator `json:"navigator"`
	Ok        bool            `json:"ok"`
}

type sensorResponse struct {
	Sensors []sensor `json:"sensors"`
}

func encodeTrackerResponse(c *fiber.Ctx, t *generator.Tracker) error {
	model := new(generator.TrackerView)
	tracker2model(t, model)
	return c.JSON(model)
}

func encodeTrackerStateResponse(c *fiber.Ctx, state *gpsgenproto.Device) error {
	d, err := protojson.Marshal(state)
	if err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	_, err = c.Write(d)
	return err
}

func encodeSuccessResponse(c *fiber.Ctx) error {
	return c.JSON(successResponse{Ok: true})
}

func encodeFlagResponse(c *fiber.Ctx, flag bool) error {
	return c.JSON(successResponse{Ok: flag})
}

func encodeRoutes(c *fiber.Ctx, routes []*gpsgen.Route) error {
	data, err := gpsgen.EncodeGeoJSONRoutes(routes)
	if err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	_, err = c.Write(data)
	return err
}

func encodeRoute(c *fiber.Ctx, route *gpsgen.Route) error {
	data, err := gpsgen.EncodeGeoJSONRoutes([]*gpsgen.Route{route})
	if err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
	_, err = c.Write(data)
	return err
}

func encodeSensorsResponse(c *fiber.Ctx, s []*types.Sensor) error {
	resp := sensorResponse{
		Sensors: make([]sensor, len(s)),
	}
	for i := 0; i < len(s); i++ {
		model := sensor{}
		sensor2model(s[i], &model)
		resp.Sensors[i] = model
	}
	return c.JSON(resp)
}
