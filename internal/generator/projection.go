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
	Model  string
	Descr  string
	Color  string
	Status int
	Limit  int64
	Offset int64
}

// SearchResult holds the result of a tracker search operation.
type SearchResult struct {
	Trackers []*TrackerView
}

// TrackerView provides a structured view of tracker information for display or transmission.
type TrackerView struct {
	ID       string `json:"id"`
	CustomID string `json:"customId"`
	Status   struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"status"`
	Model   string `json:"model"`
	Color   string `json:"color"`
	Descr   string `json:"descr"`
	Offline struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"offline"`
	Elevation struct {
		Min       float64 `json:"min"`
		Max       float64 `json:"max"`
		Amplitude int     `json:"amplitude"`
		Mode      int     `json:"mode"`
	} `json:"elevation"`
	Battery struct {
		Min        float64 `json:"min"`
		Max        float64 `json:"max"`
		ChargeTime float64 `json:"chargeTime"`
	} `json:"battery"`
	Speed struct {
		Min       float64 `json:"min"`
		Max       float64 `json:"max"`
		Amplitude int     `json:"amplitude"`
	} `json:"speed"`
	Props       properties.Properties `json:"props,omitempty"`
	NumSensors  int                   `json:"numSensors"`
	NumRoutes   int                   `json:"numRoutes"`
	SkipOffline bool                  `json:"skipOffline"`
	CreatedAt   int64                 `json:"createdAt"`
	UpdatedAt   int64                 `json:"updatedAt"`
	RunningAt   int64                 `json:"runningAt,omitempty"`
	StoppedAt   int64                 `json:"stoppedAt,omitempty"`
}
