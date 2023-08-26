package broker

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_history_ReadFrom(t *testing.T) {
	timePeriod := 5 * time.Minute
	queueCap := 64
	h := newHistory(timePeriod, queueCap)

	h.Append(historyItem{
		data:      []byte("-"),
		timestamp: time.Now().Add(-10 * time.Minute).Unix(),
	})

	h.Append(historyItem{
		data:      []byte("1"),
		timestamp: time.Now().Add(-3 * time.Minute).Unix(),
	})

	h.Append(historyItem{
		data:      []byte("2"),
		timestamp: time.Now().Add(-2 * time.Minute).Unix(),
	})

	h.Append(historyItem{
		data:      []byte("3"),
		timestamp: time.Now().Add(-1 * time.Minute).Unix(),
	})

	want := 3
	got := 0
	h.ReadFrom(time.Now(), func(i historyItem) {
		got++
	})
	require.Equal(t, want, got)
}
