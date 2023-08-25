package broker

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"google.golang.org/protobuf/proto"
)

type hub struct {
	mu      sync.RWMutex
	clients map[uuid.UUID]Client
	history *history
}

func newHub() *hub {
	return &hub{
		clients: make(map[uuid.UUID]Client),
		history: newHistory(),
	}
}

func (h *hub) Broadcast(e *gpsgendproto.Event) {
	data, _ := proto.Marshal(e)

	if e.Kind != gpsgendproto.Event_TRACKER_CHANGED {
		h.history.Append(historyItem{
			timestamp: time.Now().Unix(),
			data:      data,
		})
	}

	h.mu.RLock()
	defer h.mu.RUnlock()

	if len(h.clients) == 0 {
		return
	}

	for _, cli := range h.clients {
		cli.Send(data)
	}
}

func (h *hub) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()

	for _, cli := range h.clients {
		cli.Close()
	}

	h.clients = make(map[uuid.UUID]Client)
}

func (h *hub) Register(cid uuid.UUID, cli Client) {
	h.mu.Lock()
	h.clients[cid] = cli
	h.mu.Unlock()

	h.history.ReadFrom(time.Now(), func(i historyItem) {
		cli.Send(i.data)
	})
}

func (h *hub) Unregister(cid uuid.UUID) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	cli, ok := h.clients[cid]
	if !ok {
		return fmt.Errorf("%w %s", ErrClientNotFound, cid)
	}
	delete(h.clients, cid)

	cli.Close()

	if len(h.clients) == 0 {
		h.clients = make(map[uuid.UUID]Client)
	}

	return nil
}
