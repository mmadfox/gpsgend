package grpc

import (
	context "context"

	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

type Server struct {
	UnimplementedGeneratorServiceServer

	generator generator.Service
}

func NewServer(s generator.Service) *Server {
	return &Server{generator: s}
}

func (s *Server) NewTracker(ctx context.Context, req *NewTrackerRequest) (*NewTrackerResponse, error) {
	opts, err := decodeNewTrackerRequest(req)
	if err != nil {
		return encodeNewTrackerErrorResponse(err)
	}

	newTracker, err := s.generator.NewTracker(ctx, opts)
	if err != nil {
		return encodeNewTrackerErrorResponse(err)
	}

	return encodeNewTrackerResponse(newTracker)
}

func (s *Server) SearchTrackers(ctx context.Context, req *SearchTrackersRequest) (*SearchTrackersResponse, error) {
	filter, err := decodeSearchTrackersRequest(req)
	if err != nil {
		return encodeSearchTrackersErrorResponse(err)
	}

	result, err := s.generator.SearchTrackers(ctx, filter)
	if err != nil {
		return encodeSearchTrackersErrorResponse(err)
	}

	return encodeSearchTrackersResponse(result)
}

func (s *Server) RemoveTracker(ctx context.Context, req *RemoveTrackerRequest) (*RemoveTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRemoveTrackerErrorResponse(err)
	}

	if err := s.generator.RemoveTracker(ctx, trackerID); err != nil {
		return encodeRemoveTrackerErrorResponse(err)
	}

	return encodeRemoveTrackerResponse()
}

func (s *Server) UpdateTracker(ctx context.Context, req *UpdateTrackerRequest) (*UpdateTrackerResponse, error) {
	trackerID, opts, err := decodeUpdateTrackerRequest(req)
	if err != nil {
		return encodeUpdateTrackerErrorResponse(err)
	}

	if err := s.generator.UpdateTracker(ctx, trackerID, opts); err != nil {
		return encodeUpdateTrackerErrorResponse(err)
	}

	return encodeUpdateTrackerResponse()
}

func (s *Server) FindTracker(ctx context.Context, req *FindTrackerRequest) (*FindTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeFindTrackerErrorResponse(err)
	}

	tracker, err := s.generator.FindTracker(ctx, trackerID)
	if err != nil {
		return encodeFindTrackerErrorResponse(err)
	}

	return encodeFindTrackerResponse(tracker)
}

func (s *Server) StartTracker(ctx context.Context, req *StartTrackerRequest) (*StartTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeStartTrackerErrorResponse(err)
	}

	if err := s.generator.StartTracker(ctx, trackerID); err != nil {
		return encodeStartTrackerErrorResponse(err)
	}

	return encodeStartTrackerResponse()
}

func (s *Server) StopTracker(ctx context.Context, req *StopTrackerRequest) (*StopTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeStopTrackerErrorResponse(err)
	}

	if err := s.generator.StopTracker(ctx, trackerID); err != nil {
		return encodeStopTrackerErrorResponse(err)
	}

	return encodeStopTrackerResponse()
}

func (s *Server) TrackerState(ctx context.Context, req *TrackerStateRequest) (*TrackerStateResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeTrackerStateErrorResponse(err)
	}

	state, err := s.generator.TrackerState(ctx, trackerID)
	if err != nil {
		return encodeTrackerStateErrorResponse(err)
	}

	return encodeTrackerStateResponse(state)
}

func (s *Server) AddRoutes(ctx context.Context, req *AddRoutesRequest) (*AddRoutesResponse, error) {
	trackerID, routes, err := decodeAddRoutesRequest(req)
	if err != nil {
		return encodeAddRoutesErrorResponse(err)
	}

	if err := s.generator.AddRoutes(ctx, trackerID, routes); err != nil {
		return encodeAddRoutesErrorResponse(err)
	}

	return encodeAddRoutesResponse()
}

