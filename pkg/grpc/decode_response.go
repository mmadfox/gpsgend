package grpc

import (
	"errors"

	"github.com/mmadfox/go-gpsgen"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
)

func decodeNewTrackerResponse(resp *gpsgendproto.NewTrackerResponse) (*Tracker, error) {
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}
	return DecodeTracker(resp.Tracker)
}

func decodeSearchTrackersResponse(resp *gpsgendproto.SearchTrackersResponse) (SearchResult, error) {
	if resp.Error != nil {
		return SearchResult{}, decodeError(resp.Error)
	}
	sr := SearchResult{
		Trackers: make([]*Tracker, len(resp.Trackers)),
		Next:     resp.Next,
	}
	for i := 0; i < len(resp.Trackers); i++ {
		trk, err := DecodeTracker(resp.Trackers[i])
		if err != nil {
			return SearchResult{}, err
		}
		sr.Trackers[i] = trk
	}
	return sr, nil
}

func decodeRoutesResponse(resp *gpsgendproto.RoutesResponse) ([]*gpsgen.Route, error) {
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}
	routes, err := gpsgen.DecodeRoutes(resp.Routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func decodeRouteAtResponse(resp *gpsgendproto.RouteAtResponse) (*gpsgen.Route, error) {
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}
	routes, err := gpsgen.DecodeRoutes(resp.Route)
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return nil, ErrNoRoutes
	}
	return routes[0], nil
}

func decodeRouteByIDResponse(resp *gpsgendproto.RouteByIDResponse) (*gpsgen.Route, error) {
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}
	routes, err := gpsgen.DecodeRoutes(resp.Route)
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return nil, ErrNoRoutes
	}
	return routes[0], nil
}

func decodeToNextRouteResponse(resp *gpsgendproto.ToNextRouteResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeToPrevRouteResponse(resp *gpsgendproto.ToPrevRouteResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeMoveToRouteResponse(resp *gpsgendproto.MoveToRouteResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeMoveToRouteByIDResponse(resp *gpsgendproto.MoveToRouteByIDResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeMoveToTrackResponse(resp *gpsgendproto.MoveToTrackResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeMoveToTrackByIDResponse(resp *gpsgendproto.MoveToTrackByIDResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeMoveToSegmentResponse(resp *gpsgendproto.MoveToSegmentResponse) (*Navigator, bool, error) {
	if resp.Error != nil {
		return nil, false, decodeError(resp.Error)
	}
	navigator := DecodeNavigator(resp.Navigator)
	return navigator, resp.Ok, nil
}

func decodeSensorsResponse(resp *gpsgendproto.SensorsResponse) ([]*Sensor, error) {
	sensors := make([]*Sensor, len(resp.Sensors))
	for i := 0; i < len(resp.Sensors); i++ {
		sensors[i] = DecodeSensor(resp.Sensors[i])
	}
	return sensors, nil
}

func decodeError(e *gpsgendproto.Error) error {
	return errors.New(e.Msg)
}
