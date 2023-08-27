package generator

import (
	"github.com/mmadfox/go-gpsgen/properties"
)

type TrackerSnapshot struct {
	ID       string `json:"trackerId" bson:"tracker_id,omitempty"`
	CustomID string `json:"customId" bson:"custom_id"`
	Status   struct {
		ID   int    `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"status" bson:"status"`
	Model   string `json:"model" bson:"model"`
	Color   string `json:"color" bson:"color"`
	Descr   string `json:"descr"  bson:"descr"`
	Offline struct {
		Min int `json:"min" bson:"min"`
		Max int `json:"max" bson:"max"`
	} `json:"offline" bson:"offline"`
	Elevation struct {
		Min       float64 `json:"min" bson:"min"`
		Max       float64 `json:"max" bson:"max"`
		Amplitude int     `json:"amplitude" bson:"amplitude"`
		Mode      int     `json:"mode" bson:"mode"`
	} `json:"elevation" bson:"elevation"`
	Battery struct {
		Min        float64 `json:"min" bson:"min"`
		Max        float64 `json:"max" bson:"max"`
		ChargeTime float64 `json:"chargeTime" bson:"charge_time"`
	} `json:"battery" bson:"battery"`
	Speed struct {
		Min       float64 `json:"min" bson:"min"`
		Max       float64 `json:"max" bson:"max"`
		Amplitude int     `json:"amplitude" bson:"amplitude"`
	} `json:"speed" bson:"speed"`
	Props       properties.Properties `json:"props,omitempty" bson:"props"`
	Sensors     []SensorSnapshot      `json:"sensors" bson:"sensors"`
	NumSensors  int                   `json:"numSensors" bson:"num_sensors"`
	NumRoutes   int                   `json:"numRoutes" bson:"num_routes"`
	SkipOffline bool                  `json:"skipOffline" bson:"skip_offline"`
	Snapshot    []byte                `json:"-" bson:"snapshot"`
	Routes      []byte                `json:"-" bson:"routes"`
	CreatedAt   int64                 `json:"createdAt" bson:"created_at"`
	UpdatedAt   int64                 `json:"updatedAt" bson:"updated_at"`
	RunningAt   int64                 `json:"runningAt" bson:"running_at"`
	StoppedAt   int64                 `json:"stoppedAt" bson:"stopped_at"`
	Version     int                   `json:"version" bson:"version"`
}

type SensorSnapshot struct {
	ID        string  `json:"sensorId" bson:"sensor_id"`
	Name      string  `json:"name" bson:"name"`
	Min       float64 `json:"min" bson:"min"`
	Max       float64 `json:"max" bson:"max"`
	Amplitude int     `json:"amplitude" bson:"amplitude"`
	Mode      int     `json:"mode" bson:"mode"`
}
