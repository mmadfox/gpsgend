package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen/route"
	"github.com/mmadfox/gpsgend/internal/device"
)

const (
	deviceParamName     = "deviceID"
	routeParamName      = "routeID"
	userParamName       = "userID"
	sensorParamName     = "sensorID"
	processParamName    = "processID"
	statusParamName     = "statusID"
	modelParamName      = "model"
	colorParamName      = "color"
	speedFromParamName  = "speed:from"
	speedTillParamName  = "speed:till"
	speedParamName      = "speed"
	uploadFilename      = "route"
	countryParamName    = "country"
	limitParamName      = "limit"
	pageParamName       = "page"
	sortFieldParamName  = "sort"
	sortVectorParamName = "sort:vec"
)

var errDecodeRequest = errors.New("invalid request")

type newDeviceRequest struct {
	UsePresets  bool              `json:"usePresets"`
	Model       string            `json:"model"`
	UserID      string            `json:"userId,omitempty"`
	Properties  map[string]string `json:"properties,omitempty"`
	Description string            `json:"description,omitempty"`
	Speed       struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Battery struct {
		Max        float64 `json:"max"`
		Min        float64 `json:"min"`
		ChargeTime int64   `json:"chargeTime"`
	} `json:"battery"`
	Elevation struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"elevation"`
	Offline struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
	WithRandomRoutes bool `json:"withRandomRoutes"`
}

func (r *newDeviceRequest) ToNewDeviceParams() device.NewDeviceParams {
	params := device.NewDeviceParams{}
	params.Model = r.Model
	params.UserID = r.UserID
	params.Properties = r.Properties
	params.Description = r.Description
	params.Speed = r.Speed
	params.Battery = r.Battery
	params.Elevation = r.Elevation
	params.Offline = r.Offline
	params.WithRandomRoutes = r.WithRandomRoutes
	return params
}

func (r *newDeviceRequest) ToNewPresetsDeviceParams() device.NewPresetsDeviceParams {
	params := device.NewPresetsDeviceParams{}
	params.Description = r.Description
	params.UserID = r.UserID
	params.Properties = r.Properties
	params.WithRandomRoutes = r.WithRandomRoutes
	return params
}

type updateDeviceRequest struct {
	DeviceID    uuid.UUID          `json:"-"`
	Model       *string            `json:"model"`
	UserID      *string            `json:"userId,omitempty"`
	Properties  *map[string]string `json:"properties,omitempty"`
	Description *string            `json:"description,omitempty"`
	Speed       *struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Battery *struct {
		Max        float64 `json:"max"`
		Min        float64 `json:"min"`
		ChargeTime int64   `json:"chargeTime"`
	} `json:"battery"`
	Elevation *struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"elevation"`
	Offline *struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
}

type addSensorRequest struct {
	Name      string  `json:"name"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Amplitude int     `json:"amplitude"`
}

func (r *addSensorRequest) ToAddSensorParams() device.AddSensorParams {
	return device.AddSensorParams{
		Name:      r.Name,
		Min:       r.Min,
		Max:       r.Max,
		Amplitude: r.Amplitude,
	}
}

func (r *updateDeviceRequest) ToUpdateDeviceParams() device.UpdateDeviceParams {
	params := device.UpdateDeviceParams{}
	params.Model = r.Model
	params.UserID = r.UserID
	params.Description = r.Description
	params.Properties = r.Properties
	params.Speed = r.Speed
	params.Battery = r.Battery
	params.Elevation = r.Elevation
	params.Offline = r.Offline
	return params
}

func decodeDeviceQueryFilter(c *fiber.Ctx) (qf device.QueryFilter, err error) {
	qf.ID = stringQueryParam(c, deviceParamName)
	qf.ProcID = stringQueryParam(c, processParamName)
	qf.SensorID = stringQueryParam(c, sensorParamName)
	qf.Color = stringQueryParam(c, colorParamName)
	qf.Model = stringQueryParam(c, modelParamName)
	qf.Status = intQueryParam(c, statusParamName)
	return
}

func decodeDeviceID(c *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params(deviceParamName))
	if err != nil {
		return uuid.Nil, err
	}
	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("invalid device id")
	}
	return id, nil
}

func decodeRouteID(c *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params(routeParamName))
	if err != nil {
		return uuid.Nil, err
	}
	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("invalid route id")
	}
	return id, nil
}

var countriesMap = map[string]route.Country{
	"ru": route.Russia,
	"sa": route.SouthArabia,
	"ao": route.Angola,
	"tr": route.Turkey,
	"fr": route.France,
	"es": route.Spain,
	"cn": route.China,
}

func decodeSensorID(c *fiber.Ctx) (uuid.UUID, error) {
	id, err := uuid.Parse(c.Params(sensorParamName))
	if err != nil {
		return uuid.Nil, err
	}
	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("invalid sensor id")
	}
	return id, nil
}

func decodeCountryParam(c *fiber.Ctx, isQuery bool) (route.Country, error) {
	var param string
	switch isQuery {
	case true:
		param = c.Query(countryParamName)
	case false:
		param = c.Params(countryParamName)
	}
	country, ok := countriesMap[strings.ToLower(param)]
	if !ok {
		return -1, fmt.Errorf("country %q not supported", param)
	}
	return country, nil
}

func decodeAddSensorRequest(c *fiber.Ctx) (*addSensorRequest, error) {
	req := new(addSensorRequest)
	return req, json.NewDecoder(
		bytes.NewReader(c.Request().Body()),
	).Decode(req)
}

func decodeNewDeviceRequest(c *fiber.Ctx) (*newDeviceRequest, error) {
	req := new(newDeviceRequest)
	if err := json.NewDecoder(
		bytes.NewReader(c.Request().Body()),
	).Decode(req); err != nil {
		return nil, errDecodeRequest
	}
	return req, nil
}

func decodeUpdateDeviceRequest(c *fiber.Ctx) (*updateDeviceRequest, error) {
	req := new(updateDeviceRequest)
	deviceID, err := decodeDeviceID(c)
	if err != nil {
		return nil, err
	}
	req.DeviceID = deviceID
	return req, json.NewDecoder(
		bytes.NewReader(c.Request().Body()),
	).Decode(req)
}

func stringQueryParam(c *fiber.Ctx, name string) *[]string {
	param := c.Query(name)
	param = strings.TrimSpace(param)
	if len(param) == 0 {
		return nil
	}
	result := strings.Split(param, ",")
	return &result
}

func intQueryParam(c *fiber.Ctx, name string) *[]int {
	param := c.Query(name)
	param = strings.TrimSpace(param)
	if len(param) == 0 {
		return nil
	}
	list := strings.Split(param, ",")
	result := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		p, err := strconv.Atoi(list[i])
		if err != nil {
			continue
		}
		result[i] = p
	}
	return &result
}
