package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	gpsgenproto "github.com/mmadfox/go-gpsgen/proto"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	id           uuid.UUID
	generatorCli gpsgendproto.GeneratorServiceClient
	trackerCli   gpsgendproto.TrackerServiceClient
}

func New(conn *grpc.ClientConn) *Client {
	return &Client{
		id:           uuid.New(),
		generatorCli: gpsgendproto.NewGeneratorServiceClient(conn),
		trackerCli:   gpsgendproto.NewTrackerServiceClient(conn),
	}
}

func (c *Client) ID() uuid.UUID {
	return c.id
}

func (c *Client) AddTracker(ctx context.Context, opts *AddTrackerOptions) (*Tracker, error) {
	req, err := encodeNewTrackerRequest(opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.generatorCli.NewTracker(ctx, req)
	if err != nil {
		return nil, toError(err)
	}

	return decodeNewTrackerResponse(resp)
}

func (c *Client) SearchTrackers(ctx context.Context, f Filter) (SearchResult, error) {
	req := encodeSearchTrackersRequest(&f)

	resp, err := c.generatorCli.SearchTrackers(ctx, req)
	if err != nil {
		return SearchResult{}, toError(err)
	}

	return decodeSearchTrackersResponse(resp)
}

func (c *Client) RemoveTracker(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.RemoveTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.RemoveTracker(ctx, &req)
	if err != nil {
		return toError(err)
	}

	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) UpdateTracker(ctx context.Context, trackerID string, opts UpdateTrackerOptions) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	if opts.isEmpty() {
		return ErrNoDataForUpdate
	}

	req := encodeUpdateTrackerRequest(trackerID, &opts)

	resp, err := c.generatorCli.UpdateTracker(ctx, req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) FindTracker(ctx context.Context, trackerID string) (*Tracker, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	req := &gpsgendproto.FindTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.FindTracker(ctx, req)
	if err != nil {
		return nil, toError(err)
	}
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}

	return DecodeTracker(resp.Tracker)
}

func (c *Client) StartTracker(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.StartTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.StartTracker(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) StopTracker(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.StopTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.StopTracker(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) TrackerState(ctx context.Context, trackerID string) (*gpsgenproto.Device, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	req := &gpsgendproto.TrackerStateRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.TrackerState(ctx, req)
	if err != nil {
		return nil, toError(err)
	}
	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}
	m := new(gpsgenproto.Device)
	if err := proto.Unmarshal(resp.State, m); err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) AddRoutes(ctx context.Context, trackerID string, newRoutes ...*gpsgen.Route) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	if len(newRoutes) == 0 {
		return ErrNoRoutes
	}

	req, err := encodeAddRouteRequest(trackerID, newRoutes)
	if err != nil {
		return err
	}

	resp, err := c.generatorCli.AddRoutes(ctx, req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) RemoveRoute(ctx context.Context, trackerID, routeID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}
	if err := validateID(routeID, "route.id"); err != nil {
		return err
	}

	req := gpsgendproto.RemoveRouteRequest{
		TrackerId: trackerID,
		RouteId:   routeID,
	}

	resp, err := c.generatorCli.RemoveRoute(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) Routes(ctx context.Context, trackerID string) ([]*gpsgen.Route, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	req := gpsgendproto.RoutesRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.Routes(ctx, &req)
	if err != nil {
		return nil, toError(err)
	}

	return decodeRoutesResponse(resp)
}

func (c *Client) RouteAt(ctx context.Context, trackerID string, routeIndex int) (*gpsgen.Route, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}
	if routeIndex < 0 {
		return nil, ErrInvalidRouteIndex
	}

	req := gpsgendproto.RouteAtRequest{
		TrackerId:  trackerID,
		RouteIndex: int64(routeIndex),
	}

	resp, err := c.generatorCli.RouteAt(ctx, &req)
	if err != nil {
		return nil, toError(err)
	}

	return decodeRouteAtResponse(resp)
}

func (c *Client) RouteByID(ctx context.Context, trackerID, routeID string) (*gpsgen.Route, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}
	if err := validateID(routeID, "route.id"); err != nil {
		return nil, err
	}

	req := gpsgendproto.RouteByIDRequest{
		TrackerId: trackerID,
		RouteId:   routeID,
	}

	resp, err := c.generatorCli.RouteByID(ctx, &req)
	if err != nil {
		return nil, toError(err)
	}

	return decodeRouteByIDResponse(resp)
}

func (c *Client) ResetRoutes(ctx context.Context, trackerID string) (bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return false, err
	}

	req := gpsgendproto.ResetRoutesRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ResetRoutes(ctx, &req)
	if err != nil {
		return false, toError(err)
	}
	if resp.Error != nil {
		return false, decodeError(resp.Error)
	}

	return resp.Ok, nil
}

func (c *Client) ResetNavigator(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.ResetNavigatorRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ResetNavigator(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) ToNextRoute(ctx context.Context, trackerID string) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}

	req := gpsgendproto.ToNextRouteRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ToNextRoute(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeToNextRouteResponse(resp)
}

func (c *Client) ToPrevRoute(ctx context.Context, trackerID string) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}

	req := gpsgendproto.ToPrevRouteRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ToPrevRoute(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeToPrevRouteResponse(resp)
}

