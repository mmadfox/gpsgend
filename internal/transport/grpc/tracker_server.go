package grpc

import (
	context "context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
)

type TrackerServer struct {
	gpsgendproto.UnimplementedTrackerServiceServer

	mu      sync.RWMutex
	clients map[uuid.UUID]*client
}

func NewTrackServer() *TrackerServer {
	return &TrackerServer{
		UnimplementedTrackerServiceServer: gpsgendproto.UnimplementedTrackerServiceServer{},
		clients:                           make(map[uuid.UUID]*client),
	}
}

func (s *TrackerServer) OnPacket(data []byte) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	req := &gpsgendproto.SubscribeResponse{Packet: data}
	for _, cli := range s.clients {
		if err := cli.stream.Send(req); err != nil {
			select {
			case cli.close <- struct{}{}:
			default:
			}
		}
	}
}

func (s *TrackerServer) Subscribe(req *gpsgendproto.SubscribeRequest, stream gpsgendproto.TrackerService_SubscribeServer) error {
	cid, err := uuid.Parse(req.ClientId)
	if err != nil {
		return err
	}

	cli := newClient(cid, stream)

	s.registerCliet(cli)

	defer func() {
		s.unregisterClient(cli)
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-cli.close:
			return nil
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

	s.mu.RLock()
	cli, ok := s.clients[cid]
	s.mu.RUnlock()
	if !ok {
		return &gpsgendproto.UnsubscribeResponse{
			Error: &gpsgendproto.Error{Msg: fmt.Sprintf("client %s not found", req.ClientId)},
		}, nil
	}

	select {
	case cli.close <- struct{}{}:
	default:
	}

	return new(gpsgendproto.UnsubscribeResponse), nil
}

func (s *TrackerServer) GetClientsInfo(ctx context.Context, _ *gpsgendproto.GetClientsInfoRequest) (*gpsgendproto.GetClientsInfoResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	resp := gpsgendproto.GetClientsInfoResponse{
		Clients: make([]*gpsgendproto.ClientInfo, 0, len(s.clients)),
	}

	for _, cli := range s.clients {
		resp.Clients = append(resp.Clients, &gpsgendproto.ClientInfo{
			Id:        cli.id.String(),
			Timestamp: cli.createdAt.Unix(),
		})
	}

	return &resp, nil
}

func (s *TrackerServer) registerCliet(cli *client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients[cli.id] = cli
}

func (s *TrackerServer) unregisterClient(cli *client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.clients, cli.id)
}

type client struct {
	id        uuid.UUID
	stream    gpsgendproto.TrackerService_SubscribeServer
	close     chan struct{}
	createdAt time.Time
}

func newClient(id uuid.UUID, stream gpsgendproto.TrackerService_SubscribeServer) *client {
	return &client{
		id:        id,
		stream:    stream,
		close:     make(chan struct{}),
		createdAt: time.Now(),
	}
}
