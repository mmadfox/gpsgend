package device

import (
	"context"
)

type DeviceView struct {
	ID          string `json:"device_id" bson:"device_id,omitempty"`
	UserID      string `json:"user_id" bson:"user_id"`
	Model       string `json:"model" bson:"model"`
	Color       string `json:"color" bson:"color"`
	Description string `json:"description" bson:"description"`
	Speed       struct {
		Min       float64 `json:"min" bson:"min"`
		Max       float64 `json:"max" bson:"max"`
		Amplitude int     `json:"amplitude" bson:"amplitude"`
	} `json:"speed" bson:"speed"`
	Battery struct {
		Min        float64 `json:"min" bson:"min"`
		Max        float64 `json:"max" bson:"max"`
		ChargeTime string  `json:"chargeTime" bson:"charge_time"`
	} `json:"battery" bson:"battery"`
	Elevation struct {
		Min       float64 `json:"min" bson:"min"`
		Max       float64 `json:"max" bson:"max"`
		Amplitude int     `json:"amplitude" bson:"amplitude"`
	} `bson:"elevation"`
	Offline struct {
		Min int `json:"min" bson:"min"`
		Max int `json:"max" bson:"max"`
	} `json:"offline" bson:"offline"`
	Props  map[string]string `json:"properties" bson:"properties"`
	Status struct {
		ID   int    `json:"id" bson:"id"`
		Text string `json:"text" bson:"text"`
	} `json:"status" bson:"status"`
	NumRoutes  int   `json:"numRoutes" bson:"num_routes"`
	NumSensors int   `json:"numSensors" bson:"num_sensors"`
	CreatedAt  int64 `json:"createdAt" bson:"created_at"`
	UpdatedAt  int64 `json:"updatedAt" bson:"updated_at"`
	Version    int   `json:"version" bson:"version"`
}

type Query interface {
	Search(ctx context.Context, f QueryFilter) (SearchResult, error)
}

type SearchResult struct {
	Meta struct {
		TotalDevices int64 `json:"total"`
		Found        int64 `json:"found"`
		Page         int64 `json:"page"`
		Limit        int64 `json:"limit"`
	} `json:"meta"`
	Devices []DeviceView `json:"devices"`
}

type QueryFilter struct {
	Model  *string
	ID     *[]string
	Sensor *[]string
	User   *[]string
	Status *[]int
	Limit  int64
	Page   int64
}
