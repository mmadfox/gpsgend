package broker

import (
	"sync"
	"time"
)

type history struct {
	mu         sync.RWMutex
	data       []historyItem
	timePeriod time.Duration
	queueCap   int
	head       int
	tail       int
}

func newHistory(timePeriod time.Duration, queueCap int) *history {
	return &history{
		timePeriod: timePeriod,
		queueCap:   queueCap,
		data:       make([]historyItem, queueCap),
		head:       -1,
		tail:       0,
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
	from := now.Add(-h.timePeriod).Unix()
	to := now.Unix()
	for i := 0; i < h.queueCap; i++ {
		index := h.mod(i + h.tail)
		item := h.data[index]
		if item.timestamp > from && item.timestamp <= to {
			fn(item)
		}
	}
}

func (h *history) mod(i int) int {
	return i % h.queueCap
}
