package http

import "time"

var zeroSensors = []sensorModel{}

type deviceModel struct {
	ID     string `json:"deviceID"`
	UserID string `json:"userID"`
	Status struct {
		ID   int    `json:"id"`
		Text string `json:"text"`
	} `json:"status"`
	Color       string `json:"color"`
	Model       string `json:"model"`
	Description string `json:"description"`
	Speed       struct {
		Max       float64 `json:"max"`
		Min       float64 `json:"min"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Battery struct {
		Max        float64       `json:"max"`
		Min        float64       `json:"min"`
		ChargeTime time.Duration `json:"chargeTime"`
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
	Props   map[string]string `json:"properties"`
	Sensors []sensorModel     `json:"sensors"`
}

type sensorModel struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Amplitude int     `json:"amplitude"`
}

type routeModel struct {
	ID     string `json:"routeId"`
	Color  string `json:"color"`
	Routes []byte `json:"b64GeoJSON"`
}
