package publisher

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Publisher struct {
	mu        sync.RWMutex
	inCh      chan []byte
	stopCh    chan struct{}
	closeOnce sync.Once
	clients   map[uuid.UUID]*Client
	logger    *slog.Logger
}

type Client struct {
	ID  uuid.UUID
	Out chan []byte
}

func New(logger *slog.Logger) *Publisher {
	publisher := &Publisher{
		logger:  logger,
		inCh:    make(chan []byte, 16),
		stopCh:  make(chan struct{}),
		clients: make(map[uuid.UUID]*Client, 8),
	}

	go publisher.doSend()

	return publisher
}

func (p *Publisher) NewClient() *Client {
	pid := uuid.New()
	cli := &Client{
		ID:  pid,
		Out: make(chan []byte, 1),
	}

	p.mu.Lock()
	p.clients[pid] = cli
	p.mu.Unlock()

	return cli
}

func (p *Publisher) CloseClient(pid uuid.UUID) {
	p.mu.Lock()
	defer p.mu.Unlock()

	cli, ok := p.clients[pid]
	if !ok {
		return
	}

	close(cli.Out)
	delete(p.clients, pid)
}

func (p *Publisher) Close() {
	p.closeOnce.Do(func() {
		close(p.stopCh)
	})
}

func (p *Publisher) Publish(_ context.Context, data []byte) error {
	select {
	case p.inCh <- data:
	default:
	}
	return nil
}

func (p *Publisher) doSend() {
	defer func() {
		for _, cli := range p.clients {
			close(cli.Out)
		}
	}()
	for {
		select {
		case <-p.stopCh:
			return
		case pck, ok := <-p.inCh:

			if !ok {
				return
			}
			for _, cli := range p.clients {
				select {
				case cli.Out <- pck:
				default:
				}
			}
		}
	}
}
