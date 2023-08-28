package types

import (
	"github.com/mmadfox/go-gpsgen"
)

type Navigator struct {
	Lon             float64 `json:"lon"`
	Lat             float64 `json:"lat"`
	Distance        float64 `json:"distance"`
	RouteID         string  `json:"routeId"`
	RouteDistance   float64 `json:"routeDistance"`
	RouteIndex      int     `json:"routeIndex"`
	TrackID         string  `json:"trackId"`
	TrackDistance   float64 `json:"trackDistance"`
	TrackIndex      int     `json:"trackIndex"`
	SegmentDistance float64 `json:"segmentDistance"`
	SegmentIndex    int     `json:"segmentIndex"`
	Units           string  `json:"units"`
}

func NavigatorFromProc(proc *gpsgen.Device) Navigator {
	loc := proc.Location()
	nav := Navigator{
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
	if route := proc.CurrentRoute(); route != nil {
		nav.RouteID = route.ID()
	}
	if track := proc.CurrentTrack(); track != nil {
		nav.TrackID = track.ID()
	}
	return nav
}