func (s *Server) RemoveRoute(ctx context.Context, req *RemoveRouteRequest) (*RemoveRouteResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRemoveRouteErrorResponse(err)
	}

	routeID, err := types.ParseID(req.RouteId)
	if err != nil {
		return encodeRemoveRouteErrorResponse(err)
	}

	if err := s.generator.RemoveRoute(ctx, trackerID, routeID); err != nil {
		return encodeRemoveRouteErrorResponse(err)
	}

	return encodeRemoveRoutesResponse()
}

func (s *Server) Routes(ctx context.Context, req *RotuesRequest) (*RoutesResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRoutesErrorResponse(err)
	}

	routes, err := s.generator.Routes(ctx, trackerID)
	if err != nil {
		return encodeRoutesErrorResponse(err)
	}

	return encodeRoutesResponse(routes)
}

func (s *Server) RouteAt(ctx context.Context, req *RouteAtRequest) (*RouteAtResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRouteAtErrorResponse(err)
	}
	routeIndex := int(req.RouteIndex)

	route, err := s.generator.RouteAt(ctx, trackerID, routeIndex)
	if err != nil {
		return encodeRouteAtErrorResponse(err)
	}

	return encodeRouteAtResponse(route)
}

func (s *Server) RouteByID(ctx context.Context, req *RouteByIDRequest) (*RouteByIDResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRouteByIDErrorResponse(err)
	}
	routeID, err := types.ParseID(req.RouteId)
	if err != nil {
		return encodeRouteByIDErrorResponse(err)
	}

	route, err := s.generator.RouteByID(ctx, trackerID, routeID)
	if err != nil {
		return encodeRouteByIDErrorResponse(err)
	}

	return encodeRouteIDResponse(route)
}

func (s *Server) ResetRoutes(ctx context.Context, req *ResetRoutesRequest) (*ResetRoutesResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeResetRoutesErrorResponse(err)
	}

	ok, err := s.generator.ResetRoutes(ctx, trackerID)
	if err != nil {
		return encodeResetRoutesErrorResponse(err)
	}

	return encodeResetRoutesResponse(ok)
}

func (s *Server) ResetNavigator(ctx context.Context, req *ResetNavigatorRequest) (*ResetNavigatorResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeResetNavigatorErrorResponse(err)
	}

	if err := s.generator.ResetNavigator(ctx, trackerID); err != nil {
		return encodeResetNavigatorErrorResponse(err)
	}

	return encodeResetNavigatorResponse()
}

func (s *Server) ToNextRoute(ctx context.Context, req *ToNextRouteRequest) (*ToNextRouteResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeToNextRouteErrorResponse(err)
	}

	navigator, ok, err := s.generator.ToNextRoute(ctx, trackerID)
	if err != nil {
		return encodeToNextRouteErrorResponse(err)
	}

	return encodeToNextRouteResponse(&navigator, ok)
}

func (s *Server) ToPrevRoute(ctx context.Context, req *ToPrevRouteRequest) (*ToPrevRouteResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeToPrevRouteErrorResponse(err)
	}

	navigator, ok, err := s.generator.ToNextRoute(ctx, trackerID)
	if err != nil {
		return encodeToPrevRouteErrorResponse(err)
	}

	return encodeToPrevRouteResponse(&navigator, ok)
}

func (s *Server) MoveToRoute(ctx context.Context, req *MoveToRouteRequest) (*MoveToRouteResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeMoveToRouteErrorResponse(err)
	}
	routeIndex := int(req.RouteIndex)

	navigator, ok, err := s.generator.MoveToRoute(ctx, trackerID, routeIndex)
	if err != nil {
		return encodeMoveToRouteErrorResponse(err)
	}

	return encodeMoveToRouteResponse(&navigator, ok)
}

func (s *Server) MoveToRouteByID(ctx context.Context, req *MoveToRouteByIDRequest) (*MoveToRouteByIDResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeMoveToRouteByIDErrorResponse(err)
	}
	routeID, err := types.ParseID(req.RouteId)
	if err != nil {
		return encodeMoveToRouteByIDErrorResponse(err)
	}

	navigator, ok, err := s.generator.MoveToRouteByID(ctx, trackerID, routeID)
	if err != nil {
		return encodeMoveToRouteByIDErrorResponse(err)
	}

	return encodeMoveToRouteByIDResponse(&navigator, ok)
}

