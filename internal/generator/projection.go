package generator

import (
	"context"

	"github.com/mmadfox/go-gpsgen/properties"
)

// Query defines an interface for querying tracker data.
type Query interface {
	// SearchTrackers searches for trackers based on the given filter criteria.
	SearchTrackers(ctx context.Context, f Filter) (SearchResult, error)
}

// Filter defines filtering criteria for tracker search.
type Filter struct {
	TrackerIDs []string
	Status     int
	Term       string
	Limit      int64
	Offset     int64
}

// SearchResult holds the result of a tracker search operation.
type SearchResult struct {
	Trackers []TrackerView `json:"trackers"`
	Next     int64         `json:"next"`
}

// TrackerView provides a structured view of tracker information for display or transmission.
type TrackerView struct {
	ID       string `json:"id" bson:"tracker_id"`
	CustomID string `json:"customId" bson:"custom_id"`
	Status   struct {
		ID   int    `json:"id" bson:"id"`
		Name string `json:"name" bson:"name"`
	} `json:"status" bson:"status"`
	Model   string `json:"model" bson:"model"`
	Color   string `json:"color" bson:"color"`
	Descr   string `json:"descr" bson:"descr"`
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
	Props       properties.Properties `json:"props,omitempty" bson:"props,omitempty"`
	NumSensors  int                   `json:"numSensors" bson:"num_sensors"`
	NumRoutes   int                   `json:"numRoutes" bson:"num_routes"`
	SkipOffline bool                  `json:"skipOffline" bson:"skip_offline"`
	CreatedAt   int64                 `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt   int64                 `json:"updatedAt,omitempty" bson:"updated_at"`
	RunningAt   int64                 `json:"runningAt,omitempty" bson:"running_at"`
	StoppedAt   int64                 `json:"stoppedAt,omitempty" bson:"stopped_at"`
}