func (c *Client) MoveToRoute(ctx context.Context, trackerID string, routeIndex int) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}
	if routeIndex < 0 {
		return nil, false, ErrInvalidRouteIndex
	}

	req := gpsgendproto.MoveToRouteRequest{
		TrackerId:  trackerID,
		RouteIndex: int64(routeIndex),
	}

	resp, err := c.generatorCli.MoveToRoute(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeMoveToRouteResponse(resp)
}

func (c *Client) MoveToRouteByID(ctx context.Context, trackerID, routeID string) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}
	if err := validateID(routeID, "route.id"); err != nil {
		return nil, false, err
	}

	req := gpsgendproto.MoveToRouteByIDRequest{
		TrackerId: trackerID,
		RouteId:   routeID,
	}

	resp, err := c.generatorCli.MoveToRouteByID(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeMoveToRouteByIDResponse(resp)
}

func (c *Client) MoveToTrack(ctx context.Context, trackerID string, routeIndex, trackIndex int) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}
	if routeIndex < 0 {
		return nil, false, ErrInvalidRouteIndex
	}
	if trackIndex < 0 {
		return nil, false, ErrInvalidTrackIndex
	}

	req := gpsgendproto.MoveToTrackRequest{
		TrackerId:  trackerID,
		RouteIndex: int64(routeIndex),
		TrackIndex: int64(trackIndex),
	}

	resp, err := c.generatorCli.MoveToTrack(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeMoveToTrackResponse(resp)
}

func (c *Client) MoveToTrackByID(ctx context.Context, trackerID, routeID, trackID string) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}
	if err := validateID(routeID, "route.id"); err != nil {
		return nil, false, err
	}
	if err := validateID(trackID, "track.id"); err != nil {
		return nil, false, err
	}

	req := gpsgendproto.MoveToTrackByIDRequest{
		TrackerId: trackerID,
		RouteId:   routeID,
		TrackId:   trackID,
	}

	resp, err := c.generatorCli.MoveToTrackByID(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeMoveToTrackByIDResponse(resp)
}

func (c *Client) MoveToSegment(ctx context.Context, trackerID string, routeIndex, trackIndex, segmentIndex int) (*Navigator, bool, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, false, err
	}
	if routeIndex < 0 {
		return nil, false, ErrInvalidRouteIndex
	}
	if trackIndex < 0 {
		return nil, false, ErrInvalidTrackIndex
	}
	if segmentIndex < 0 {
		return nil, false, ErrInvalidSegmentIndex
	}

	req := gpsgendproto.MoveToSegmentRequest{
		TrackerId:    trackerID,
		RouteIndex:   int64(routeIndex),
		TrackIndex:   int64(trackIndex),
		SegmentIndex: int64(segmentIndex),
	}

	resp, err := c.generatorCli.MoveToSegment(ctx, &req)
	if err != nil {
		return nil, false, toError(err)
	}

	return decodeMoveToSegmentResponse(resp)
}

func (c *Client) AddSensor(
	ctx context.Context,
	trackerID string,
	name string,
	min float64,
	max float64,
	amplitude int,
	mode int,
) (string, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return "", err
	}

	req := gpsgendproto.AddSensorRequest{
		TrackerId: trackerID,
		Name:      name,
		Min:       min,
		Max:       max,
		Amplitude: int64(amplitude),
		Mode:      int64(mode),
	}

	resp, err := c.generatorCli.AddSensor(ctx, &req)
	if err != nil {
		return "", toError(err)
	}
	if resp.Error != nil {
		return "", decodeError(resp.Error)
	}

	return resp.SensorId, nil
}

func (c *Client) RemoveSensor(ctx context.Context, trackerID, sensorID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}
	if err := validateID(sensorID, "sensor.id"); err != nil {
		return err
	}

	req := gpsgendproto.RemoveSensorRequest{
		TrackerId: trackerID,
		SensorId:  sensorID,
	}

	resp, err := c.generatorCli.RemoveSensor(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) Sensors(ctx context.Context, trackerID string) ([]*Sensor, error) {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return nil, err
	}

	req := gpsgendproto.SensorsRequest{
		TrackerId: trackerID,
	}

	resp, err := c.generatorCli.Sensors(ctx, &req)
	if err != nil {
		return nil, toError(err)
	}

	return decodeSensorsResponse(resp)
}

func (c *Client) ShutdownTracker(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.ShutdownTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ShutdownTracker(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) ResumeTracker(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	req := gpsgendproto.ResumeTrackerRequest{TrackerId: trackerID}

	resp, err := c.generatorCli.ResumeTracker(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) Stats(ctx context.Context) ([]StatsItem, error) {
	resp, err := c.generatorCli.Stats(ctx, new(gpsgendproto.EmptyRequest))
	if err != nil {
		return nil, toError(err)
	}

	if resp.Error != nil {
		return nil, decodeError(resp.Error)
	}

	return DecodeStats(resp.Stats), nil
}

func (c *Client) Sync(ctx context.Context, trackerID string) error {
	if err := validateID(trackerID, "tracker.id"); err != nil {
		return err
	}

	resp, err := c.generatorCli.Sync(ctx, &gpsgendproto.SyncRequest{TrackerId: trackerID})
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}
