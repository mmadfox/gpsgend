package websocket

import "sync"

type client struct {
	out       chan []byte
	closeCh   chan struct{}
	onceClose sync.Once
}

func newClient() *client {
	return &client{
		out:     make(chan []byte, 24),
		closeCh: make(chan struct{}),
	}
}

func (c *client) Close() {
	c.onceClose.Do(func() {
		close(c.closeCh)
	})
}

func (c *client) Send(data []byte) {
	select {
	case c.out <- data:
	default:
	}
}
