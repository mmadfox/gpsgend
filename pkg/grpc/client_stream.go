package grpc

import (
	"context"
	"time"

	"github.com/mmadfox/go-gpsgen"
	gpsgenproto "github.com/mmadfox/go-gpsgen/proto"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
)

const waitStream = 3 * time.Second

func (c *Client) ActiveWatchers(ctx context.Context) ([]Watcher, error) {
	req := gpsgendproto.GetClientsInfoRequest{}

	resp, err := c.trackerCli.GetClientsInfo(ctx, &req)
	if err != nil {
		return nil, toError(err)
	}

	watchers := make([]Watcher, len(resp.Clients))
	for i := 0; i < len(resp.Clients); i++ {
		cli := resp.Clients[i]
		watchers[i] = Watcher{
			ID:        cli.Id,
			StartedAt: time.Unix(cli.Timestamp, 0),
		}
	}

	return watchers, nil
}

func (c *Client) Unwatch(ctx context.Context) error {
	req := gpsgendproto.UnsubscribeRequest{ClientId: c.id.String()}

	resp, err := c.trackerCli.Unsubscribe(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) Watch(
	ctx context.Context,
	onPacket func(*gpsgenproto.Packet),
	onError func(error) bool,
) error {
	var (
		stream gpsgendproto.TrackerService_SubscribeClient
		err    error
	)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if stream == nil {
			stream, err = c.subscribe(ctx)
			if err != nil {
				stream = nil
				if onError != nil {
					if stop := onError(err); stop {
						return nil
					}
				}
				time.Sleep(waitStream)
				continue
			}
		}

		resp, err := stream.Recv()
		if err != nil {
			stream = nil
			if onError != nil {
				if stop := onError(err); stop {
					return nil
				}
			}
			time.Sleep(waitStream)
			continue
		}
		if onPacket != nil {
			packet, err := gpsgen.PacketFromBytes(resp.Packet)
			if err != nil {
				if onError != nil {
					if stop := onError(err); stop {
						return nil
					}
				}
				continue
			}
			onPacket(packet)
		}
	}
}

func (c *Client) subscribe(ctx context.Context) (gpsgendproto.TrackerService_SubscribeClient, error) {
	return c.trackerCli.Subscribe(ctx, &gpsgendproto.SubscribeRequest{
		ClientId: c.id.String(),
	})
}
