package grpc

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/mmadfox/go-gpsgen"
	gpsgenproto "github.com/mmadfox/go-gpsgen/proto"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"google.golang.org/protobuf/proto"
)

type Watcher interface {
	OnEvent(e *gpsgendproto.Event)
	OnPacket(*gpsgenproto.Packet)
	OnError(error) bool
}

const waitStream = 3 * time.Second

func (c *Client) Unwatch(ctx context.Context) error {
	req := gpsgendproto.UnsubscribeRequest{ClientId: c.id.String()}

	atomic.StoreUint32(&c.watch, 0)

	resp, err := c.trackerCli.Unsubscribe(ctx, &req)
	if err != nil {
		return toError(err)
	}
	if resp.Error != nil {
		return decodeError(resp.Error)
	}

	return nil
}

func (c *Client) Watch(ctx context.Context, w Watcher) error {
	var stream gpsgendproto.TrackerService_SubscribeClient
	var err error

	atomic.StoreUint32(&c.watch, 1)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
		}

		if c.isUnwatch() {
			return nil
		}

		if stream == nil {
			stream, err = c.subscribe(ctx)
			if err != nil {
				stream = nil
				if stop := w.OnError(err); stop {
					return err
				}
				time.Sleep(waitStream)
				continue
			}
		}

		resp, err := stream.Recv()
		if err != nil {
			if c.isUnwatch() {
				return nil
			}
			stream = nil
			if stop := w.OnError(err); stop {
				return err
			}
			time.Sleep(waitStream)
			continue
		}

		event := new(gpsgendproto.Event)
		if err := proto.Unmarshal(resp.Event, event); err != nil {
			if stop := w.OnError(err); stop {
				return err
			}
		}

		if event.Kind == gpsgendproto.Event_TRACKER_CHANGED {
			payload, ok := event.Payload.(*gpsgendproto.Event_TrackerChangedEvent)
			if !ok {
				continue
			}
			pck, err := gpsgen.PacketFromBytes(payload.TrackerChangedEvent.Packet)
			if err != nil {
				if stop := w.OnError(err); stop {
					return err
				}
				continue
			}
			w.OnPacket(pck)
		} else {
			w.OnEvent(event)
		}
	}
}

func (c *Client) subscribe(ctx context.Context) (gpsgendproto.TrackerService_SubscribeClient, error) {
	return c.trackerCli.Subscribe(ctx, &gpsgendproto.SubscribeRequest{
		ClientId: c.id.String(),
	})
}

func (c *Client) isUnwatch() bool {
	return atomic.LoadUint32(&c.watch) == 0
}
