package device

import "context"

type Publisher interface {
	Publish(ctx context.Context, packet []byte) error
}