func (s *Server) MoveToTrack(ctx context.Context, req *MoveToTrackRequest) (*MoveToTrackResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeMoveToTrackErrorResponse(err)
	}
	routeIndex := int(req.RouteIndex)
	trackIndex := int(req.TrackIndex)

	navigator, ok, err := s.generator.MoveToTrack(ctx, trackerID, routeIndex, trackIndex)
	if err != nil {
		return encodeMoveToTrackErrorResponse(err)
	}

	return encodeMoveToTrackResponse(&navigator, ok)
}

func (s *Server) MoveToTrackByID(ctx context.Context, req *MoveToTrackByIDRequest) (*MoveToTrackByIDResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeMoveToTrackByIDErrorResponse(err)
	}
	routeID, err := types.ParseID(req.RouteId)
	if err != nil {
		return encodeMoveToTrackByIDErrorResponse(err)
	}
	trackID, err := types.ParseID(req.TrackId)
	if err != nil {
		return encodeMoveToTrackByIDErrorResponse(err)
	}

	navigator, ok, err := s.generator.MoveToTrackByID(ctx, trackerID, routeID, trackID)
	if err != nil {
		return encodeMoveToTrackByIDErrorResponse(err)
	}

	return encodeMoveToTrackByIDResponse(&navigator, ok)
}

func (s *Server) MoveToSegment(ctx context.Context, req *MoveToSegmentRequest) (*MoveToSegmentResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeMoveToSegmentErrorResponse(err)
	}
	routeIndex := int(req.RouteIndex)
	trackIndex := int(req.TrackIndex)
	segmentIndex := int(req.SegmentIndex)

	navigator, ok, err := s.generator.MoveToSegment(ctx, trackerID, routeIndex, trackIndex, segmentIndex)
	if err != nil {
		return encodeMoveToSegmentErrorResponse(err)
	}

	return encodeMoveToSegmentResponse(&navigator, ok)
}

func (s *Server) AddSensor(ctx context.Context, req *AddSensorRequest) (*AddSensorResponse, error) {
	trackerID, sensor, err := decodeAddSensorRequest(req)
	if err != nil {
		return encodeAddSensorErrorResponse(err)
	}

	if err := s.generator.AddSensor(ctx, trackerID, sensor); err != nil {
		return encodeAddSensorErrorResponse(err)
	}

	return encodeAddSensorResponse(sensor)
}

func (s *Server) RemoveSensor(ctx context.Context, req *RemoveSensorRequest) (*RemoveSensorResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRemoveSensorErrorResponse(err)
	}
	sensorID, err := types.ParseID(req.SensorId)
	if err != nil {
		return encodeRemoveSensorErrorResponse(err)
	}

	if err := s.generator.RemoveSensor(ctx, trackerID, sensorID); err != nil {
		return encodeRemoveSensorErrorResponse(err)
	}

	return encodeRemoveSensorResponse()
}

func (s *Server) Sensors(ctx context.Context, req *SensorsRequest) (*SensorsResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeSensorsErrorResponse(err)
	}

	sensors, err := s.generator.Sensors(ctx, trackerID)
	if err != nil {
		return encodeSensorsErrorResponse(err)
	}

	return encodeSensorsResponse(sensors)
}

func (s *Server) ShutdownTracker(ctx context.Context, req *ShutdownTrackerRequest) (*ShutdownTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeShutdownTrackerErrorResponse(err)
	}

	if err := s.generator.ShutdownTracker(ctx, trackerID); err != nil {
		return encodeShutdownTrackerErrorResponse(err)
	}

	return encodeShutdownTrackerResponse()
}

func (s *Server) ResumeTracker(ctx context.Context, req *ResumeTrackerRequest) (*ResumeTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeResumeTrackerErrorResponse(err)
	}

	if err := s.generator.ResumeTracker(ctx, trackerID); err != nil {
		return encodeResumeTrackerErrorResponse(err)
	}

	return encodeResumeTrackerResponse()
}
