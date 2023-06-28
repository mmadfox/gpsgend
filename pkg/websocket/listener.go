package websocket

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/fasthttp/websocket"
	stdfasthttpws "github.com/fasthttp/websocket"
	gpsgenpb "github.com/mmadfox/go-gpsgen/proto"
	"google.golang.org/protobuf/proto"
)

const (
	maxMessageSize = 512
	pongWait       = 60 * time.Second
	writeWait      = 10 * time.Second
	pingPeriod     = (pongWait * 9) / 10
)

var errBreak = errors.New("break")

type Handler interface {
	ServeGPS(*gpsgenpb.Device)
}

type HandlerFunc func(*gpsgenpb.Device)

func (hf HandlerFunc) ServeGPS(dev *gpsgenpb.Device) {
	hf(dev)
}

type Listener struct {
	OnResponse func(*http.Response) error
	OnError    func(error)

	conn    *stdfasthttpws.Conn
	headers http.Header
}

func NewListener() *Listener {
	return &Listener{
		headers: make(http.Header),
	}
}

func (l *Listener) Headers() http.Header {
	return l.headers
}

func (l *Listener) Close() error {
	if l.conn == nil {
		return nil
	}
	return nil
}

func (l *Listener) ListenAndServe(ctx context.Context, addr string, h Handler) error {
	ticker := backoff.NewTicker(backoff.NewExponentialBackOff())
	opFunc := func() error {
		return l.listenAndServe(ctx, addr, h)
	}
	var err error
	for range ticker.C {
		if err = opFunc(); err != nil {
			l.onError(err)
			if errors.Is(err, context.Canceled) || errors.Is(err, errBreak) {
				break
			}
			continue
		}
		ticker.Stop()
		break
	}
	return err
}

func (l *Listener) listenAndServe(ctx context.Context, addr string, handler Handler) error {
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, addr, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := l.onResponse(resp); err != nil {
		return err
	}
	l.conn = conn

	var readErr error
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPingHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		msgTyp, msg, err := conn.ReadMessage()
		if err != nil {
			readErr = err
			break
		}
		if msgTyp != websocket.BinaryMessage {
			return errBreak
		}
		if len(msg) == 0 {
			continue
		}
		dev := new(gpsgenpb.Device)
		if err := proto.Unmarshal(msg, dev); err != nil {
			readErr = err
			break
		}
		handler.ServeGPS(dev)
	}
	return readErr
}

func (l *Listener) onResponse(resp *http.Response) error {
	if l.OnResponse == nil {
		return nil
	}
	return l.OnResponse(resp)
}

func (l *Listener) onError(err error) {
	if l.OnError == nil {
		return
	}
	l.OnError(err)
}
