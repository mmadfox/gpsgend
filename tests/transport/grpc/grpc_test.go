package grpc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mmadfox/gpsgend/internal/transport/grpc"
	transportgrpc "github.com/mmadfox/gpsgend/pkg/grpc"
	mockgenerator "github.com/mmadfox/gpsgend/tests/mocks/generator"
)

func server(t *testing.T) (*grpc.GeneratorServer, *mockgenerator.MockService) {
	ctrl := gomock.NewController(t)
	generator := mockgenerator.NewMockService(ctrl)
	srv := grpc.NewGeneratorServer(generator)
	return srv, generator
}

func TestGRPC_NewGeneratorServer(t *testing.T) {
	cli := transportgrpc.New()
	cli.NewTracker(context.Background(), transportgrpc.NewTrackerOptions{})
	_ = cli
}
