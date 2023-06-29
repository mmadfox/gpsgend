package mongo

type deviceModel struct {
	ID          string `bson:"device_id,omitempty"`
	UserID      string `bson:"user_id"`
	Model       string `bson:"model"`
	Color       string `bson:"color"`
	Description string `bson:"description"`
	Speed       struct {
		Min       float64 `bson:"min"`
		Max       float64 `bson:"max"`
		Amplitude int     `bson:"amplitude"`
	} `bson:"speed"`
	Battery struct {
		Min        float64 `bson:"min"`
		Max        float64 `bson:"max"`
		ChargeTime string  `bson:"charge_time"`
	} `bson:"battery" `
	Elevation struct {
		Min       float64 `bson:"min"`
		Max       float64 `bson:"max"`
		Amplitude int     `bson:"amplitude"`
	} `bson:"elevation"`
	Offline struct {
		Min int `bson:"min"`
		Max int `bson:"max"`
	} `bson:"offline"`
	Props    map[string]string `bson:"properties"`
	Snapshot []byte            `bson:"snapshot"`
	Status   struct {
		ID   int    `bson:"id"`
		Text string `bson:"text"`
	} `bson:"status"`
	Sensors    []sensorModel `bson:"sensors"`
	Routes     []routeModel  `bson:"routes"`
	NumRoutes  int           `bson:"num_routes"`
	NumSensors int           `bson:"num_sensors"`
	CreatedAt  int64         `bson:"created_at"`
	UpdatedAt  int64         `bson:"updated_at"`
	Version    int           `bson:"version"`
}

type sensorModel struct {
	ID        string  `bson:"id"`
	Name      string  `bson:"name"`
	Min       float64 `bson:"min"`
	Max       float64 `bson:"max"`
	Amplitude int     `bson:"amplitude"`
}

type routeModel struct {
	ID    string `bson:"id"`
	Color string `bson:"color"`
	Route []byte `bson:"routes"`
}
