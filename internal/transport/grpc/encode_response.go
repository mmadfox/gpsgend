package grpc

import (
	gpsgen "github.com/mmadfox/go-gpsgen"
	gpsgenpb "github.com/mmadfox/go-gpsgen/proto"
	"github.com/mmadfox/gpsgend/internal/generator"
	transportcodes "github.com/mmadfox/gpsgend/internal/transport/codes"
	"github.com/mmadfox/gpsgend/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func encodeNewTrackerResponse(t *generator.Tracker) (*NewTrackerResponse, error) {
	trk, err := tracker2model(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &NewTrackerResponse{
		Tracker: trk,
	}, nil
}

func encodeSearchTrackersResponse(result generator.SearchResult) (*SearchTrackersResponse, error) {
	trackers := make([]*Tracker, len(result.Trackers))
	for i := 0; i < len(result.Trackers); i++ {
		model, err := trackerView2model(result.Trackers[i])
		if err != nil {
			return nil, err
		}
		trackers[i] = model
	}
	return &SearchTrackersResponse{
		Trackers: trackers,
	}, nil
}

func encodeRemoveTrackerResponse() (*RemoveTrackerResponse, error) {
	return &RemoveTrackerResponse{}, nil
}

func encodeUpdateTrackerResponse() (*UpdateTrackerResponse, error) {
	return &UpdateTrackerResponse{}, nil
}

func encodeFindTrackerResponse(trk *generator.Tracker) (*FindTrackerResponse, error) {
	model, err := tracker2model(trk)
	if err != nil {
		return nil, err
	}
	return &FindTrackerResponse{
		Tracker: model,
	}, nil
}

func encodeStartTrackerResponse() (*StartTrackerResponse, error) {
	return &StartTrackerResponse{}, nil
}

func encodeStopTrackerResponse() (*StopTrackerResponse, error) {
	return &StopTrackerResponse{}, nil
}

func encodeTrackerStateResponse(state *gpsgenpb.Device) (*TrackerStateResponse, error) {
	data, err := proto.Marshal(state)
	if err != nil {
		return nil, err
	}
	return &TrackerStateResponse{
		State: data,
	}, nil
}

func encodeAddRoutesResponse() (*AddRoutesResponse, error) {
	return &AddRoutesResponse{}, nil
}

func encodeRemoveRoutesResponse() (*RemoveRouteResponse, error) {
	return &RemoveRouteResponse{}, nil
}

func encodeRoutesResponse(routes []*gpsgen.Route) (*RoutesResponse, error) {
	data, err := gpsgen.EncodeRoutes(routes)
	if err != nil {
		return nil, err
	}
	return &RoutesResponse{
		Routes: data,
	}, nil
}

func encodeRouteAtResponse(route *gpsgen.Route) (*RouteAtResponse, error) {
	data, err := gpsgen.EncodeRoutes([]*gpsgen.Route{route})
	if err != nil {
		return nil, err
	}
	return &RouteAtResponse{
		Route: data,
	}, nil
}

func encodeRouteIDResponse(route *gpsgen.Route) (*RouteByIDResponse, error) {
	data, err := gpsgen.EncodeRoutes([]*gpsgen.Route{route})
	if err != nil {
		return nil, err
	}
	return &RouteByIDResponse{
		Route: data,
	}, nil
}

func encodeResetRoutesResponse(ok bool) (*ResetRoutesResponse, error) {
	return &ResetRoutesResponse{
		Ok: ok,
	}, nil
}

func encodeResetNavigatorResponse() (*ResetNavigatorResponse, error) {
	return &ResetNavigatorResponse{}, nil
}

func encodeToNextRouteResponse(nav *types.Navigator, ok bool) (*ToNextRouteResponse, error) {
	return &ToNextRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeToPrevRouteResponse(nav *types.Navigator, ok bool) (*ToPrevRouteResponse, error) {
	return &ToPrevRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToRouteResponse(nav *types.Navigator, ok bool) (*MoveToRouteResponse, error) {
	return &MoveToRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToRouteByIDResponse(nav *types.Navigator, ok bool) (*MoveToRouteByIDResponse, error) {
	return &MoveToRouteByIDResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToTrackResponse(nav *types.Navigator, ok bool) (*MoveToTrackResponse, error) {
	return &MoveToTrackResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToTrackByIDResponse(nav *types.Navigator, ok bool) (*MoveToTrackByIDResponse, error) {
	return &MoveToTrackByIDResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToSegmentResponse(nav *types.Navigator, ok bool) (*MoveToSegmentResponse, error) {
	return &MoveToSegmentResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeAddSensorResponse(s *gpsgen.Sensor) (*AddSensorResponse, error) {
	return &AddSensorResponse{
		SensorId: s.ID(),
	}, nil
}

func encodeRemoveSensorResponse() (*RemoveSensorResponse, error) {
	return &RemoveSensorResponse{}, nil
}

func encodeSensorsResponse(sensors []*types.Sensor) (*SensorsResponse, error) {
	return &SensorsResponse{
		Sensors: sensors2model(sensors),
	}, nil
}

func encodeShutdownTrackerResponse() (*ShutdownTrackerResponse, error) {
	return &ShutdownTrackerResponse{}, nil
}

func encodeResumeTrackerResponse() (*ResumeTrackerResponse, error) {
	return &ResumeTrackerResponse{}, nil
}

func encodeNewTrackerErrorResponse(err error) (*NewTrackerResponse, error) {
	return &NewTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeSearchTrackersErrorResponse(err error) (*SearchTrackersResponse, error) {
	return &SearchTrackersResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveTrackerErrorResponse(err error) (*RemoveTrackerResponse, error) {
	return &RemoveTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeUpdateTrackerErrorResponse(err error) (*UpdateTrackerResponse, error) {
	return &UpdateTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeFindTrackerErrorResponse(err error) (*FindTrackerResponse, error) {
	return &FindTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeStartTrackerErrorResponse(err error) (*StartTrackerResponse, error) {
	return &StartTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeStopTrackerErrorResponse(err error) (*StopTrackerResponse, error) {
	return &StopTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeTrackerStateErrorResponse(err error) (*TrackerStateResponse, error) {
	return &TrackerStateResponse{
		Error: newError(err),
	}, nil
}

func encodeAddRoutesErrorResponse(err error) (*AddRoutesResponse, error) {
	return &AddRoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveRouteErrorResponse(err error) (*RemoveRouteResponse, error) {
	return &RemoveRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeRoutesErrorResponse(err error) (*RoutesResponse, error) {
	return &RoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeRouteByIDErrorResponse(err error) (*RouteByIDResponse, error) {
	return &RouteByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeRouteAtErrorResponse(err error) (*RouteAtResponse, error) {
	return &RouteAtResponse{
		Error: newError(err),
	}, nil
}

func encodeResetRoutesErrorResponse(err error) (*ResetRoutesResponse, error) {
	return &ResetRoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeResetNavigatorErrorResponse(err error) (*ResetNavigatorResponse, error) {
	return &ResetNavigatorResponse{
		Error: newError(err),
	}, nil
}

func encodeToNextRouteErrorResponse(err error) (*ToNextRouteResponse, error) {
	return &ToNextRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeToPrevRouteErrorResponse(err error) (*ToPrevRouteResponse, error) {
	return &ToPrevRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToRouteErrorResponse(err error) (*MoveToRouteResponse, error) {
	return &MoveToRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToRouteByIDErrorResponse(err error) (*MoveToRouteByIDResponse, error) {
	return &MoveToRouteByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToTrackErrorResponse(err error) (*MoveToTrackResponse, error) {
	return &MoveToTrackResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToTrackByIDErrorResponse(err error) (*MoveToTrackByIDResponse, error) {
	return &MoveToTrackByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToSegmentErrorResponse(err error) (*MoveToSegmentResponse, error) {
	return &MoveToSegmentResponse{
		Error: newError(err),
	}, nil
}

func encodeAddSensorErrorResponse(err error) (*AddSensorResponse, error) {
	return &AddSensorResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveSensorErrorResponse(err error) (*RemoveSensorResponse, error) {
	return &RemoveSensorResponse{
		Error: newError(err),
	}, nil
}

func encodeSensorsErrorResponse(err error) (*SensorsResponse, error) {
	return &SensorsResponse{
		Error: newError(err),
	}, nil
}

func encodeShutdownTrackerErrorResponse(err error) (*ShutdownTrackerResponse, error) {
	return &ShutdownTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeResumeTrackerErrorResponse(err error) (*ResumeTrackerResponse, error) {
	return &ResumeTrackerResponse{
		Error: newError(err),
	}, nil
}

func newError(err error) *Error {
	return &Error{
		Code: int64(transportcodes.FromError(err)),
		Msg:  err.Error(),
	}
}
