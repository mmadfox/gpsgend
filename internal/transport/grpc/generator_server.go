package grpc

import (
	context "context"

	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
)

type GeneratorServer struct {
	gpsgendproto.UnimplementedGeneratorServiceServer

	generator generator.Service
}

func NewGeneratorServer(s generator.Service) *GeneratorServer {
	return &GeneratorServer{generator: s}
}

func (s *GeneratorServer) NewTracker(ctx context.Context, req *gpsgendproto.NewTrackerRequest) (*gpsgendproto.NewTrackerResponse, error) {
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

func (s *GeneratorServer) SearchTrackers(ctx context.Context, req *gpsgendproto.SearchTrackersRequest) (*gpsgendproto.SearchTrackersResponse, error) {
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

func (s *GeneratorServer) RemoveTracker(ctx context.Context, req *gpsgendproto.RemoveTrackerRequest) (*gpsgendproto.RemoveTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeRemoveTrackerErrorResponse(err)
	}

	if err := s.generator.RemoveTracker(ctx, trackerID); err != nil {
		return encodeRemoveTrackerErrorResponse(err)
	}

	return encodeRemoveTrackerResponse()
}

func (s *GeneratorServer) UpdateTracker(ctx context.Context, req *gpsgendproto.UpdateTrackerRequest) (*gpsgendproto.UpdateTrackerResponse, error) {
	trackerID, opts, err := decodeUpdateTrackerRequest(req)
	if err != nil {
		return encodeUpdateTrackerErrorResponse(err)
	}

	if err := s.generator.UpdateTracker(ctx, trackerID, opts); err != nil {
		return encodeUpdateTrackerErrorResponse(err)
	}

	return encodeUpdateTrackerResponse()
}

func (s *GeneratorServer) FindTracker(ctx context.Context, req *gpsgendproto.FindTrackerRequest) (*gpsgendproto.FindTrackerResponse, error) {
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

func (s *GeneratorServer) StartTracker(ctx context.Context, req *gpsgendproto.StartTrackerRequest) (*gpsgendproto.StartTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeStartTrackerErrorResponse(err)
	}

	if err := s.generator.StartTracker(ctx, trackerID); err != nil {
		return encodeStartTrackerErrorResponse(err)
	}

	return encodeStartTrackerResponse()
}

func (s *GeneratorServer) StopTracker(ctx context.Context, req *gpsgendproto.StopTrackerRequest) (*gpsgendproto.StopTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeStopTrackerErrorResponse(err)
	}

	if err := s.generator.StopTracker(ctx, trackerID); err != nil {
		return encodeStopTrackerErrorResponse(err)
	}

	return encodeStopTrackerResponse()
}

