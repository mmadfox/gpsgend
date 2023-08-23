package grpc

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/generator"
	transportgrpc "github.com/mmadfox/gpsgend/internal/transport/grpc"
	"github.com/mmadfox/gpsgend/internal/types"
)

type (
	Tracker           = generator.TrackerView
	Sensor            = types.Sensor
	NewTrackerOptions = generator.NewTrackerOptions
)

type Client struct {
	cli *transportgrpc.Client
}

func New() *Client {
	return &Client{}
}

func (c *Client) NewTracker(ctx context.Context, opts NewTrackerOptions) (*Tracker, error) {
	return c.cli.NewTracker(ctx, opts)
}
