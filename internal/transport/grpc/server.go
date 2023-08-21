package grpc

import (
	context "context"

	"github.com/mmadfox/gpsgend/internal/generator"
)

type Server struct {
	UnimplementedGeneratorServiceServer

	generator generator.Service
}

func NewServer(s generator.Service) *Server {
	return &Server{generator: s}
}

func (s *Server) NewTracker(ctx context.Context, req *NewTrackerRequest) (*NewTrackerRequest, error) {
	return nil, nil
}
