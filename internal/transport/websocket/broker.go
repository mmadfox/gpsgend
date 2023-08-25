package websocket

import (
	"github.com/google/uuid"
	"github.com/mmadfox/gpsgend/internal/broker"
)

type Broker interface {
	RegisterClient(id uuid.UUID, c broker.Client)
	Unregister(id uuid.UUID) error
}
