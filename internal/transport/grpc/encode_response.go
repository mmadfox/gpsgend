package grpc

import (
	gpsgen "github.com/mmadfox/go-gpsgen"
	gpsgenpb "github.com/mmadfox/go-gpsgen/proto"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/generator"
	transportcodes "github.com/mmadfox/gpsgend/internal/transport/codes"
	"github.com/mmadfox/gpsgend/internal/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func encodeNewTrackerResponse(t *generator.Tracker) (*gpsgendproto.NewTrackerResponse, error) {
	trk, err := tracker2model(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &gpsgendproto.NewTrackerResponse{
		Tracker: trk,
	}, nil
}

func encodeSearchTrackersResponse(result generator.SearchResult) (*gpsgendproto.SearchTrackersResponse, error) {
	trackers := make([]*gpsgendproto.Tracker, len(result.Trackers))
	for i := 0; i < len(result.Trackers); i++ {
		model, err := trackerView2model(result.Trackers[i])
		if err != nil {
			return nil, err
		}
		trackers[i] = model
	}
	return &gpsgendproto.SearchTrackersResponse{
		Trackers: trackers,
	}, nil
}

func encodeRemoveTrackerResponse() (*gpsgendproto.RemoveTrackerResponse, error) {
	return &gpsgendproto.RemoveTrackerResponse{}, nil
}

func encodeUpdateTrackerResponse() (*gpsgendproto.UpdateTrackerResponse, error) {
	return &gpsgendproto.UpdateTrackerResponse{}, nil
}

func encodeFindTrackerResponse(trk *generator.Tracker) (*gpsgendproto.FindTrackerResponse, error) {
	model, err := tracker2model(trk)
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.FindTrackerResponse{
		Tracker: model,
	}, nil
}

func encodeStartTrackerResponse() (*gpsgendproto.StartTrackerResponse, error) {
	return &gpsgendproto.StartTrackerResponse{}, nil
}

func encodeStopTrackerResponse() (*gpsgendproto.StopTrackerResponse, error) {
	return &gpsgendproto.StopTrackerResponse{}, nil
}

func encodeTrackerStateResponse(state *gpsgenpb.Device) (*gpsgendproto.TrackerStateResponse, error) {
	data, err := proto.Marshal(state)
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.TrackerStateResponse{
		State: data,
	}, nil
}

func encodeAddRoutesResponse() (*gpsgendproto.AddRoutesResponse, error) {
	return &gpsgendproto.AddRoutesResponse{}, nil
}

func encodeRemoveRoutesResponse() (*gpsgendproto.RemoveRouteResponse, error) {
	return &gpsgendproto.RemoveRouteResponse{}, nil
}

func encodeRoutesResponse(routes []*gpsgen.Route) (*gpsgendproto.RoutesResponse, error) {
	data, err := gpsgen.EncodeRoutes(routes)
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.RoutesResponse{
		Routes: data,
	}, nil
}

func encodeRouteAtResponse(route *gpsgen.Route) (*gpsgendproto.RouteAtResponse, error) {
	data, err := gpsgen.EncodeRoutes([]*gpsgen.Route{route})
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.RouteAtResponse{
		Route: data,
	}, nil
}

func encodeRouteIDResponse(route *gpsgen.Route) (*gpsgendproto.RouteByIDResponse, error) {
	data, err := gpsgen.EncodeRoutes([]*gpsgen.Route{route})
	if err != nil {
		return nil, err
	}
	return &gpsgendproto.RouteByIDResponse{
		Route: data,
	}, nil
}

func encodeResetRoutesResponse(ok bool) (*gpsgendproto.ResetRoutesResponse, error) {
	return &gpsgendproto.ResetRoutesResponse{
		Ok: ok,
	}, nil
}

func encodeResetNavigatorResponse() (*gpsgendproto.ResetNavigatorResponse, error) {
	return &gpsgendproto.ResetNavigatorResponse{}, nil
}

func encodeToNextRouteResponse(nav *types.Navigator, ok bool) (*gpsgendproto.ToNextRouteResponse, error) {
	return &gpsgendproto.ToNextRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeToPrevRouteResponse(nav *types.Navigator, ok bool) (*gpsgendproto.ToPrevRouteResponse, error) {
	return &gpsgendproto.ToPrevRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToRouteResponse(nav *types.Navigator, ok bool) (*gpsgendproto.MoveToRouteResponse, error) {
	return &gpsgendproto.MoveToRouteResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToRouteByIDResponse(nav *types.Navigator, ok bool) (*gpsgendproto.MoveToRouteByIDResponse, error) {
	return &gpsgendproto.MoveToRouteByIDResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToTrackResponse(nav *types.Navigator, ok bool) (*gpsgendproto.MoveToTrackResponse, error) {
	return &gpsgendproto.MoveToTrackResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToTrackByIDResponse(nav *types.Navigator, ok bool) (*gpsgendproto.MoveToTrackByIDResponse, error) {
	return &gpsgendproto.MoveToTrackByIDResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeMoveToSegmentResponse(nav *types.Navigator, ok bool) (*gpsgendproto.MoveToSegmentResponse, error) {
	return &gpsgendproto.MoveToSegmentResponse{
		Navigator: navigator2model(nav),
		Ok:        ok,
	}, nil
}

func encodeAddSensorResponse(s *types.Sensor) (*gpsgendproto.AddSensorResponse, error) {
	return &gpsgendproto.AddSensorResponse{
		SensorId: s.ID().String(),
	}, nil
}

func encodeRemoveSensorResponse() (*gpsgendproto.RemoveSensorResponse, error) {
	return &gpsgendproto.RemoveSensorResponse{}, nil
}

func encodeSensorsResponse(sensors []*types.Sensor) (*gpsgendproto.SensorsResponse, error) {
	return &gpsgendproto.SensorsResponse{
		Sensors: sensors2model(sensors),
	}, nil
}

func encodeShutdownTrackerResponse() (*gpsgendproto.ShutdownTrackerResponse, error) {
	return &gpsgendproto.ShutdownTrackerResponse{}, nil
}

func encodeResumeTrackerResponse() (*gpsgendproto.ResumeTrackerResponse, error) {
	return &gpsgendproto.ResumeTrackerResponse{}, nil
}

func encodeNewTrackerErrorResponse(err error) (*gpsgendproto.NewTrackerResponse, error) {
	return &gpsgendproto.NewTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeSearchTrackersErrorResponse(err error) (*gpsgendproto.SearchTrackersResponse, error) {
	return &gpsgendproto.SearchTrackersResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveTrackerErrorResponse(err error) (*gpsgendproto.RemoveTrackerResponse, error) {
	return &gpsgendproto.RemoveTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeUpdateTrackerErrorResponse(err error) (*gpsgendproto.UpdateTrackerResponse, error) {
	return &gpsgendproto.UpdateTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeFindTrackerErrorResponse(err error) (*gpsgendproto.FindTrackerResponse, error) {
	return &gpsgendproto.FindTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeStartTrackerErrorResponse(err error) (*gpsgendproto.StartTrackerResponse, error) {
	return &gpsgendproto.StartTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeStopTrackerErrorResponse(err error) (*gpsgendproto.StopTrackerResponse, error) {
	return &gpsgendproto.StopTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeTrackerStateErrorResponse(err error) (*gpsgendproto.TrackerStateResponse, error) {
	return &gpsgendproto.TrackerStateResponse{
		Error: newError(err),
	}, nil
}

func encodeAddRoutesErrorResponse(err error) (*gpsgendproto.AddRoutesResponse, error) {
	return &gpsgendproto.AddRoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveRouteErrorResponse(err error) (*gpsgendproto.RemoveRouteResponse, error) {
	return &gpsgendproto.RemoveRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeRoutesErrorResponse(err error) (*gpsgendproto.RoutesResponse, error) {
	return &gpsgendproto.RoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeRouteByIDErrorResponse(err error) (*gpsgendproto.RouteByIDResponse, error) {
	return &gpsgendproto.RouteByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeRouteAtErrorResponse(err error) (*gpsgendproto.RouteAtResponse, error) {
	return &gpsgendproto.RouteAtResponse{
		Error: newError(err),
	}, nil
}

func encodeResetRoutesErrorResponse(err error) (*gpsgendproto.ResetRoutesResponse, error) {
	return &gpsgendproto.ResetRoutesResponse{
		Error: newError(err),
	}, nil
}

func encodeResetNavigatorErrorResponse(err error) (*gpsgendproto.ResetNavigatorResponse, error) {
	return &gpsgendproto.ResetNavigatorResponse{
		Error: newError(err),
	}, nil
}

func encodeToNextRouteErrorResponse(err error) (*gpsgendproto.ToNextRouteResponse, error) {
	return &gpsgendproto.ToNextRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeToPrevRouteErrorResponse(err error) (*gpsgendproto.ToPrevRouteResponse, error) {
	return &gpsgendproto.ToPrevRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToRouteErrorResponse(err error) (*gpsgendproto.MoveToRouteResponse, error) {
	return &gpsgendproto.MoveToRouteResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToRouteByIDErrorResponse(err error) (*gpsgendproto.MoveToRouteByIDResponse, error) {
	return &gpsgendproto.MoveToRouteByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToTrackErrorResponse(err error) (*gpsgendproto.MoveToTrackResponse, error) {
	return &gpsgendproto.MoveToTrackResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToTrackByIDErrorResponse(err error) (*gpsgendproto.MoveToTrackByIDResponse, error) {
	return &gpsgendproto.MoveToTrackByIDResponse{
		Error: newError(err),
	}, nil
}

func encodeMoveToSegmentErrorResponse(err error) (*gpsgendproto.MoveToSegmentResponse, error) {
	return &gpsgendproto.MoveToSegmentResponse{
		Error: newError(err),
	}, nil
}

func encodeAddSensorErrorResponse(err error) (*gpsgendproto.AddSensorResponse, error) {
	return &gpsgendproto.AddSensorResponse{
		Error: newError(err),
	}, nil
}

func encodeRemoveSensorErrorResponse(err error) (*gpsgendproto.RemoveSensorResponse, error) {
	return &gpsgendproto.RemoveSensorResponse{
		Error: newError(err),
	}, nil
}

func encodeSensorsErrorResponse(err error) (*gpsgendproto.SensorsResponse, error) {
	return &gpsgendproto.SensorsResponse{
		Error: newError(err),
	}, nil
}

func encodeShutdownTrackerErrorResponse(err error) (*gpsgendproto.ShutdownTrackerResponse, error) {
	return &gpsgendproto.ShutdownTrackerResponse{
		Error: newError(err),
	}, nil
}

func encodeResumeTrackerErrorResponse(err error) (*gpsgendproto.ResumeTrackerResponse, error) {
	return &gpsgendproto.ResumeTrackerResponse{
		Error: newError(err),
	}, nil
}

func newError(err error) *gpsgendproto.Error {
	return &gpsgendproto.Error{
		Code: int64(transportcodes.FromError(err)),
		Msg:  err.Error(),
	}
}
