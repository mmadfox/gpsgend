package grpc

import (
	context "context"
	"log/slog"
	"sync"

	"github.com/google/uuid"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
)

type TrackerServer struct {
	gpsgendproto.UnimplementedTrackerServiceServer

	broker Broker
	logger *slog.Logger 
}

func NewTrackServer(b Broker, logger *slog.Logger) *TrackerServer {
	return &TrackerServer{
		UnimplementedTrackerServiceServer: gpsgendproto.UnimplementedTrackerServiceServer{},
		broker:                            b,
		logger: logger,
	}
}

func (s *TrackerServer) Subscribe(
	req *gpsgendproto.SubscribeRequest,
	stream gpsgendproto.TrackerService_SubscribeServer,
) error {
	cid, err := uuid.Parse(req.ClientId)
	if err != nil {
		return err
	}

	cli := newClient()
	s.broker.RegisterClient(cid, cli)

	s.logger.Info("client connected", "id", cid)

	defer func() {
		s.broker.Unregister(cid)
		cli.Close()
		close(cli.out)
		s.logger.Info("client disconnected", "id", cid)
	}()

	resp := new(gpsgendproto.SubscribeResponse)

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-cli.closeCh:
			return nil
		case data, ok := <-cli.out:
			if !ok {
				return nil
			}
			resp.Event = data
			if err := stream.Send(resp); err != nil {
				return err
			}
		}
	}
}

func (s *TrackerServer) Unsubscribe(ctx context.Context, req *gpsgendproto.UnsubscribeRequest) (*gpsgendproto.UnsubscribeResponse, error) {
	cid, err := uuid.Parse(req.ClientId)
	if err != nil {
		return &gpsgendproto.UnsubscribeResponse{
			Error: newError(err),
		}, nil
	}

	if err := s.broker.Unregister(cid); err != nil {
		return &gpsgendproto.UnsubscribeResponse{
			Error: newError(err),
		}, nil
	}

	return new(gpsgendproto.UnsubscribeResponse), nil
}

func (s *TrackerServer) GetClientsInfo(ctx context.Context, _ *gpsgendproto.GetClientsInfoRequest) (*gpsgendproto.GetClientsInfoResponse, error) {
	return &gpsgendproto.GetClientsInfoResponse{}, nil
}

type client struct {
	out       chan []byte
	closeCh   chan struct{}
	onceClose sync.Once
}

func newClient() *client {
	return &client{
		out:     make(chan []byte, 24),
		closeCh: make(chan struct{}),
	}
}

func (c *client) Close() {
	c.onceClose.Do(func() {
		close(c.closeCh)
	})
}

func (c *client) Send(data []byte) {
	select {
	case c.out <- data:
	default:
	}
}
