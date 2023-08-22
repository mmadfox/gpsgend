package grpc

import (
	context "context"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type TrackerServer struct {
	UnimplementedTrackerServiceServer
	mu      sync.RWMutex
	clients map[uuid.UUID]*client
}

func NewTrackServer() *TrackerServer {
	return &TrackerServer{
		clients: make(map[uuid.UUID]*client),
	}
}

func (s *TrackerServer) OnPacket(data []byte) {
	s.mu.RLock()
	defer s.mu.RLock()

	req := &SubscribeResponse{Packet: data}
	for _, cli := range s.clients {
		if err := cli.stream.Send(req); err != nil {
			select {
			case cli.close <- struct{}{}:
			default:
			}
		}
	}
}

func (s *TrackerServer) Subscribe(req *SubscribeRequest, stream TrackerService_SubscribeServer) error {
	cid, err := uuid.Parse(req.ClientId)
	if err != nil {
		return err
	}

	cli := newClient(cid, stream)

	s.registerCliet(cli)

	for {
		select {
		case <-stream.Context().Done():
			s.unregisterClient(cli)
			return nil
		case <-cli.close:
			s.unregisterClient(cli)
			return nil
		}
	}
}

func (s *TrackerServer) Unsubscribe(ctx context.Context, req *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	cid, err := uuid.Parse(req.ClientId)
	if err != nil {
		return &UnsubscribeResponse{
			Error: newError(err),
		}, nil
	}

	s.mu.RLock()
	cli, ok := s.clients[cid]
	s.mu.RUnlock()
	if !ok {
		return &UnsubscribeResponse{
			Error: &Error{Msg: fmt.Sprintf("client %s not found", req.ClientId)},
		}, nil
	}

	select {
	case cli.close <- struct{}{}:
	default:
	}

	return new(UnsubscribeResponse), nil
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
	id     uuid.UUID
	stream TrackerService_SubscribeServer
	close  chan struct{}
}

func newClient(id uuid.UUID, stream TrackerService_SubscribeServer) *client {
	return &client{
		id:     id,
		stream: stream,
		close:  make(chan struct{}),
	}
}
