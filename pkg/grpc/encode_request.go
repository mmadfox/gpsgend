package grpc

import (
	"encoding/json"

	"github.com/mmadfox/go-gpsgen"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
)

func encodeNewTrackerRequest(opts *AddTrackerOptions) (*gpsgendproto.NewTrackerRequest, error) {
	var rawProps []byte
	if len(opts.Props) > 0 {
		data, err := json.Marshal(opts.Props)
		if err != nil {
			return nil, err
		}
		rawProps = data
	}
	return &gpsgendproto.NewTrackerRequest{
		Model:       opts.Model,
		Color:       opts.Color,
		CustomId:    opts.CustomID,
		Descr:       opts.Descr,
		Props:       rawProps,
		SkipOffline: opts.SkipOffline,
		Offline: &gpsgendproto.Offline{
			Min: int64(opts.Offline.Min),
			Max: int64(opts.Offline.Max),
		},
		Elevation: &gpsgendproto.Elevation{
			Min:       opts.Elevation.Min,
			Max:       opts.Elevation.Max,
			Amplitude: int64(opts.Elevation.Amplitude),
			Mode:      int64(opts.Elevation.Mode),
		},
		Battery: &gpsgendproto.Battery{
			Min:        opts.Battery.Min,
			Max:        opts.Battery.Max,
			ChargeTime: int64(opts.Battery.ChargeTime),
		},
		Speed: &gpsgendproto.Speed{
			Min:       opts.Speed.Min,
			Max:       opts.Speed.Max,
			Amplitude: int64(opts.Speed.Amplitude),
		},
	}, nil
}

func encodeSearchTrackersRequest(f *Filter) *gpsgendproto.SearchTrackersRequest {
	return &gpsgendproto.SearchTrackersRequest{
		Filter: &gpsgendproto.Filter{
			TrackerId: f.TrackerIDs,
			Term:      f.Term,
			Status:    int64(f.Status),
			Limit:     f.Limit,
			Offset:    f.Offset,
		},
	}
}

func encodeUpdateTrackerRequest(trackerID string, opts *UpdateTrackerOptions) *gpsgendproto.UpdateTrackerRequest {
	return &gpsgendproto.UpdateTrackerRequest{
		TrackerId: trackerID,
		Model:     opts.Model,
		Color:     opts.Color,
		CustomId:  opts.CustomID,
		Descr:     opts.Descr,
	}
}

func encodeAddRouteRequest(trackerID string, newRoutes []*gpsgen.Route) (*gpsgendproto.AddRoutesRequest, error) {
	data, err := gpsgen.EncodeRoutes(newRoutes)
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.AddRoutesRequest{
		TrackerId: trackerID,
		Routes:    data,
	}, nil
}
