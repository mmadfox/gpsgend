package ws

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/mmadfox/go-gpsgen"
	gpsgenproto "github.com/mmadfox/go-gpsgen/proto"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"google.golang.org/protobuf/proto"
)

const (
	pongWait   = 60 * time.Second
	writeWait  = 10 * time.Second
	pingPeriod = (pongWait * 9) / 10
	wait       = 5 * time.Second
)

type Watcher interface {
	OnEvent(e *gpsgendproto.Event)
	OnPacket(*gpsgenproto.Packet)
	OnError(error) bool
}

type Client struct {
	addr   string
	closed uint32
	conn   *websocket.Conn
}

func New(addr string) *Client {
	addr = strings.TrimRight(addr, "/")
	return &Client{
		addr: addr + "/gpsgend/ws",
	}
}

func (c *Client) Watch(ctx context.Context, w Watcher) error {
	atomic.StoreUint32(&c.closed, 0)

	var watchErr error

loop:
	for i := 0; i < 30; i++ {
		conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.addr, nil)
		if err != nil {
			watchErr = err
			if stop := w.OnError(err); stop {
				break loop
			}
			time.Sleep(wait)
			continue
		}

		c.conn = conn
		watchErr = nil

		conn.SetReadDeadline(time.Now().Add(pongWait))
		conn.SetPingHandler(func(string) error {
			conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		})

	innerLoop:
		for {
			msgTyp, msg, err := conn.ReadMessage()
			if err != nil {
				if c.isClose() {
					break loop
				}
				break innerLoop
			}

			if msgTyp != websocket.BinaryMessage {
				continue
			}
			if len(msg) == 0 {
				continue
			}

			event := new(gpsgendproto.Event)
			if err := proto.Unmarshal(msg, event); err != nil {
				if stop := w.OnError(err); stop {
					watchErr = err
					break loop
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
						watchErr = err
						break loop
					}
					continue
				}
				w.OnPacket(pck)
			} else {
				w.OnEvent(event)
			}
		}
	}

	return watchErr
}

func (c *Client) Unwatch() error {
	if c.conn == nil {
		return nil
	}
	atomic.StoreUint32(&c.closed, 1)
	err := c.conn.Close()
	c.conn = nil
	return err
}

func (c *Client) isClose() bool {
	return atomic.LoadUint32(&c.closed) == 1
}