func (s *GeneratorServer) TrackerState(ctx context.Context, req *gpsgendproto.TrackerStateRequest) (*gpsgendproto.TrackerStateResponse, error) {
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

func (s *GeneratorServer) AddRoutes(ctx context.Context, req *gpsgendproto.AddRoutesRequest) (*gpsgendproto.AddRoutesResponse, error) {
	trackerID, routes, err := decodeAddRoutesRequest(req)
	if err != nil {
		return encodeAddRoutesErrorResponse(err)
	}

	if err := s.generator.AddRoutes(ctx, trackerID, routes); err != nil {
		return encodeAddRoutesErrorResponse(err)
	}

	return encodeAddRoutesResponse()
}

func (s *GeneratorServer) RemoveRoute(ctx context.Context, req *gpsgendproto.RemoveRouteRequest) (*gpsgendproto.RemoveRouteResponse, error) {
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

func (s *GeneratorServer) Routes(ctx context.Context, req *gpsgendproto.RoutesRequest) (*gpsgendproto.RoutesResponse, error) {
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

func (s *GeneratorServer) RouteAt(ctx context.Context, req *gpsgendproto.RouteAtRequest) (*gpsgendproto.RouteAtResponse, error) {
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

func (s *GeneratorServer) RouteByID(ctx context.Context, req *gpsgendproto.RouteByIDRequest) (*gpsgendproto.RouteByIDResponse, error) {
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

func (s *GeneratorServer) ResetRoutes(ctx context.Context, req *gpsgendproto.ResetRoutesRequest) (*gpsgendproto.ResetRoutesResponse, error) {
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

func (s *GeneratorServer) ResetNavigator(ctx context.Context, req *gpsgendproto.ResetNavigatorRequest) (*gpsgendproto.ResetNavigatorResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeResetNavigatorErrorResponse(err)
	}

	if err := s.generator.ResetNavigator(ctx, trackerID); err != nil {
		return encodeResetNavigatorErrorResponse(err)
	}

	return encodeResetNavigatorResponse()
}

func (s *GeneratorServer) ToNextRoute(ctx context.Context, req *gpsgendproto.ToNextRouteRequest) (*gpsgendproto.ToNextRouteResponse, error) {
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

func (s *GeneratorServer) ToPrevRoute(ctx context.Context, req *gpsgendproto.ToPrevRouteRequest) (*gpsgendproto.ToPrevRouteResponse, error) {
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

func (s *GeneratorServer) MoveToRoute(ctx context.Context, req *gpsgendproto.MoveToRouteRequest) (*gpsgendproto.MoveToRouteResponse, error) {
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

func (s *GeneratorServer) MoveToRouteByID(ctx context.Context, req *gpsgendproto.MoveToRouteByIDRequest) (*gpsgendproto.MoveToRouteByIDResponse, error) {
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

func (s *GeneratorServer) MoveToTrack(ctx context.Context, req *gpsgendproto.MoveToTrackRequest) (*gpsgendproto.MoveToTrackResponse, error) {
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

func (s *GeneratorServer) MoveToTrackByID(ctx context.Context, req *gpsgendproto.MoveToTrackByIDRequest) (*gpsgendproto.MoveToTrackByIDResponse, error) {
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

func (s *GeneratorServer) MoveToSegment(ctx context.Context, req *gpsgendproto.MoveToSegmentRequest) (*gpsgendproto.MoveToSegmentResponse, error) {
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

func (s *GeneratorServer) AddSensor(ctx context.Context, req *gpsgendproto.AddSensorRequest) (*gpsgendproto.AddSensorResponse, error) {
	trackerID, sensor, err := decodeAddSensorRequest(req)
	if err != nil {
		return encodeAddSensorErrorResponse(err)
	}

	if err := s.generator.AddSensor(ctx, trackerID, sensor); err != nil {
		return encodeAddSensorErrorResponse(err)
	}

	return encodeAddSensorResponse(sensor)
}

func (s *GeneratorServer) RemoveSensor(ctx context.Context, req *gpsgendproto.RemoveSensorRequest) (*gpsgendproto.RemoveSensorResponse, error) {
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

func (s *GeneratorServer) Sensors(ctx context.Context, req *gpsgendproto.SensorsRequest) (*gpsgendproto.SensorsResponse, error) {
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

func (s *GeneratorServer) ShutdownTracker(ctx context.Context, req *gpsgendproto.ShutdownTrackerRequest) (*gpsgendproto.ShutdownTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeShutdownTrackerErrorResponse(err)
	}

	if err := s.generator.ShutdownTracker(ctx, trackerID); err != nil {
		return encodeShutdownTrackerErrorResponse(err)
	}

	return encodeShutdownTrackerResponse()
}

func (s *GeneratorServer) ResumeTracker(ctx context.Context, req *gpsgendproto.ResumeTrackerRequest) (*gpsgendproto.ResumeTrackerResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeResumeTrackerErrorResponse(err)
	}

	if err := s.generator.ResumeTracker(ctx, trackerID); err != nil {
		return encodeResumeTrackerErrorResponse(err)
	}

	return encodeResumeTrackerResponse()
}

func (s *GeneratorServer) Stats(ctx context.Context, req *gpsgendproto.EmptyRequest) (*gpsgendproto.StatsResponse, error) {
	statsItems, err := s.generator.Stats(ctx)
	if err != nil {
		return encodeStatsErrorResponse(err)
	}

	return encodeStatsResponse(statsItems)
}

func (s *GeneratorServer) Sync(ctx context.Context, req *gpsgendproto.SyncRequest) (*gpsgendproto.SyncResponse, error) {
	trackerID, err := types.ParseID(req.TrackerId)
	if err != nil {
		return encodeSyncErrorResponse(err)
	}

	if err := s.generator.Sync(ctx, trackerID); err != nil {
		return encodeSyncErrorResponse(err)
	}

	return encodeSyncResponse()
}
