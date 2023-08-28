package grpc

import (
	"time"

	"github.com/mmadfox/go-gpsgen/properties"
)

type Tracker struct {
	ID       string
	CustomID string
	Status   struct {
		ID   int
		Name string
	}
	Model   string
	Color   string
	Descr   string
	Offline struct {
		Min int
		Max int
	}
	Elevation struct {
		Min       float64
		Max       float64
		Amplitude int
		Mode      int
	}
	Battery struct {
		Min        float64
		Max        float64
		ChargeTime time.Duration
	}
	Speed struct {
		Min       float64
		Max       float64
		Amplitude int
	}
	Props       properties.Properties
	NumSensors  int
	NumRoutes   int
	SkipOffline bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	RunningAt   time.Time
	StoppedAt   time.Time
}

type Sensor struct {
	ID        string
	Name      string
	Min, Max  float64
	Amplitude int
	Mode      int
}

type Filter struct {
	TrackerIDs []string
	Term       string
	Status     int
	Limit      int64
	Offset     int64
}

type Navigator struct {
	Lon             float64
	Lat             float64
	Distance        float64
	RouteiD         string
	RouteDistance   float64
	RouteIndex      int
	TrackID         string
	TrackDistance   float64
	TrackIndex      int
	SegmentDistance float64
	SegmentIndex    int
	Units           string
}

type StatsItem struct {
	Status string
	Total  int
}

type SearchResult struct {
	Trackers []*Tracker
	Next     int64
}
