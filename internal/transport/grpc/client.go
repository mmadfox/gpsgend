package grpc

import (
	"context"

	"github.com/mmadfox/gpsgend/internal/generator"
)

type Client struct {
	GeneratorServiceClient
	TrackerServiceClient
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) NewTracker(ctx context.Context, opts generator.NewTrackerOptions) (*generator.TrackerView, error) {
	return nil, nil
}
