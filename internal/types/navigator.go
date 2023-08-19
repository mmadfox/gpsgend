package types

import (
	"github.com/mmadfox/go-gpsgen"
)

type Navigator struct {
	Lon             float64 `json:"lon"`
	Lat             float64 `json:"lat"`
	Distance        float64 `json:"distance"`
	RouteDistance   float64 `json:"routeDistance"`
	RouteIndex      int     `json:"routeIndex"`
	TrackDistance   float64 `json:"trackDistance"`
	TrackIndex      int     `json:"trackIndex"`
	SegmentDistance float64 `json:"segmentDistance"`
	SegmentIndex    int     `json:"segmentIndex"`
	Units           string  `json:"units"`
}

func NavigatorFromProc(proc *gpsgen.Device) Navigator {
	loc := proc.Location()
	return Navigator{
		Lon:             loc.Lon,
		Lat:             loc.Lat,
		Distance:        proc.Distance(),
		RouteDistance:   proc.Distance(),
		RouteIndex:      proc.RouteIndex(),
		TrackDistance:   proc.TrackDistance(),
		TrackIndex:      proc.TrackIndex(),
		SegmentDistance: proc.SegmentDistance(),
		SegmentIndex:    proc.SegmentIndex(),
		Units:           "meters",
	}
}
