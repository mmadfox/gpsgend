package broker

import (
	"sync"
	"time"
)

const (
	cap    = 64
	window = 5 * time.Minute
)

type history struct {
	mu   sync.RWMutex
	data []historyItem
	head int
	tail int
}

func newHistory() *history {
	return &history{
		data: make([]historyItem, cap),
		head: -1,
		tail: 0,
	}
}

type historyItem struct {
	timestamp int64
	data      []byte
}

func (h *history) Append(item historyItem) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data[h.mod(h.head+1)] = item
	old := h.head
	h.head = h.mod(h.head + 1)
	if old != -1 && h.head == h.tail {
		h.tail = h.mod(h.tail + 1)
	}
}

func (h *history) ReadFrom(now time.Time, fn func(historyItem)) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	from := now.Add(-window).Unix()
	to := now.Unix()
	for i := 0; i < cap; i++ {
		index := h.mod(i + h.tail)
		item := h.data[index]
		if item.timestamp > from && item.timestamp <= to {
			fn(item)
		}
	}
}

func (h *history) mod(i int) int {
	return i % cap
}
